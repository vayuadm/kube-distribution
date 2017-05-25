package main

import (
	log "github.com/Sirupsen/logrus"

	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/vayuadm/kube-distribution/dockerhub"
	"github.com/vayuadm/kube-distribution/utils"
)

var (
	kube          KubeClient
	watchBranches utils.Set
)

const delimiter = "--"

func main() {

	kube = NewKubeClient()
	watchBranches = getWatchBranches()

	http.HandleFunc("/events", handler)
	http.ListenAndServe("0.0.0.0:5050", nil)
}

func handler(writer http.ResponseWriter, request *http.Request) {

	if repository, err := dockerhub.GetPushEventRepositories(request.Body); err != nil {
		log.Error(err)
		writer.WriteHeader(http.StatusBadRequest)
	} else {
		branch, namespace, deployment, version, err := parseTag(repository.Tag)
		if err != nil {
			log.Error(err)
			return
		}

		if watchBranches.Contains(branch) {
			image := fmt.Sprintf("%s:%s", repository.Name, version)
			if err := kube.UpdateDeployment(deployment, namespace, image); err != nil {
				log.Error(err)
			}
		}
		writer.WriteHeader(http.StatusOK)
	}
}

// example tag: master--default--ceribrouideplyment--7
func parseTag(tag string) (branch, namespace, deployment, version string, err error) {

	ret := strings.Split(tag, delimiter)
	if len(ret) != 4 || ret[0] == "" || ret[1] == "" || ret[2] == "" || ret[3] == "" {
		return "", "", "", "", fmt.Errorf(
			"Failed to parse docker image tag: %s. Format should be: <branch>%s<kubernetes namespace>%s<kubernetes deployment name>%s<version>", tag, delimiter, delimiter, delimiter)
	}

	return ret[0], ret[1], ret[2], ret[3], nil
}

func getWatchBranches() utils.Set {

	ret := utils.NewSet()
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
