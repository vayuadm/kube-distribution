package itest

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func SendEvent(image, tag string) (response *http.Response, err error) {

	return http.Post("http://localhost:5050/events",
		"",
		getEventEnvelope(image, tag))
}

func getEventEnvelope(image, tag string) io.Reader {

	envelope := strings.TrimSpace(`
		{[{
		  "push_data": {
		    "pushed_at": 1494748295,
		    "images": [],
		    "tag": "${tag}",
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
		    "repo_name": "${image}"
		  }
		}]
		}`)
	envelope = strings.Replace(envelope, "${image}", image, -1)
	envelope = strings.Replace(envelope, "${tag}", tag, -1)

	fmt.Println(envelope)

	return strings.NewReader(envelope)
}
