package http

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type LoggingTransport struct {
	Transport http.RoundTripper
}

// RoundTrip executes a single HTTP transaction and logs the request and response
func (t *LoggingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	startTime := time.Now()

	// Log the request
	logRequest(req)

	resp, err := t.Transport.RoundTrip(req)
	if err != nil {
		log.Printf("Error making request: %v", err)
		return nil, err
	}

	// Log the response
	logResponse(resp, startTime)

	return resp, nil
}

// logRequest logs details of the HTTP request
func logRequest(req *http.Request) {
	log.Printf("Request: %s %s\nHeaders: %v\n",
		req.Method, req.URL.String(), req.Header)
}

// logResponse logs details of the HTTP response
func logResponse(resp *http.Response, startTime time.Time) {
	duration := time.Since(startTime)
	log.Printf("Response: %s %s\nStatus: %s\nHeaders: %v\nDuration: %v\n",
		resp.Request.Method, resp.Request.URL.String(), resp.Status, resp.Header, duration)
}

// copyRequestBody copies the request body for logging
func copyRequestBody(req *http.Request) string {
	if req.Body == nil {
		return ""
	}
	buf, err := io.ReadAll(req.Body)
	if err != nil {
		return fmt.Sprintf("Error reading request body: %v", err)
	}
	req.Body = io.NopCloser(bytes.NewBuffer(buf)) // Restore the body
	return string(buf)
}

// copyResponseBody copies the response body for logging
func copyResponseBody(resp *http.Response) string {
	if resp.Body == nil {
		return ""
	}
	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Sprintf("Error reading response body: %v", err)
	}
	resp.Body = io.NopCloser(bytes.NewBuffer(buf)) // Restore the body
	return string(buf)
}
