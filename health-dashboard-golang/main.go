package main

import (
	"context"
	"github.com/hashicorp/go-retryablehttp"
	"health-dashboard-notifier/notifier"
)


type NotificationHandler struct {
	client     *retryablehttp.Client
	webhookUrl string
}

func (n *NotificationHandler) Handler(ctx context.Context, event events.CloudWatchEvent) error {

	message, err := notifier.CloudWatchEventToMessage(event.Detail)
	if err != nil {
		return err
	}
	if err := n.notifySlack(ctx, message); err != nil {
		return err
	}

	return nil
}
