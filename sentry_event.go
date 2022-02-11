package main

import (
	"encoding/json"
	"fmt"
)

type SentryEvent struct {
	ProjectName string `json:"project_name"`
	Level       string `json:"level"`
	Message     string `json:"message"`
	URL         string `json:"url"`
	Event       struct {
		Title string `json:"title"`
	} `json:"event"`
}

func DecodeSentryEvent(data string) (*SentryEvent, error) {
	sentryEvent := new(SentryEvent)
	err := json.Unmarshal([]byte(data), sentryEvent)
	if err != nil {
		return nil, fmt.Errorf("sentry event decode failed with error %w", err)
	}
	return sentryEvent, nil
}
