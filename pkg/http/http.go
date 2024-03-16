package http

import (
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"

	"golang.org/x/net/publicsuffix"
)

func NewClient(hostURL string) (*Client, error) {
	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		return nil, err
	}
	return &Client{
		httpClient: &http.Client{
			Jar: jar,
		},
		hostURL: hostURL,
	}, nil
}

type Client struct {
	httpClient *http.Client
	hostURL    string
}

func (c *Client) HostURL() string {
	return c.hostURL
}

func (c *Client) Jar() http.CookieJar {
	return c.httpClient.Jar
}

func (c *Client) Do(req *http.Request) ([]byte, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode > 300 || resp.StatusCode < 200 {
		return nil, fmt.Errorf("invalid resposne code :%d", resp.StatusCode)
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}
