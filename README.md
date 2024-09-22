# monitoring-service

Pub/Sub monitoring service for subscribing to sites and publishing events when they go down or come back up. Developed using Go, TypeScript, PostgreSQL, Webhooks, Encore.

This is an application developed using Go and TypeScript that continuously monitors the uptime of a list of websites provided.

If the website goes down from an "up" state or goes up from a "down" state, Slack messages will be sent to a pre-defined channel using Webhooks. Additional notifications and emails will also be sent using AWS SNS and SES.

