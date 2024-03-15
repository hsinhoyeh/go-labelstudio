package http

import (
	"net/url"
	"path/filepath"
)

func JoinURL(host, subpath string) (string, error) {
	u, err := url.Parse(host)
	if err != nil {
		return "", err
	}
	u.Path = filepath.Join(u.Path, subpath)
	return u.String(), nil
}
