package dockerhub

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"

	log "github.com/Sirupsen/logrus"
)

type DockerRepo struct {
	Name string
	Tag  string
}

func GetPushEventRepositories(dockerHubMessageReader io.Reader) (*DockerRepo, error) {

	log.Info("Parsing docker registry events...")
	var ret *DockerRepo
	dockerhubMessage, err := toDockerhubMessage(dockerHubMessageReader)
	if err == nil && dockerhubMessage != (Webhook{}) {
		log.Infof("Found dockerhub webhook message")
		log.Infof("Message: %s, Image: %s:%s", dockerhubMessage,
			dockerhubMessage.Repository.RepoName, dockerhubMessage.PushData.Tag)
		ret = &DockerRepo{
			Name: dockerhubMessage.Repository.RepoName,
			Tag:  dockerhubMessage.PushData.Tag}
	}

	return ret, err
}

func toDockerhubMessage(dockerHubMessageReader io.Reader) (Webhook, error) {

	var dockerHubMessage Webhook
	decoder := json.NewDecoder(dockerHubMessageReader)
	err := decoder.Decode(&dockerHubMessage)
	if err != nil {
		message := fmt.Sprintf("Failed to decode docker hub webhook: %s", err)
		return Webhook{}, errors.New(message)
	}

	return dockerHubMessage, nil
}
