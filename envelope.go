package main

import (
	"encoding/json"
	"fmt"
	"errors"
	"io"
	"strings"

	"github.com/docker/distribution/notifications"
	log "github.com/Sirupsen/logrus"
)

type Repository struct {
	Name string
	Tag  string
}

func GetPushEventRepositories(envelope io.Reader) ([]Repository, error) {

	log.Info("Parsing docker registry events...")
	var ret []Repository
	events, err := toEvents(envelope);
	if err == nil {
		for _, currEvent := range events {
			log.Infof("Event: %s, Image: %s:%s", currEvent.Action,
				currEvent.Target.Repository, currEvent.Target.Tag)
			if strings.EqualFold(currEvent.Action, "push") {
				ret = append(ret, Repository{
					Name: currEvent.Target.Repository,
					Tag: currEvent.Target.Tag})
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