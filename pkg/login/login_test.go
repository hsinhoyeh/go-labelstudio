package login

import (
	"context"
	"testing"

	lsgoquery "github.com/hsinhoyeh/go-labelstudio/pkg/goquery"
	lstestutil "github.com/hsinhoyeh/go-labelstudio/pkg/http/testutil"
	"github.com/hsinhoyeh/go-labelstudio/pkg/invite"
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

	sessionResp, err := NewLoginService(c).DefaultLogin(context.Background())
	assert.NoError(t, err)
	assert.True(t, len(sessionResp.SessionID) > 0)
}

func TestSignup(t *testing.T) {
	ctx := context.Background()

	c1, err := lstestutil.NewTestClient()
	assert.NoError(t, err)

	_, err = NewLoginService(c1).DefaultLogin(ctx)
	assert.NoError(t, err)

	inviteService := invite.NewInviteService(c1)
	token, err := inviteService.GetInvitationToken(ctx)
	assert.NoError(t, err)

	c2, err := lstestutil.NewTestClient()
	assert.NoError(t, err)
	loginService2 := NewLoginService(c2)
	signupResponse, err := loginService2.SignUp(ctx, "foo@example.com", "barbarbar", token)
	assert.NoError(t, err)
	assert.True(t, len(signupResponse.SessionID) > 0)

	c3, err := lstestutil.NewTestClient()
	assert.NoError(t, err)
	loginService3 := NewLoginService(c3)
	loginResponse, err := loginService3.LogMeIn(ctx, "foo@example.com", "barbarbar")
	assert.NoError(t, err)
	assert.True(t, len(loginResponse.SessionID) > 0)
}

func TestMakeSignupOrLoginUrl(t *testing.T) {
	for _, testcase := range []struct {
		hosturl     string
		requesturi  string
		nexturi     string
		expectedUrl string
	}{
		{
			hosturl:     "https://foo.bar.com/subpath",
			requesturi:  "/user/login",
			nexturi:     "projects",
			expectedUrl: "https://foo.bar.com/subpath/user/login%3Fnext=/subpath/projects",
		},
		{
			hosturl:     "https://foo.bar.com/subpath",
			requesturi:  "/user/login",
			nexturi:     "",
			expectedUrl: "https://foo.bar.com/subpath/user/login%3Fnext=/subpath/projects",
		},

		{
			hosturl:     "https://foo.bar.com",
			requesturi:  "/user/login",
			nexturi:     "projects",
			expectedUrl: "https://foo.bar.com/user/login%3Fnext=projects",
		},
		{
			hosturl:     "https://foo.bar.com",
			requesturi:  "/user/login",
			nexturi:     "",
			expectedUrl: "https://foo.bar.com/user/login%3Fnext=/projects",
		},
	} {

		p, err := makeLoginUrl(testcase.hosturl, testcase.requesturi, testcase.nexturi)
		assert.NoError(t, err)
		assert.EqualValues(t, testcase.expectedUrl, p)

	}
}
