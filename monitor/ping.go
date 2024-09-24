package monitor

import (
	"context"
	"net/http"
	"strings"
)

// PingRespone is the response from the Ping endpoint
type PingResponse struct {
	Up bool `json:"up"`
}

type PingMetadata struct {
	SiteID int `json:"site_id"`
	URL string 	`json:"url"`
	CheckedAt string `json:"checked_at"`
	Up bool `json:"up"`
}

type PingResponses struct {
	Responses []PingMetadata `json:"responses"`
}

// Ping pings a specific site and determines whether it's up or down
// 
//encore:api public path=/ping/*url
func Ping(context context.Context, url string) (*PingResponse, error) {
	// if the url does not start with "http:" or "https:", default it to "https:"
	if !strings.HasPrefix(url, "http:") && !strings.HasPrefix(url, "https:") {
		url = "https://" + url
	}

	// Make an HTTP request to check if it's up
	req, err := http.NewRequestWithContext(context, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return &PingResponse{Up: false}, nil
	}

	res.Body.Close()

	// 2xx and 3xx status codes are considered up
	up := res.StatusCode < 400
	return &PingResponse{Up: up}, nil
}