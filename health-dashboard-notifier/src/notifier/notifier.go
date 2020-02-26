package notifier

import (
	"encoding/json"
	"fmt"
	"strings"
)

type SlackMessage struct {
	Text      string `json:"text"`
	IconEmoji string `json:"icon_emoji,omitempty"`
}

type Description struct {
	LatestDescription string `json:"latestDescription"`
}

type AffectedEntity struct {
	EntityValue string `json:"entityValue"`
}

type CloudwatchEvent struct {
	EventType        string           `json:"eventTypeCode"`
	Service          string           `json:"service"`
	EventDescription []Description    `json:"eventDescription"`
	AffectedEntities []AffectedEntity `json:"affectedEntities"`
}

func CloudWatchEventToMessage(source []byte) (*SlackMessage, error) {
	var event CloudwatchEvent
	if err := json.Unmarshal(source, &event); err != nil {
		return nil, err
	}

	msg := &SlackMessage{
		IconEmoji: ":warning:",
		Text:      textForEvent(event),
	}
	return msg, nil
}

func textForEvent(event CloudwatchEvent) string {
	description := strings.Split(event.EventDescription[0].LatestDescription, "  ")
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("*Service:* %v\n*Event Name:* \n%v\n\n*Event Description:* \n%v\n\n*Affected Entities:* \n*%v*",
		event.Service, event.EventType, description[0], event.AffectedEntities[0].EntityValue))
	return sb.String()
}
