package http

import (
	"net/url"
	"path/filepath"
)

func JoinURL(host string, subpaths ...string) (string, error) {
	u, err := url.Parse(host)
	if err != nil {
		return "", err
	}
	joinPaths := append([]string{u.Path}, subpaths...)
	u.Path = filepath.Join(joinPaths...)
	return u.String(), nil
}
