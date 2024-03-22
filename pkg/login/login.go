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

// SignUp allows a new user to be registered into labelstudio
// invitedToken is required when the public registered is diabled when LABEL_STUDIO_DISABLE_SIGNUP_WITHOUT_LINK is turned on.
func (l *LoginService) SignUp(ctx context.Context, email, password string, invitedToken string) (*SessionResponse, error) {
	signupUrl, err := makeLoginUrl(l.client.HostURL(), "/user/signup", "/projects")
	if err != nil {
		return nil, err
	}
	if len(invitedToken) > 0 {
		signupUrl = fmt.Sprintf("%s/?token=%s", signupUrl, invitedToken)
	}
	csrfToken, err := retrieveCSRFToken(ctx, l.client, signupUrl)
	if err != nil {
		return nil, err
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
	_, err = l.client.Do(req)
	if err != nil {
		return nil, err
	}
	return l.parseSessionResponse(signupUrl)
}

func makeLoginUrl(hosturl string, signupOrLogin string, next string) (string, error) {
	if next == "" {
		next = "/projects"
	}
	u, _ := url.Parse(hosturl)
	return lshttp.JoinURL(hosturl, fmt.Sprintf("%s?next=%s", signupOrLogin, filepath.Join(u.Path, next)))
}

func (l *LoginService) parseSessionResponse(siteUrl string) (*SessionResponse, error) {
	u, _ := url.Parse(siteUrl)
	for _, cookie := range l.client.Jar().Cookies(u) {
		if cookie.Name == "sessionid" {
			return &SessionResponse{
				SessionID: cookie.Value,
			}, nil
		}
	}
	return nil, errors.New("sessionid not found")
}

type SessionResponse struct {
	SessionID string `json:"sessionid"`
}

func (l *LoginService) DefaultLogin(ctx context.Context) (*SessionResponse, error) {
	account := os.Getenv("LABEL_STUDIO_USERNAME")
	password := os.Getenv("LABEL_STUDIO_PASSWORD")

	return l.LogMeIn(ctx, account, password)
}

func (l *LoginService) LogMeIn(ctx context.Context, account string, password string) (*SessionResponse, error) {
	loginUrl, err := makeLoginUrl(l.client.HostURL(), "/user/login", "/projects")
	if err != nil {
		return nil, err
	}
	csrfToken, err := retrieveCSRFToken(ctx, l.client, loginUrl)
	if err != nil {
		return nil, err
	}

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
	_, err = l.client.Do(req)
	if err != nil {
		return nil, err
	}
	return l.parseSessionResponse(loginUrl)
}

// retrieveCSRFToken issue a $GET requests with xxx/user/login to get started the whole conversation
// the server side would issue a CSRF token in cookies
// and a csrfmiddlewaretoken inside the form, where this function would parse and return it.
//
// example:
// Set-Cookie:
// csrftoken=58r3SWadB9FQ17Y9u8ytuLhkTyR1PuT2IEdZdJtUDhrOrfpwWf6aucIUqEWjCVGc; expires=Fri, 14 Mar 2025 06:22:47 GMT; Max-Age=31449600; Path=/; SameSite=Lax
// Set-Cookie:
// sessionid=eyJ1aWQiOiJiZDMwNmMwMC1lY2FjLTRhNDMtYWI2YS03NGY2YTQ5YTQ1ODEiLCJvcmdhbml6YXRpb25fcGsiOjF9:1rl0xr:B2Md-s6kvMBe69k7LzIENQB211Dug--Q0v_bRYyoYxA; expires=Fri, 29 Mar 2024 06:22:47 GMT; HttpOnly; Max-
// <form id="login-form" action="/labeling/user/login/?next=/labeling/projects/" method="post">
//
//	  <input type="hidden" name="csrfmiddlewaretoken" value="ynwRhcis7Cti9Bzcrpag0eJssA8Zz4RSqRp2sGzbei6CYONHfQ0vY60VdOJw3Wd0">
//	  <p><input type="text" class="ls-input" name="email" id="email" placeholder="Email" value=""></p>
//	  <p><input type="password" class="ls-input" name="password" id="password" placeholder="Password"></p>
//	  <p>
//	    <input type="checkbox" id="persist_session" name="persist_session" class="ls-checkbox" checked="checked" style="width: auto;" />
//	    <label for="persist_session">Keep me logged in this browser</label>
//	  </p>
//	  <p><button type="submit" aria-label="Log In" class="ls-button ls-button_look_primary">Log in</button></p>
//	</form>
func retrieveCSRFToken(ctx context.Context, client *lshttp.Client, url string) (string, error) {
	val, found, err := lsgoquery.ParseHTML(ctx, client, url, csrfParser)
	if err != nil {
		return "", err
	}
	if !found {
		return "", errors.New("found nothing with parser")
	}
	return val, nil
}

func csrfParser(doc *goquery.Document) (string, bool) {
	query := "[name=csrfmiddlewaretoken]"
	return doc.Find(query).First().Attr("value")
}
