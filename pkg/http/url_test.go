package http

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type urlTestCase struct {
	givenHost     string
	givenSubpaths []string
	expectedUrl   string
}

func TestJoinURL(t *testing.T) {
	for _, testcase := range []urlTestCase{
		urlTestCase{
			givenHost:     "http://example.com",
			givenSubpaths: []string{"foo/bar"},
			expectedUrl:   "http://example.com/foo/bar",
		},
		urlTestCase{
			givenHost:     "http://example.com/foo",
			givenSubpaths: []string{"bar/hao"},
			expectedUrl:   "http://example.com/foo/bar/hao",
		},
		urlTestCase{
			givenHost:     "http://example.com/foo",
			givenSubpaths: []string{"bar/hao", "b/ahd"},
			expectedUrl:   "http://example.com/foo/bar/hao/b/ahd",
		},
	} {
		gotUrl, err := JoinURL(testcase.givenHost, testcase.givenSubpaths...)
		assert.NoError(t, err)
		assert.EqualValues(t, testcase.expectedUrl, gotUrl)
	}

}
