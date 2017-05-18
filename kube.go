package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/lavalamp/client-go-flat/apimachinery/pkg/apis/meta/v1"
	"github.com/lavalamp/client-go-flat/kubernetes"
	"github.com/lavalamp/client-go-flat/pkg/apis/extensions/v1beta1"

	"errors"
	"fmt"
	"github.com/lavalamp/client-go-flat/tools/clientcmd"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

const (
	kubeConfigPath = "KUBERNETES_CONFIG"
)

type KubeClient struct {
	api *kubernetes.Clientset
}

func NewKubeClient() KubeClient {

	config, err := clientcmd.BuildConfigFromFlags("", "")
	if err != nil {
		panic(err.Error())
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal("Failed to create k8s client.", err)
	}

	return KubeClient{api: client}
}

func (kube KubeClient) UpdateDeployment(name, namespace, image string) error {

	log.Infof("Looking for deployment: %s, namespace: %s", name, namespace)
	deployment, err := findDeployment(name, namespace)
	if err != nil {
		return err
	}

	log.Infof("Updating deployment: %s image to %s (namespace: %s)", name, image, namespace)
	if _, err := kube.api.Deployments(namespace).Update(prepareKubeDeployment(deployment, image)); err != nil {
		return errors.New(fmt.Sprintf("Failed to update deployment: %s (namespace: %s, image: %s). %v", name, namespace, image, err))
	}
	log.Infof("Deployment %s has been updated to image %s (namespace %s)", name, image, namespace)

	return nil
}

func findDeployment(name, namespace string) (*v1beta1.Deployment, error) {

	deployments, err := kube.api.Deployments(namespace).List(v1.ListOptions{})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to get deployments. %v", err))
	}
	for _, currDeployment := range deployments.Items {
		if strings.EqualFold(currDeployment.Name, name) {
			return &currDeployment, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("Deployment %s not found (namespace: %s).", name, namespace))
}

func prepareKubeDeployment(deployment *v1beta1.Deployment, image string) *v1beta1.Deployment {

	deployment.Spec.Template.Spec.Containers[0].Image = image
	deployment.ObjectMeta.SetUID("")
	deployment.ObjectMeta.ResourceVersion = ""

	return deployment
}

func getConfigFile() string {

	ret := os.Getenv(kubeConfigPath)
	if ret == "" {
		log.Infof("%s is not defined, using default location", kubeConfigPath)
		usr, err := user.Current()
		if err != nil {
			log.Fatalln("Failed to get home directory.", err)
		}
		ret = filepath.Join(usr.HomeDir, ".kube", "config")
		if _, err := os.Stat(ret); os.IsNotExist(err) {
			log.Fatalf("File %s does not exists.", ret)
		}
	}
	log.Infof("Kube config path: %s", ret)

	return ret
}
