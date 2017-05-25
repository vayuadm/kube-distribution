package itest

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestETE(t *testing.T) {

	const iTest = "INTEGRATION_TESTS"
	if os.Getenv(iTest) != "true" {
		t.Skipf("TestETE skipped. To run it, please, add environment variable: %s=true", iTest)
	}

	res, err := SendEvent("nginx", "master--default--nginxdeployment--1.10")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}
