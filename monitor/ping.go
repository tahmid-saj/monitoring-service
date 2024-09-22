package monitor

import (
	"context"
	"net/http"
	"strings"
)

// response from the Ping endpoint
type PingResponse struct {
	Up bool `json:"up"`
}

// Ping pings a specific site and checks if its up / down
// encore:api public path=/ping/*url
func Ping(context context.Context, url string) (*PingResponse, error) {
	// if the url does not start with "http:" or "https:", default to "https:"
	if !strings.HasPrefix(url, "http:") && !strings.HasPrefix(url, "https:") {
		url = "https://" + url
	}

	// make a request to check if the url is up
	req, err := http.NewRequestWithContext(context, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)


	// if the url is not up
	if err != nil {
		return &PingResponse{Up: false}, nil
	}
	res.Body.Close()

	// 2xx and 3xx status codes are considered up
	up := res.StatusCode < 400
	
	return &PingResponse{Up: up}, nil
}