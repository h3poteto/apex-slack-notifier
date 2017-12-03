package main

import (
	"time"
)

type Event struct {
	Timestamp time.Time
	Message   *Message
}

func NewEvent(timestamp time.Time, message *Message) *Event {
	return &Event{
		timestamp,
		message,
	}
}

func (e *Event) PostSlack(url, channel string) error {
	s := NewSlack(url)
	return s.postEvent(e, channel)
}

type Message struct {
	AlarmName        string `json:"AlarmName"`
	AlarmDescription string `json:"AlarmDescription"`
	NewStateValue    string `json:"NewStateValue"`
	NewStateReason   string `json:"NewStateReason"`
	OldStateValue    string `json:"OldStateValue"`
}
