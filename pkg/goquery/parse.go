package goquery

import (
	"bytes"
	"context"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
	log "github.com/golang/glog"

	lshttp "github.com/hsinhoyeh/go-labelstudio/pkg/http"
)

type Parser func(*goquery.Document) (string, bool)

func ParseHTML(ctx context.Context, httpClient *lshttp.Client, url string, p Parser) (string, bool, error) {
	debugEnabled := os.Getenv("LABEL_STUDIO_DEBUG") != ""
	
	if debugEnabled {
		log.Info("[DEBUG] ParseHTML called for URL: %s", url)
	}
	
	loginFormInBytes, err := httpGet(ctx, httpClient, url)
	if err != nil {
		if debugEnabled {
			log.Error("[DEBUG] HTTP GET failed: %+v", err)
		}
		return "", false, err
	}
	
	if debugEnabled {
		log.Info("[DEBUG] Retrieved %d bytes of HTML content", len(loginFormInBytes))
	}
	
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
	debugEnabled := os.Getenv("LABEL_STUDIO_DEBUG") != ""
	
	if debugEnabled {
		log.Info("[DEBUG] ParseRawHTML called with %d bytes", len(htmlBody))
	}
	
	document, err := goquery.NewDocumentFromReader(bytes.NewReader(htmlBody))
	if err != nil {
		if debugEnabled {
			log.Error("[DEBUG] Failed to create goquery document: %+v", err)
		}
		return "", false, err
	}

	val, found := p(document)
	
	if debugEnabled {
		log.Info("[DEBUG] Parser result - found: %v, value: %s", found, val)
	}
	
	return val, found, nil
}
