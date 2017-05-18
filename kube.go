package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/lavalamp/client-go-flat/apimachinery/pkg/apis/meta/v1"
	"github.com/lavalamp/client-go-flat/kubernetes"
	"github.com/lavalamp/client-go-flat/pkg/apis/extensions/v1beta1"
	"github.com/lavalamp/client-go-flat/rest"

	"errors"
	"fmt"
	"github.com/lavalamp/client-go-flat/tools/clientcmd"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

const (
	hostUrl        = "KUBERNETES_HOST"
	caFile         = "KUBERNETES_CA_FILE"
	secretToken    = "KUBERNETES_TOKEN"
	kubeConfigPath = "KUBERNETES_CONFIG"
)

type KubeClient struct {
	api *kubernetes.Clientset
}

func NewKubeClient() KubeClient {

	config, err := clientcmd.BuildConfigFromFlags("", os.Getenv(kubeConfigPath))
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

func getKubeConfig() *rest.Config {

	host := os.Getenv(hostUrl)
	if host == "" {
		host = "https://192.168.99.100:8443"
		log.Infof("%s is not defined, using %s", hostUrl, host)
	}

	token := os.Getenv(secretToken)
	if token == "" {
		log.Fatalf("Empty %s", secretToken)
	}

	return &rest.Config{
		Host:            host,
		BearerToken:     token,
		TLSClientConfig: rest.TLSClientConfig{CAFile: getCaFile()},
	}
}

func getCaFile() string {

	ret := os.Getenv(caFile)
	if ret == "" {
		log.Infof("%s is not defined, looking for ca.crt file in .minikube folder under home directory.", caFile)
		usr, err := user.Current()
		if err != nil {
			log.Fatalln("Failed to get home directory.", err)
		}
		ret = filepath.Join(usr.HomeDir, ".minikube", "ca.crt")
		if _, err := os.Stat(ret); os.IsNotExist(err) {
			log.Fatalf("File %s does not exists.", ret)
		}
	}

	return ret
}
