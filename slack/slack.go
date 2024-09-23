package slack

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"encore.app/monitor"
	"encore.dev/pubsub"
)

type NotifyParams struct {
	// Test is the Slack message text to send
	Text string `json:"text"`
}

var _ = pubsub.NewSubscription(monitor.TransitionTopic, "slack-notification", 
	pubsub.SubscriptionConfig[*monitor.TransitionEvent]{
		Handler: func(context context.Context, event *monitor.TransitionEvent) error {
			// slack message
			msg := fmt.Sprintf("*%s is down!*", event.Site.URL)

			if event.Up {
				msg = fmt.Sprintf("*%s is back up*", event.Site.URL)
			}

			// send the slack notification
			return Notify(context, &NotifyParams{Text: msg})
		},
	},
)

// Notify sends a Slack message to a pre-configured channel using a Slack incoming webhook
// 
//encore:api private
func Notify(context context.Context, p *NotifyParams) error {
	reqBody, err := json.Marshal(p)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(context, "POST", secrets.SlackWebhookURL, bytes.NewReader(reqBody))
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode >= 400 {
		body, _ := io.ReadAll(res.Body)
		return fmt.Errorf("error notifying slack: %s: %s", res.Status, body)
	}

	return nil
}

var secrets struct {
	// SlackWebhookURL defines the slack webhook URL to send uptime notifications to
	SlackWebhookURL string
}