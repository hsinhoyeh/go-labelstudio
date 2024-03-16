package testutil

import (
	"errors"
	"os"

	lshttp "github.com/hsinhoyeh/go-labelstudio/pkg/http"
)

func NewTestClient() (*lshttp.Client, error) {
	hostUrl := os.Getenv("LABEL_STUDIO_HOST")

	if hostUrl == "" {
		return nil, errors.New("hostUrl is invalid")
	}

	c, err := lshttp.NewClient(hostUrl)
	if err != nil {
		return nil, err
	}
	return c, nil
}
