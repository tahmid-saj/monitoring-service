package monitor

import (
	"context"
	"testing"
)

func TestPing(t *testing.T) {
	context := context.Background()
	tests := []struct {
		URL string
		Up bool
	} {
		{ "google.com", true },
		{ "youtube.com", true },
		// Test both with and without "https://"
		{"httpbin.org/status/200", true},
		{"https://httpbin.org/status/200", true},

		// 4xx and 5xx should considered down.
		{"httpbin.org/status/400", false},
		{"https://httpbin.org/status/500", false},
		// Invalid URLs should be considered down.
		{"invalid://scheme", false},
	}

	for _, test := range tests {
		res, err := Ping(context, test.URL)

		if err != nil {
			t.Errorf("url %s: unexpected error occurred: %v", test.URL, err)
		} else if res.Up != test.Up {
			t.Errorf("url %s: got up=%v, want %v", test.URL, res.Up, test.Up)
		}
	}
}