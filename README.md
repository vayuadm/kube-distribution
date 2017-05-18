# kube-distribution

[![CircleCI](https://circleci.com/gh/vayuadm/kube-distribution.svg?style=svg)](https://circleci.com/gh/vayuadm/kube-distribution)

_kube-distribution_ is a process that listen to docker registry push events events,
and automatically run updated [Kubernetes Deployment](https://kubernetes.io/docs/user-guide/deployments/) with new pushed image.

Docker image tag format: {branch}--{kubernetes namespace}--{kubernetes deployment}--{version}
(example: master--default--ceribrodeplyment--7)

## Running _Kubernetes Distribution_ inside a Docker container
```bash
$ docker run -d -e KUBERNETES_CONFIG=$HOME/.kube/config --name kube-distribution -p 5050:5050 vayuadm/kube-distribution
```
- `KUBERNETES_CONFIG` - Path to kube config (If empty - in-cluster configuration is assumed).

## Running _Kubernetes Distribution_ as Kubernetes Pod
When running inside kubernetes cluster, no need to set KUBERNETES_CONFIG environment variable as kube-distribution will load the configuration on its own.

## Building the project in OSX
```
$ env GOOS=linux GOARCH=386 go build -o .dist/kube-distribution -v
```
