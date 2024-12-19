package http

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"os"

	"golang.org/x/net/publicsuffix"

	lsopenapi "github.com/hsinhoyeh/go-labelstudio/api/go-openapiv2"
)

type Options struct {
	caCert           string
	insureSkipVerify bool
}

type clientOption interface {
	apply(*Options)
}

func WithCaCert(caCert string) clientOption {
	return withCaCert{caCert}
}

type withCaCert struct {
	caCert string
}

func (w withCaCert) apply(o *Options) {
	o.caCert = w.caCert
}

func WithSkipVerify(b bool) clientOption {
	return withSkipVerify{b}
}

type withSkipVerify struct {
	insureSkipVerify bool
}

func (w withSkipVerify) apply(o *Options) {
	o.insureSkipVerify = w.insureSkipVerify
}

func NewClient(hostURL string, opts ...clientOption) (*Client, error) {
	o := &Options{}
	for _, opt := range opts {
		opt.apply(o)
	}

	var transport http.RoundTripper
	if o.caCert != "" {
		caCert, err := os.ReadFile(o.caCert)
		if err != nil {
			return nil, err
		}
		caCertPool := x509.NewCertPool()
		if !caCertPool.AppendCertsFromPEM(caCert) {
			log.Fatalf("Failed to append CA certificate to the pool")
		}
		transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: caCertPool,
			},
		}
	} else if o.insureSkipVerify {
		transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}
	}

	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		return nil, err
	}
	return &Client{
		httpClient: &http.Client{
			Transport: transport,
			Jar:       jar,
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

func (c *Client) OpenAPIClient() *lsopenapi.APIClient {
	apiclient := lsopenapi.NewAPIClient(
		&lsopenapi.Configuration{
			BasePath:   c.hostURL,
			HTTPClient: c.httpClient,
			DefaultHeader: map[string]string{
				"Referer": c.hostURL,
			},
		},
	)
	return apiclient
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
