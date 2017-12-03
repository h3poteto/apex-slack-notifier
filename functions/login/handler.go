package main

import (
	"encoding/json"
	"github.com/apex/go-apex"
	"github.com/apex/go-apex/cloudwatch"
	"github.com/kelseyhightower/envconfig"
	"log"
)

type EnvConfig struct {
	SlackToken string `envconfig:"SLACK_TOKEN" required:"true"`
}

func handler(evt *cloudwatch.Event, ctx *apex.Context) error {
	var env EnvConfig
	if err := envconfig.Process("", &env); err != nil {
		return err
	}

	// for debug logging
	var d interface{}
	json.Unmarshal(evt.Detail, &d)
	log.Printf("[DEBUG] detail: %+v\n", d)

	var detail *Detail
	err := json.Unmarshal(evt.Detail, &detail)
	if err != nil {
		return err
	}
	e := NewEvent(
		evt.Account,
		detail.UserIdentity.UserName,
		detail.EventType,
		detail.ResponseElements,
		detail.SourceIPAddress,
		evt.Time,
	)

	return e.PostSlack(env.SlackToken)
}
