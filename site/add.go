package site

import "context"

// AddParams are the parameters for adding a site to be monitored
type AddParams struct {
	// URL is the URL of the site. If it doesn't contain a schema
	// (like "http:" or "https:") it defaults to "https:"
	URL string `json:"url"`
}

// Add adds a new site to the list of monitored sites
// 
//encore:api public method=POST path=/site
func (s *Service) Add(context context.Context, p *AddParams) (*Site, error) {
	site := &Site{
		URL: p.URL,
	}
	if err := s.db.Create(site).Error; err != nil {
		return nil, err
	}

	return site, nil
}