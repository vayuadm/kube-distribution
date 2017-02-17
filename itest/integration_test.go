package itest

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"net/http"
)

func TestETE(t *testing.T) {

	res, err := SendEvent(Push, "nginx", "master-default-nginxdeployment-1.10")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}
