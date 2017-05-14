package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"

	log "github.com/Sirupsen/logrus"
	"github.com/vayuadm/kube-distribution/dockerhub"
)

type DockerRepo struct {
	Name string
	Tag  string
}

func GetPushEventRepositories(dockerHubMessageReader io.Reader) ([]DockerRepo, error) {

	log.Info("Parsing docker registry events...")
	var ret []DockerRepo
	dockerhubMessages, err := toDockerhubMessages(dockerHubMessageReader)
	if err == nil {
		log.Infof("Found %d dockerhub webhook message(s)", len(dockerhubMessages))
		for _, currMessage := range dockerhubMessages {
			log.Infof("Message: %s, Image: %s:%s", currMessage,
				currMessage.Repository.RepoName, currMessage.PushData.Tag)
			ret = append(ret, DockerRepo{
				Name: currMessage.Repository.RepoName,
				Tag:  currMessage.PushData.Tag})
		}
	}

	return ret, err
}

func toDockerhubMessages(dockerHubMessageReader io.Reader) ([]dockerhub.Webhook, error) {

	var dockerHubMessage []dockerhub.Webhook
	decoder := json.NewDecoder(dockerHubMessageReader)
	err := decoder.Decode(&dockerHubMessage)
	if err != nil {
		message := fmt.Sprintf("Failed to decode docker hub webhook: %s", err)

		return nil, errors.New(message)
	}

	return dockerHubMessage, nil
}
