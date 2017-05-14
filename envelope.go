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

func GetPushEventRepositories(dockerHubMessageReader io.Reader) (*DockerRepo, error) {

	log.Info("Parsing docker registry events...")
	var ret *DockerRepo
	dockerhubMessage, err := toDockerhubMessage(dockerHubMessageReader)
	if err == nil && dockerhubMessage != (dockerhub.Webhook{}) {
		log.Infof("Found dockerhub webhook message")
		log.Infof("Message: %s, Image: %s:%s", dockerhubMessage,
			dockerhubMessage.Repository.RepoName, dockerhubMessage.PushData.Tag)
		ret = &DockerRepo{
			Name: dockerhubMessage.Repository.RepoName,
			Tag:  dockerhubMessage.PushData.Tag}
	}

	return ret, err
}

func toDockerhubMessage(dockerHubMessageReader io.Reader) (dockerhub.Webhook, error) {

	var dockerHubMessage dockerhub.Webhook
	decoder := json.NewDecoder(dockerHubMessageReader)
	err := decoder.Decode(&dockerHubMessage)
	if err != nil {
		message := fmt.Sprintf("Failed to decode docker hub webhook: %s", err)
		return dockerhub.Webhook{}, errors.New(message)
	}

	return dockerHubMessage, nil
}
