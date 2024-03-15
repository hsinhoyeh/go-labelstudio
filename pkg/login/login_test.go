package login

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"testing"

	lshttp "github.com/hsinhoyeh/go-labelstudio/pkg/http"
	"github.com/hsinhoyeh/go-labelstudio/testdata"
	"github.com/stretchr/testify/assert"
)

func TestParseCSRFToken(t *testing.T) {
	userLoginBytes, err := testdata.HtmlFs.ReadFile("html/userlogin.html")
	assert.NoError(t, err)

	v, err := parseCSRFToken(userLoginBytes)
	assert.NoError(t, err)
	assert.EqualValues(t, "ynwRhcis7Cti9Bzcrpag0eJssA8Zz4RSqRp2sGzbei6CYONHfQ0vY60VdOJw3Wd0", v)
}

func TestLogin(t *testing.T) {
	account := os.Getenv("LABEL_STUDIO_USERNAME")
	password := os.Getenv("LABEL_STUDIO_PASSWORD")
	hostUrl := os.Getenv("LABEL_STUDIO_HOST")

	if account == "" || password == "" || hostUrl == "" {
		t.Skip("skip this test as envs are not configured")
		return
	}
	u, err := url.Parse(hostUrl)
	assert.NoError(t, err)

	c, err := lshttp.NewClient()
	assert.NoError(t, err)

	ls := NewLoginService(c, hostUrl)
	assert.NoError(t, ls.LogMeIn(context.Background(), account, password))

	for _, cookie := range c.Jar.Cookies(u) {
		fmt.Printf("  %s: %s\n", cookie.Name, cookie.Value)
	}

}
