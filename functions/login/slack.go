package main

import (
	"github.com/bluele/slack"
	"github.com/pkg/errors"
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
	var result slack.AttachmentField
	if evt.Result.ConsoleLogin != "" {
		result = slack.AttachmentField{
			Title: "ConsoleLogin",
			Value: evt.Result.ConsoleLogin,
		}
	} else if evt.Result.CheckMfa != "" {
		result = slack.AttachmentField{
			Title: "CheckMfa",
			Value: evt.Result.CheckMfa,
		}
	} else {
		return errors.New("unknown response elements")
	}

	params := slack.WebHookPostPayload{
		Username:  "AWS Notifier",
		IconEmoji: ":warning:",
		Channel:   channel,
		Text:      "*Console Login Alert*",
	}
	attachment := slack.Attachment{
		Color: "warning",
		Fields: []*slack.AttachmentField{
			&slack.AttachmentField{
				Title: "AccountID",
				Value: evt.AccountID,
			},
			&slack.AttachmentField{
				Title: "UserName",
				Value: evt.UserName,
			},
			&slack.AttachmentField{
				Title: "EventType",
				Value: evt.EventType,
			},
			&slack.AttachmentField{
				Title: "IPAddress",
				Value: evt.IPAddress,
			},
			&result,
			&slack.AttachmentField{
				Title: "Time",
				Value: evt.Time.Format(time.RFC3339),
			},
		},
	}
	params.Attachments = []*slack.Attachment{&attachment}
	return s.client.PostMessage(&params)
}
