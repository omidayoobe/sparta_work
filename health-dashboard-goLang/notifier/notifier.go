package notifier

import "encoding/json"

type SlackMessage struct {
	Text      string  `json:"text"`
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
		Text:     event.EventType,
	}
	return msg, nil
}