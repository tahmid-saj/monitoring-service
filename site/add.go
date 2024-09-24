package site

import (
	"context"
	"time"
)

// AddParams are the parameters for adding a site to be monitored
type AddParams struct {
	// URL is the URL of the site. If it doesn't contain a schema
	// (like "http:" or "https:") it defaults to "https:"
	URL string `json:"url"`
}

type AddResponse struct {
	// ID is a unique ID for the site
	ID int `json:"id"`
	// URL is the site's URL
	URL string `json:"url"`
	CreatedAt string `json:"created_at"`
	Up bool `json:"up"`
}

// Add adds a new site to the list of monitored sites
// 
//encore:api public method=POST path=/site
func (s *Service) Add(context context.Context, p *AddParams) (*AddResponse, error) {
	site := &Site{
		URL: p.URL,
	}
	if err := s.db.Create(site).Error; err != nil {
		return nil, err
	}

	return &AddResponse{
		ID: site.ID,
		URL: site.URL,
		CreatedAt: time.Now().String(),
		Up: true,
	}, nil
}