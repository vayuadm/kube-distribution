package dockerhub

import (
	"strings"
	"testing"

	"io"

	"github.com/stretchr/testify/assert"
)

func TestEnvelope_GetPushEventRepositories(t *testing.T) {

	repository, err := GetPushEventRepositories(getEventEnvelope())
	assert.NoError(t, err)
	assert.NotNil(t, repository)
	assert.Equal(t, "vayuadm/kube-distribution", repository.Name)
	assert.Equal(t, "master--default--ceribrouideplyment--7", repository.Tag)
}

func TestEnvelope_GetPushEventRepositories_EmptyEnvelope(t *testing.T) {

	repository, err := GetPushEventRepositories(strings.NewReader("{}"))
	assert.NoError(t, err)
	assert.Nil(t, repository)
}

func TestEnvelope_GetPushEventRepositories_CorruptedEnvelope(t *testing.T) {

	_, err := GetPushEventRepositories(strings.NewReader("{"))
	assert.Error(t, err)
}

func getEventEnvelope() io.Reader {

	return strings.NewReader(strings.TrimSpace(`
				{
		  "push_data": {
		    "pushed_at": 1494748295,
		    "images": [],
		    "tag": "master--default--ceribrouideplyment--7",
		    "pusher": "effoeffi"
		  },
		  "callback_url": "https://registry.hub.docker.com/u/vayuadm/kube-distribution/hook/25i05b0gidb0j4gg4dbbe1g2hfhfi13i1/",
		  "repository": {
		    "status": "Active",
		    "description": "",
		    "is_trusted": false,
		    "full_description": null,
		    "repo_url": "https://hub.docker.com/r/vayuadm/kube-distribution",
		    "owner": "vayuadm",
		    "is_official": false,
		    "is_private": false,
		    "name": "kube-distribution",
		    "namespace": "vayuadm",
		    "star_count": 0,
		    "comment_count": 0,
		    "date_created": 1488206469,
		    "repo_name": "vayuadm/kube-distribution"
		  }
		}`))
}
