package monitor

import (
	"context"
	"errors"

	"encore.app/site"
	"encore.dev/pubsub"
	"encore.dev/storage/sqldb"
)

// TransitionEvent describes a transition of a monitored site
// from up -> down or from down -> up
type TransitionEvent struct {
	// Site is the monitored site in question
	Site *site.Site `json:"site"`
	// Up is whether the site is now up or down
	Up bool `json:"up"`
}

// TransitionTopic is a pubsub topic with transition events for when a monitored site transitions from up -> down or down -> up
var TransitionTopic = pubsub.NewTopic[*TransitionEvent]("uptime-transition", pubsub.TopicConfig{
	DeliveryGuarantee: pubsub.AtLeastOnce,
})

// getPreviousMeasurement reports whether the given site was up or down in the previous ping
func getPreviousMeasurement(context context.Context, siteID int) (up bool, err error) {
	err = sqldb.QueryRow(context, `
		SELECT up FROM checks
		WHERE site_id = $1
		ORDER BY checked_at DESC
		LIMIT 1
	`, siteID).Scan(&up)

	if errors.Is(err, sqldb.ErrNoRows) {
		// there was no previous ping so treat this as if the site was up before
		return true, nil
	} else if err != nil {
		return false, err
	}

	return up, nil
}

func publishOnTransition(context context.Context, site *site.Site, isUp bool) error {
	// if the site was up then went down
	wasUp, err := getPreviousMeasurement(context, site.ID)
	if err != nil {
		return err
	}
	
	if isUp == wasUp {
		// nothing to do since the site is still up
		return nil
	}

	// publish the event if the site went back up
	_, err = TransitionTopic.Publish(context, &TransitionEvent{
		Site: site,
		Up: isUp,
	})

	return err
}