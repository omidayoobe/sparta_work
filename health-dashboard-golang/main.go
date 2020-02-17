package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/hashicorp/go-retryablehttp"
	"health-dashboard-notifier/notifier"
	"log"
	"net/http"
	"os"
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

func (n *NotificationHandler) notifySlack(ctx context.Context, message *notifier.SlackMessage) error {
	w := &bytes.Buffer{}
	if err := json.NewEncoder(w).Encode(*message); err != nil {
		return err
	}
	req, err := retryablehttp.NewRequest(http.MethodPost, n.webhookUrl, w)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")

	res, err := n.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("expected successful, got response %v", res.Status)
	}
	return nil
}

func main() {
	client := retryablehttp.NewClient()
	webhookUrl, ok := os.LookupEnv("SLACK_WEBHOOK")
	if !ok {
		log.Fatalf("could not look up SLACK_WEBHOOK")
	}
	h := &NotificationHandler{
		client:     client,
		webhookUrl: webhookUrl,
	}
	lambda.Start(h.Handler)
}


