package site

import "context"

// Delete deletes a site by the ID
//
//encore:api public method=DELETE path=/path/:siteID
func (s *Service) Delete(context context.Context, siteID int) error {
	return s.db.Delete(&Site{ID: siteID}).Error
}