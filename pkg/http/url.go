package http

import (
	"net/url"
)

func JoinURL(host string, subpaths ...string) (string, error) {
	return url.JoinPath(host, subpaths...)
}
