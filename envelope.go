package main

import (
	"encoding/json"
	"fmt"
	"errors"
	"io"

	"github.com/docker/distribution/notifications"
	log "github.com/Sirupsen/logrus"
)

type repository struct {
	Name string
	Tag  string
}

func GetPushEventRepositories(envelope io.Reader) ([]repository, error) {

	var ret []repository
	events, err := toEvents(envelope);
	if err == nil {
		for _, currEvent := range events {
			if currEvent.Action == "push" {
				ret = append(ret, repository{Name:currEvent.Target.Repository, Tag:currEvent.Target.Tag})
			}

		}
	}

	return ret, err
}

func toEvents(envelopeReader io.Reader) ([]notifications.Event, error) {

	var envelope notifications.Envelope
	decoder := json.NewDecoder(envelopeReader)
	err := decoder.Decode(&envelope)
	if err != nil {
		message := fmt.Sprintf("Failed to decode docker registry event's envelope: %s", err)
		log.Error(message)

		return nil, errors.New(message)
	}

	return envelope.Events, nil
}