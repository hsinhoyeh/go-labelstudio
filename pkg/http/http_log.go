package http

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type LoggingTransport struct {
	Transport http.RoundTripper
}

// LabelStudioTransport wraps another transport to fix Authorization header format
type LabelStudioTransport struct {
	Transport http.RoundTripper
}

// RoundTrip fixes the Authorization header for Label Studio's "Token" format
func (t *LabelStudioTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	debugEnabled := os.Getenv("LABEL_STUDIO_DEBUG") != ""
	
	// Clone the request to avoid modifying the original
	newReq := req.Clone(req.Context())
	
	// Fix Authorization header: "Bearer token" -> "Token token"
	if auth := newReq.Header.Get("Authorization"); auth != "" && strings.HasPrefix(auth, "Bearer ") {
		token := strings.TrimPrefix(auth, "Bearer ")
		newReq.Header.Set("Authorization", "Token "+token)
		
		if debugEnabled {
			log.Printf("[DEBUG] Fixed Authorization header: Bearer -> Token")
		}
	}
	
	return t.Transport.RoundTrip(newReq)
}

// RoundTrip executes a single HTTP transaction and logs the request and response
func (t *LoggingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	startTime := time.Now()

	// Check if debug mode is enabled via environment variable
	debugEnabled := os.Getenv("LABEL_STUDIO_DEBUG") != ""

	// Log the request
	logRequest(req, debugEnabled)

	resp, err := t.Transport.RoundTrip(req)
	if err != nil {
		log.Printf("Error making request: %v", err)
		return nil, err
	}

	// Log the response
	logResponse(resp, startTime, debugEnabled)

	return resp, nil
}

// logRequest logs details of the HTTP request
func logRequest(req *http.Request, debugEnabled bool) {
	if debugEnabled {
		body := copyRequestBody(req)
		log.Printf("[DEBUG] Request: %s %s\nHeaders: %v\nBody: %s\n",
			req.Method, req.URL.String(), req.Header, body)
		
		// Log invitation token specifically if present in URL
		if strings.Contains(req.URL.String(), "token=") {
			log.Printf("[DEBUG] Invitation token detected in URL: %s\n", req.URL.String())
		}
	} else {
		log.Printf("Request: %s %s\nHeaders: %v\n",
			req.Method, req.URL.String(), req.Header)
	}
}

// logResponse logs details of the HTTP response
func logResponse(resp *http.Response, startTime time.Time, debugEnabled bool) {
	duration := time.Since(startTime)
	if debugEnabled {
		body := copyResponseBody(resp)
		log.Printf("[DEBUG] Response: %s %s\nStatus: %s\nHeaders: %v\nBody: %s\nDuration: %v\n",
			resp.Request.Method, resp.Request.URL.String(), resp.Status, resp.Header, body, duration)
		
		// Log cookies for debugging signup/login issues
		if len(resp.Cookies()) > 0 {
			log.Printf("[DEBUG] Response cookies: %v\n", resp.Cookies())
		}
	} else {
		log.Printf("Response: %s %s\nStatus: %s\nHeaders: %v\nDuration: %v\n",
			resp.Request.Method, resp.Request.URL.String(), resp.Status, resp.Header, duration)
	}
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
