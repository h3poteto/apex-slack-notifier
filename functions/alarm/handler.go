package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/kelseyhightower/envconfig"
)

type EnvConfig struct {
	SlackURL     string `envconfig:"SLACK_URL" required:"true"`
	SlackChannel string `envconfig:"SLACK_CHANNEL" required:"true"`
}

func handler(ctx context.Context, event events.SNSEvent) error {
	var env EnvConfig
	if err := envconfig.Process("", &env); err != nil {
		return err
	}

	j := []byte(event.Records[0].SNS.Message)
	var message Message
	err := json.Unmarshal(j, &message)
	if err != nil {
		return err
	}

	e := NewEvent(event.Records[0].SNS.Timestamp, &message)
	return e.PostSlack(env.SlackURL, env.SlackChannel)
}
