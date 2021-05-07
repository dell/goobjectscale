package crr

import (
	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/rest/client"
	"net/http"
	"path"
	"strconv"
)

// CRR is a REST implementation of the CRR interface
type CRR struct {
	Client *client.Client
}

// PauseReplication implements the CRR interface
func (c *CRR) PauseReplication(destObjectScale string, destObjectStore string, durationMills int, params map[string]string) error {
	req := client.Request{
		Method:      http.MethodPost,
		Path:        path.Join("replication", "control", destObjectScale, destObjectStore, "pause", strconv.Itoa(durationMills)),
		ContentType: client.ContentTypeXML,
		Params:      params,
	}
	return c.Client.MakeRemoteCall(req, nil)
}
