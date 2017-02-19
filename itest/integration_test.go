package itest

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
)

func TestETE(t *testing.T) {

	const itest = "INTEGRATION_TESTS"
	if os.Getenv(itest) != "true" {
		t.Skipf("TestETE skiped. To run it, please, add enviroment varible: %s=true", itest)
	}

	res, err := SendEvent(Push, "nginx", "master-default-nginxdeployment-1.10")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}
