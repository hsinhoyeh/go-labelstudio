package goquery

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"

	lshttp "github.com/hsinhoyeh/go-labelstudio/pkg/http"
)

type Parser func(*goquery.Document) (string, bool)

func ParseHTML(ctx context.Context, httpClient *lshttp.Client, url string, p Parser) (string, bool, error) {
	loginFormInBytes, err := httpGet(ctx, httpClient, url)
	if err != nil {
		return "", false, err
	}
	fmt.Printf("loginFormInBytes:%s\n", string(loginFormInBytes))
	return ParseRawHTML(loginFormInBytes, p)
}

// httpGet sends $GET reqeust to url and return the response in raw bytes
func httpGet(ctx context.Context, httpClient *lshttp.Client, url string) ([]byte, error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := httpClient.Do(r)
	return resp, err
}

func ParseRawHTML(htmlBody []byte, p Parser) (string, bool, error) {
	document, err := goquery.NewDocumentFromReader(bytes.NewReader(htmlBody))
	if err != nil {
		return "", false, err
	}

	val, found := p(document)
	return val, found, nil
}
