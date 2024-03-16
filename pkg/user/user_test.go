package user

import (
	"testing"

	"github.com/stretchr/testify/assert"

	lsgoquery "github.com/hsinhoyeh/go-labelstudio/pkg/goquery"
	"github.com/hsinhoyeh/go-labelstudio/testdata"
)

func TestUserToken(t *testing.T) {
	userLoginBytes, err := testdata.HtmlFs.ReadFile("html/useraccount.html")
	assert.NoError(t, err)

	val, found, err := lsgoquery.ParseRawHTML(userLoginBytes, accessTokenParser)
	assert.NoError(t, err)
	assert.EqualValues(t, "123456789456123456789", val)
	assert.True(t, found)
}
