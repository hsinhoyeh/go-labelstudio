package login

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	lshttp "github.com/hsinhoyeh/go-labelstudio/pkg/http"
)

func NewLoginService(client *http.Client, hostURL string) *LoginService {
	return &LoginService{
		client:  client,
		hostURL: hostURL, // like $LABEL_STUDIO_HOST configured in labelstudio
	}
}

type LoginService struct {
	client  *http.Client
	hostURL string
}

func (l *LoginService) LogMeIn(ctx context.Context, account string, password string) error {

	loginUrl, err := lshttp.JoinURL(l.hostURL, "/user/login")
	if err != nil {
		return err
	}
	csrfToken, err := retrieveCSRFToken(ctx, l.client, loginUrl)
	if err != nil {
		return err
	}

	form := url.Values{}
	form.Add("csrfmiddlewaretoken", csrfToken)
	form.Add("email", account)
	form.Add("password", password)
	req, err := http.NewRequestWithContext(ctx, "POST", loginUrl, strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Referer", loginUrl)
	_, err = httpDo(l.client, req)
	if err != nil {
		return err
	}
	return nil
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
func retrieveCSRFToken(ctx context.Context, client *http.Client, url string) (string, error) {
	loginFormInBytes, err := httpGet(ctx, client, url)
	if err != nil {
		return "", err
	}
	return parseCSRFToken(loginFormInBytes)
}

// httpGet sends $GET reqeust to url and return the response in raw bytes
func httpGet(ctx context.Context, httpClient *http.Client, url string) ([]byte, error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := httpDo(httpClient, r)
	return resp, err
}

func httpDo(httpClient *http.Client, req *http.Request) ([]byte, error) {
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode > 300 || resp.StatusCode < 200 {
		return nil, fmt.Errorf("invalid resposne code :%d", resp.StatusCode)
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func parseCSRFToken(htmlBody []byte) (string, error) {
	fieldName := "csrfmiddlewaretoken"
	return parseHttpForm(htmlBody, fieldName)
}

func parseHttpForm(htmlBody []byte, formFieldName string) (string, error) {
	document, err := goquery.NewDocumentFromReader(bytes.NewReader(htmlBody))
	if err != nil {
		return "", err
	}
	query := fmt.Sprintf("[name=%s]", formFieldName)
	val, exists := document.Find(query).First().Attr("value")
	if !exists {
		return "", errors.New("not found")
	}
	return val, nil
}
