package main

import (
	"os"
	"path/filepath"
	"os/user"

	log "github.com/Sirupsen/logrus"
	"github.com/lavalamp/client-go-flat/kubernetes"
	"github.com/lavalamp/client-go-flat/rest"
	"github.com/lavalamp/client-go-flat/apimachinery/pkg/apis/meta/v1"
	"github.com/lavalamp/client-go-flat/pkg/apis/extensions/v1beta1"
)

const (
	hostUrl = "KUBERNETES_HOST"
	caFile = "KUBERNETES_CA_FILE"
	secretToken = "KUBERNETES_TOKEN"
)

type KubeClient struct {
	api *kubernetes.Clientset
}

func NewKubeClient() KubeClient {

	client, err := kubernetes.NewForConfig(getKubeConfig())
	if err != nil {
		log.Fatal("Failed to create k8s client.", err)
	}

	return KubeClient{api: client}
}

func (kube KubeClient) UpdateDeployment(name, namespace, image string) {

	deployments, err := kube.api.Deployments(namespace).List(v1.ListOptions{FieldSelector: name})
	if err != nil {
		log.Fatal("Failed to get deployment name:", name, err)
	}

	if _, err := kube.api.Deployments(namespace).Update(
		prepareKubeDeployment(&deployments.Items[0], image)); err != nil {
		log.Fatal("Failed to update deployment:", name, namespace, image, err)
	}

	log.Infof("Deployment %s has been updated to image %s (namespace %s)", name, namespace, image)
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

	ca := os.Getenv(caFile)
	if ca == "" {
		log.Infof("%s is not defined, looking for ca.crt file in .minikube folder under home directory.", caFile)
		usr, err := user.Current()
		if err != nil {
			log.Fatalln("Failed to get home directory.", err)
		}
		ca = filepath.Join(usr.HomeDir, ".minikube", "ca.crt")
		if _, err := os.Stat(ca); os.IsNotExist(err) {
			log.Fatalf("File %s does not exists.", ca)
		}
	}
	tlsClientConfig := rest.TLSClientConfig{CAFile: ca}

	token := os.Getenv(secretToken)
	if token == "" {
		log.Fatalf("Empty %s", secretToken)
	}

	return &rest.Config{
		Host:            host,
		BearerToken:     token,
		TLSClientConfig: tlsClientConfig,
	}
}