package main

import (
	"testing"
	"strings"

	"github.com/stretchr/testify/assert"
	"io"
)

func TestGetPushEventRepositories(t *testing.T) {

	repositories := GetPushEventRepositories(getEventEnvelope())
	assert.Equal(t, 1, len(repositories))
}

func TestGetPushEventRepositories_NoPushEvents(t *testing.T) {

	repositories := GetPushEventRepositories(getEventEnvelope_NoPushEvents())
	assert.Equal(t, 0, len(repositories))
}

func TestGetPushEventRepositories_EmptyEnvelop(t *testing.T) {

	repositories := GetPushEventRepositories(strings.NewReader("{}"))
	assert.Equal(t, 0, len(repositories))
}

func getEventEnvelope_NoPushEvents() io.Reader {

	return strings.NewReader(strings.TrimSpace(`
		{"events": [
		      {
			 "id": "asdf-asdf-asdf-asdf-0",
			 "timestamp": "2006-01-02T15:04:05Z",
			 "action": "pull",
			 "target": {
			    "mediaType": "application/vnd.docker.distribution.manifest.v1+json",
			    "length": 1,
			    "digest": "sha256:fea8895f450959fa676bcc1df0611ea93823a735a01205fd8622846041d0c7cf",
			    "repository": "library/test",
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
		      },
		      {
			 "id": "asdf-asdf-asdf-asdf-1",
			 "timestamp": "2006-01-02T15:04:05Z",
			 "action": "pull",
			 "target": {
			    "mediaType": "application/vnd.docker.container.image.rootfs.diff+x-gtar",
			    "length": 2,
			    "digest": "sha256:c3b3692957d439ac1928219a83fac91e7bf96c153725526874673ae1f2023f8d5",
			    "repository": "library/test",
			    "url": "http://example.com/v2/library/test/blobs/sha256:c3b3692957d439ac1928219a83fac91e7bf96c153725526874673ae1f2023f8d5"
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
		}`))
}

func getEventEnvelope() io.Reader {

	return strings.NewReader(strings.TrimSpace(`
		{"events": [
		      {
			 "id": "asdf-asdf-asdf-asdf-0",
			 "timestamp": "2006-01-02T15:04:05Z",
			 "action": "pull",
			 "target": {
			    "mediaType": "application/vnd.docker.distribution.manifest.v1+json",
			    "length": 1,
			    "digest": "sha256:fea8895f450959fa676bcc1df0611ea93823a735a01205fd8622846041d0c7cf",
			    "repository": "library/test",
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
		      },
		      {
			 "id": "asdf-asdf-asdf-asdf-1",
			 "timestamp": "2006-01-02T15:04:05Z",
			 "action": "push",
			 "target": {
			    "mediaType": "application/vnd.docker.container.image.rootfs.diff+x-gtar",
			    "length": 2,
			    "digest": "sha256:c3b3692957d439ac1928219a83fac91e7bf96c153725526874673ae1f2023f8d5",
			    "repository": "library/test",
			    "url": "http://example.com/v2/library/test/blobs/sha256:c3b3692957d439ac1928219a83fac91e7bf96c153725526874673ae1f2023f8d5"
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
		      },
		      {
			 "id": "asdf-asdf-asdf-asdf-2",
			 "timestamp": "2006-01-02T15:04:05Z",
			 "action": "pull",
			 "target": {
			    "mediaType": "application/vnd.docker.container.image.rootfs.diff+x-gtar",
			    "length": 3,
			    "digest": "sha256:c3b3692957d439ac1928219a83fac91e7bf96c153725526874673ae1f2023f8d5",
			    "repository": "library/test",
			    "url": "http://example.com/v2/library/test/blobs/sha256:c3b3692957d439ac1928219a83fac91e7bf96c153725526874673ae1f2023f8d5"
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
		}`))
}