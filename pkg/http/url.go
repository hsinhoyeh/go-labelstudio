package http

import (
	"net/url"
	"os"

	log "github.com/golang/glog"
)

func JoinURL(host string, subpaths ...string) (string, error) {
	debugEnabled := os.Getenv("LABEL_STUDIO_DEBUG") != ""
	
	if debugEnabled {
		log.Info("[DEBUG] JoinURL called with host: %s, subpaths: %v", host, subpaths)
	}
	
	result, err := url.JoinPath(host, subpaths...)
	
	if debugEnabled {
		if err != nil {
			log.Error("[DEBUG] JoinURL failed: %+v", err)
		} else {
			log.Info("[DEBUG] JoinURL result: %s", result)
		}
	}
	
	return result, err
}
