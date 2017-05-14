package main

import (
	log "github.com/Sirupsen/logrus"

	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
)

var (
	kube          KubeClient
	watchBranches Set
)

func main() {

	kube = NewKubeClient()
	watchBranches = getWatchBranches()

	http.HandleFunc("/events", handler)
	http.ListenAndServe(":5050", nil)
}

func handler(writer http.ResponseWriter, request *http.Request) {

	if repositories, err := GetPushEventRepositories(request.Body); err != nil {
		log.Error(err)
		writer.WriteHeader(http.StatusBadRequest)
	} else {
		for _, currRepository := range repositories {
			branch, namespace, deployment, version, err := parseTag(currRepository.Tag)
			if err != nil {
				log.Error(err)
				continue
			}

			if watchBranches.Contains(branch) {
				image := fmt.Sprintf("%s:%s", currRepository.Name, version)
				if err := kube.UpdateDeployment(deployment, namespace, image); err != nil {
					log.Error(err)
				}
			}
		}
		writer.WriteHeader(http.StatusOK)
	}
}

// example tag: master-default-ceribrouideplyment-7
func parseTag(tag string) (branch, namespace, deployment, version string, err error) {

	ret := strings.Split(tag, "-")
	if len(ret) != 4 || ret[0] == "" || ret[1] == "" || ret[2] == "" || ret[3] == "" {
		return "", "", "", "", errors.New(fmt.Sprintf(
			"Failed to parse docker image tag: %s. Format should be: <branch>-<kubernetes namespace>-<kubernetes deployment name>-<version>", tag))
	}

	return ret[0], ret[1], ret[2], ret[3], nil
}

func getWatchBranches() Set {

	ret := NewSet()
	names := os.Getenv("WATCH_BRANCES")
	if len(names) > 0 {
		for _, currBranch := range strings.Split(names, ",") {
			ret.Add(currBranch)
		}
	} else {
		ret.Add("master")
	}
	log.Info("Watch branches: ", ret.ToArray())

	return ret
}
