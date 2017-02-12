package main

import (
	"testing"
	"strings"

	"github.com/stretchr/testify/assert"
	"io"
)

func TestGetPushEventRepositories(t *testing.T) {

	repositories := GetPushEventRepositories(getPushEventEnvelope())
	assert.Equal(t, 1, len(repositories))
}

func getPushEventEnvelope() io.Reader {

	return strings.NewReader(strings.TrimSpace(`{"events": [{
		"id": "asdf-asdf-asdf-asdf-0",
		"timestamp": "2006-01-02T15:04:05Z",
		"action": "push",
		"target": {
			"mediaType": "application/vnd.docker.distribution.manifest.v1+json",
			"size": 1,
			"digest": "sha256:0123456789abcdef0",
			"length": 1,
			"repository": "ceribro",
			"url": "http://example.com/v2/library/test/manifests/latest"
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
	}]}`))
}