package itest

import (
	"net/http"
	"io"
	"strings"
	"fmt"
)

const (
	Push = "push"
)

func SendEvent(eventType, image, tag string) (response *http.Response, err error) {

	return http.Post("http://localhost:5050/events",
		"",
		getEventEnvelope(eventType, image, tag))
}

func getEventEnvelope(eventType, image, tag string) io.Reader {

	envelope := strings.TrimSpace(`
		{"events": [
		      {
			 "id": "asdf-asdf-asdf-asdf-0",
			 "timestamp": "2006-01-02T15:04:05Z",
			 "action": "${method}",
			 "target": {
			    "mediaType": "application/vnd.docker.distribution.manifest.v1+json",
			    "length": 1,
			    "digest": "sha256:fea8895f450959fa676bcc1df0611ea93823a735a01205fd8622846041d0c7cf",
			    "repository": "${image}",
			    "tag": "${tag}",
			    "url": "http://example.com/v2/library/test/manifests/sha256:c3b3692957d439ac1928219a83fac91e7bf96c153725526874673ae1f2023f8d5"
			 },
			 "request": {
			    "id": "asdfasdf",
			    "addr": "client.local",
			    "host": "registrycluster.local",
			    "method": "PUT",
			    "useragent": "test/0.1"
			 },
			 "actor": {
			    "name": "test-actor"
			 },
			 "source": {
			    "addr": "hostname.local:port"
			 }
		      }
		   ]
		}`)
	envelope = strings.Replace(envelope, "${method}", eventType, -1)
	envelope = strings.Replace(envelope, "${image}", image, -1)
	envelope = strings.Replace(envelope, "${tag}", tag, -1)

	fmt.Println(envelope)

	return strings.NewReader(envelope)
}