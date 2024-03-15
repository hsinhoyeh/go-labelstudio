package http

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type urlTestCase struct {
	givenHost    string
	givenSubpath string
	expectedUrl  string
}

func TestJoinURL(t *testing.T) {
	for _, testcase := range []urlTestCase{
		urlTestCase{
			givenHost:    "http://example.com",
			givenSubpath: "foo/bar",
			expectedUrl:  "http://example.com/foo/bar",
		},
		urlTestCase{
			givenHost:    "http://example.com/foo",
			givenSubpath: "bar/hao",
			expectedUrl:  "http://example.com/foo/bar/hao",
		},
	} {
		gotUrl, err := JoinURL(testcase.givenHost, testcase.givenSubpath)
		assert.NoError(t, err)
		assert.EqualValues(t, testcase.expectedUrl, gotUrl)
	}

}
