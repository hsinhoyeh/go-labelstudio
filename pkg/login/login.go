package login

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/PuerkitoBio/goquery"
	log "github.com/golang/glog"
	lsgoquery "github.com/hsinhoyeh/go-labelstudio/pkg/goquery"
	lshttp "github.com/hsinhoyeh/go-labelstudio/pkg/http"
)

func NewLoginService(client *lshttp.Client) *LoginService {
	return &LoginService{
		client: client,
	}
}

type LoginService struct {
	client *lshttp.Client
}

// SessionResponse contains the session information returned after authentication
type SessionResponse struct {
	SessionID string `json:"sessionid"`
	CSRFToken string `json:"csrftoken,omitempty"`
}

// SignUp allows a new user to be registered into labelstudio
// invitedToken is required when the public registered is diabled when LABEL_STUDIO_DISABLE_SIGNUP_WITHOUT_LINK is turned on.
func SignUp(ctx context.Context, client *lshttp.Client, email, password string, invitedToken string) (*SessionResponse, error) {
	signupUrl, err := makeLoginUrl(client.HostURL(), "/user/signup", "/projects")
	if err != nil {
		log.Error("ls: generate signup url failed, err:%+v\n", err)
		return nil, err
	}
	if len(invitedToken) > 0 {
		u, _ := url.Parse(signupUrl)
		queries := u.Query()
		queries.Set("token", invitedToken)
		u.RawQuery = queries.Encode()

		signupUrl = u.String()
	}
	csrfToken, err := retrieveCSRFToken(ctx, client, signupUrl)
	if err != nil {
		log.Error("ls:retrieve CSRF token failed, err:%+v\n", err)
		return nil, err
	}

	if csrfToken == "" {
		return nil, errors.New("CSRF token is required for signup but none was found")
	}

	form := url.Values{}
	form.Add("csrfmiddlewaretoken", csrfToken)
	form.Add("email", email)
	form.Add("password", password)
	form.Add("allow_newsletters", "false")
	req, err := http.NewRequestWithContext(ctx, "POST", signupUrl, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Referer", signupUrl)
	_, err = client.Do(req)
	if err != nil {
		log.Error("ls: post signup failed, err:%+v\n", err)
		return nil, err
	}

	// For debugging
	log.Info("ls: signup completed")

	// Get session response and include CSRF token if available
	sessionResp, err := parseSessionResponse(client, signupUrl)
	if err != nil {
		return nil, err
	}

	// Check for CSRF token in cookies after signup
	parsedURL, _ := url.Parse(signupUrl)
	for _, cookie := range client.Jar().Cookies(parsedURL) {
		if cookie.Name == "csrftoken" {
			sessionResp.CSRFToken = cookie.Value
			break
		}
	}

	return sessionResp, nil
}

func makeLoginUrl(hosturl string, signupOrLogin string, next string) (string, error) {
	if next == "" {
		next = "/projects"
	}
	u, _ := url.Parse(hosturl)

	s, err := lshttp.JoinURL(hosturl, fmt.Sprintf("%s?next=%s", signupOrLogin, filepath.Join(u.Path, next)))
	if err != nil {
		log.Error("ls: generate login url failed, url:%s, err:%+v\n", s, err)
		return "", err
	}
	return strings.Replace(s, "%3F", "?", -1), nil
}

func parseSessionResponse(client *lshttp.Client, siteUrl string) (*SessionResponse, error) {
	u, _ := url.Parse(siteUrl)
	sessionResp := &SessionResponse{}

	for _, cookie := range client.Jar().Cookies(u) {
		if cookie.Name == "sessionid" {
			sessionResp.SessionID = cookie.Value
		}
		if cookie.Name == "csrftoken" {
			sessionResp.CSRFToken = cookie.Value
		}
	}

	if sessionResp.SessionID == "" {
		return nil, errors.New("sessionid not found")
	}

	// Check if the sessionid format indicates a logged-in user
	// (Session IDs for logged-in users contain dots)
	isLoggedIn := strings.Contains(sessionResp.SessionID, ".")
	log.Info("ls: session detected, logged in status: %v\n", isLoggedIn)

	return sessionResp, nil
}

func (l *LoginService) DefaultLogin(ctx context.Context) (*SessionResponse, error) {
	account := os.Getenv("LABEL_STUDIO_USERNAME")
	password := os.Getenv("LABEL_STUDIO_PASSWORD")

	return LogMeIn(ctx, l.client, account, password)
}

// LogMeIn handles the login process, including managing CSRF tokens
func LogMeIn(ctx context.Context, client *lshttp.Client, account string, password string) (*SessionResponse, error) {
	loginUrl, err := makeLoginUrl(client.HostURL(), "/user/login", "/projects")
	if err != nil {
		log.Error("ls: generate logmein failed, url:%s, err:%+v\n", loginUrl, err)
		return nil, err
	}

	// Check if we're already logged in first
	sessionResp, err := parseSessionResponse(client, loginUrl)
	if err == nil && sessionResp.SessionID != "" && strings.Contains(sessionResp.SessionID, ".") {
		log.Info("ls: user already has valid session (contains dot), skipping login\n")
		return sessionResp, nil
	}

	// Try to retrieve CSRF token for login
	csrfToken, err := retrieveCSRFToken(ctx, client, loginUrl)
	if err != nil {
		log.Error("ls: parse csrf token failed, err:%+v\n", err)
		return nil, err
	}

	// If no CSRF token found but no error, something is wrong
	// Label Studio should always provide a CSRF token for login page
	if csrfToken == "" {
		log.Error("ls: no csrf token found in login page, cannot proceed with login\n")
		return nil, errors.New("missing CSRF token for login")
	}

	// Proceed with login using CSRF token
	form := url.Values{}
	form.Add("csrfmiddlewaretoken", csrfToken)
	form.Add("email", account)
	form.Add("password", password)
	req, err := http.NewRequestWithContext(ctx, "POST", loginUrl, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Referer", loginUrl)
	_, err = client.Do(req)
	if err != nil {
		return nil, err
	}

	// For debugging
	log.Info("ls: login completed")

	// Parse session and include CSRF token if available
	return parseSessionResponse(client, loginUrl)
}

// retrieveCSRFToken tries to get a CSRF token from the page or from cookies
// Returns the token if found, or an empty string and no error if user is already logged in
func retrieveCSRFToken(ctx context.Context, client *lshttp.Client, loginUrl string) (string, error) {
	log.Info("ls: retrieve csrf token, url:%+s\n", loginUrl)

	// First check if we already have a CSRF token in cookies
	parsedURL, err := url.Parse(loginUrl)
	if err != nil {
		return "", err
	}

	for _, cookie := range client.Jar().Cookies(parsedURL) {
		if cookie.Name == "csrftoken" {
			log.Info("ls: found csrf token in cookies: %s\n", cookie.Value)
			return cookie.Value, nil
		}
	}

	// If not in cookies, try to parse it from the HTML response
	val, found, err := lsgoquery.ParseHTML(ctx, client, loginUrl, csrfParser)
	log.Info("ls: csrf val:%+v, found:%+v, err:%+v\n", val, found, err)
	if err != nil {
		return "", err
	}
	if !found {
		// User might be already logged in - this is not necessarily an error
		log.Info("ls: csrf token not found in response, user might be already logged in\n")
		return "", nil
	}
	return val, nil
}

func csrfParser(doc *goquery.Document) (string, bool) {
	query := "[name=csrfmiddlewaretoken]"
	return doc.Find(query).First().Attr("value")
}
