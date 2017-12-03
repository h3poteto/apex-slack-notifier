package main

import (
	"encoding/json"
	"github.com/apex/go-apex"
	"github.com/apex/go-apex/sns"
	"github.com/kelseyhightower/envconfig"
)

type EnvConfig struct {
	SlackURL     string `envconfig:"SLACK_URL" required:"true"`
	SlackChannel string `envconfig:"SLACK_CHANNEL" required:"true"`
}

func handler(evt *sns.Event, ctx *apex.Context) error {
	var env EnvConfig
	if err := envconfig.Process("", &env); err != nil {
		return err
	}

	j := []byte(evt.Records[0].SNS.Message)
	var message Message
	err := json.Unmarshal(j, &message)
	if err != nil {
		return err
	}

	e := NewEvent(evt.Records[0].SNS.Timestamp, &message)
	return e.PostSlack(env.SlackURL, env.SlackChannel)
}
