package main

import (
	"github.com/bluele/slack"
	"time"
)

type Slack struct {
	client *slack.WebHook
}

func NewSlack(url string) *Slack {
	client := slack.NewWebHook(url)
	return &Slack{
		client,
	}
}

func (s *Slack) postEvent(evt *Event, channel string) error {
	color := "warning"
	icon := ":warning:"
	text := ""
	if evt.Message.NewStateValue == "OK" {
		color = "good"
		icon = ":ok_hand:"
		text = "*[OK]" + evt.Message.AlarmName + "*"
	} else if evt.Message.NewStateValue == "ALARM" {
		color = "danger"
		text = "*[ALARM]" + evt.Message.AlarmName + "*"
	}

	params := slack.WebHookPostPayload{
		Username:  "AWS Notifier",
		IconEmoji: icon,
		Channel:   channel,
		Text:      text,
	}
	attachment := slack.Attachment{
		Color: color,
		Fields: []*slack.AttachmentField{
			&slack.AttachmentField{
				Title: "Body",
				Value: evt.Message.AlarmDescription,
			},
			&slack.AttachmentField{
				Title: "Reason",
				Value: evt.Message.NewStateReason,
			},
			&slack.AttachmentField{
				Title: "OldState",
				Value: evt.Message.OldStateValue,
			},
			&slack.AttachmentField{
				Title: "Time",
				Value: evt.Timestamp.Format(time.RFC3339),
			},
		},
	}
	params.Attachments = []*slack.Attachment{&attachment}
	return s.client.PostMessage(&params)
}
