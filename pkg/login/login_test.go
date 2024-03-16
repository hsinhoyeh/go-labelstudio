package login

import (
	"context"
	"fmt"
	"net/url"
	"testing"

	lsgoquery "github.com/hsinhoyeh/go-labelstudio/pkg/goquery"
	lstestutil "github.com/hsinhoyeh/go-labelstudio/pkg/http/testutil"
	"github.com/hsinhoyeh/go-labelstudio/testdata"
	"github.com/stretchr/testify/assert"
)

func TestParseCSRFToken(t *testing.T) {
	userLoginBytes, err := testdata.HtmlFs.ReadFile("html/userlogin.html")
	assert.NoError(t, err)

	val, found, err := lsgoquery.ParseRawHTML(userLoginBytes, csrfParser)
	assert.NoError(t, err)
	assert.EqualValues(t, "ynwRhcis7Cti9Bzcrpag0eJssA8Zz4RSqRp2sGzbei6CYONHfQ0vY60VdOJw3Wd0", val)
	assert.True(t, found)
}

func TestLogin(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping testing in short mode")
		return
	}
	c, err := lstestutil.NewTestClient()
	assert.NoError(t, err)

	assert.NoError(
		t,
		NewLoginService(c).DefaultLogin(context.Background()),
	)

	u, err := url.Parse(c.HostURL())
	assert.NoError(t, err)
	for _, cookie := range c.Jar().Cookies(u) {
		fmt.Printf("  %s: %s\n", cookie.Name, cookie.Value)
	}
}
