# kube-distribution

_kube-distribution_ is a process that listen to docker registry push events events,
and automatically run updated [Kubernetes Deployment](https://kubernetes.io/docs/user-guide/deployments/) with new pushed image.

Docker image tag format: <branch>-<kubernetes namespace>-<kubernetes deployment>-<version>
(exaple: master-default-ceribrodeplyment-7)

## Running _Kubernetes Distribution_ as Kubernetes Pod
TODO

## Running _Kubernetes Distribution_ inside a Docker container
```bash
docker run -d -e KUBERNETES_HOST=<Address> -e KUBERNETES_CA_FILE=<ca.cert Path> -e KUBERNETES_TOKEN=<authentication token> --name kube-distribution -p 5050:5050 vayuadm/kube-distribution
```
- `KUBERNETES_HOST` - Kubernetes master host address (default: https://192.168.99.100:8443).
- `KUBERNETES_CA_FILE` - path to `ca.crt` file (default: /<home dir>/.minikube/ca.crt).
- `KUBERNETES_TOKEN` - Kubernetes authentication token.

Example running a _Kubernetes Distribution_ container on minikube:
```
docker run -it --rm -e KUBERNETES_CA_FILE=/mnt/.minikube/ca.crt -e KUBERNETES_TOKEN=$(kubectl describe secret $(kubectl get secrets | grep default | cut -f1 -d ' ') | grep -E '^token' | cut -f2 -d':' | tr -d '\t') -v $HOME/.minikube:/mnt/.minikube -p 5050:5050 vayuadm/kube-distribution
```

## Building the project in OSX
```
$ env GOOS=linux GOARCH=386 go build -o .dist/kube-distribution -v
```