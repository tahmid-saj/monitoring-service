package site

import "context"

type ListResponse struct {
	// Sites is the list of monitored sites
	Sites []*Site `json:"sites"`
}

// List lists the monitored sites
//
//encore:api public method=GET path=/site
func (s *Service) List(context context.Context) (*ListResponse, error) {
	var sites []*Site
	if err := s.db.Find(&sites).Error; err != nil {
		return nil, err
	}

	return &ListResponse{Sites: sites}, nil
}