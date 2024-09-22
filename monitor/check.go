package monitor

import (
	"context"

	"encore.app/site"
	"encore.dev/cron"
	"encore.dev/storage/sqldb"
	"golang.org/x/sync/errgroup"
)

// cron job to check all tracked sites every 5 minutes
var _ = cron.NewJob("check-all", cron.JobConfig{
	Title: "Check all sites",
	Endpoint: CheckAll,
	Every: 5 * cron.Minute,
})

// Check checks a single site
//
//encore:api public method=POST path=/check/:siteID
func Check(context context.Context, siteID int) error {
	// retrieve the site details from the table
	site, err := site.Get(context, siteID)
	if err != nil {
		return err
	}

	return check(context, site)
}

func check(context context.Context, site *site.Site) error {
	// initiate a Ping to the site
	result, err := Ping(context, site.URL)
	if err != nil {
		return err
	}

	// Publish a Pub/Sub event if the site transitions from up -> down or down -> up
	if err := publishOnTransition(context, site, result.Up); err != nil {
		return err
	}
	
	// insert the update to the table
	_, err = sqldb.Exec(context, `
		INSERT INTO checks (site_id, up, checked_at)
		VALUES ($1, $2, NOW())
	`, site.ID, result.Up)
	
	return err
}

// CheckAll checks all the sites provided
//
//encore:api public method=POST path=/check-all
func CheckAll(context context.Context) error {
	// Get all the tracked sites
	res, err := site.List(context)
	if err != nil {
		return err
	}

	// Check up to 8 sites concurrently
	g, contcontext := errgroup.WithContext(context)
	g.SetLimit(8)

	// check all 8 sites
	for _, site := range res.Sites {
		// capture for closure
		site := site
		g.Go(func() error {
			return check(contcontext, site)
		})
	}

	return g.Wait()
}