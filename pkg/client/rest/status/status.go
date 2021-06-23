package status

import (
	"net/http"

	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/model"
	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/rest/client"
)

// Status is a REST implementation of the Status interface
type Status struct {
	Client *client.Client
}

// GetRebuildStatus implements the status interface
func (b *Status) GetRebuildStatus(objStoreName, ssPodName, ssPodNameSpace, level string, params map[string]string) (*model.RebuildInfo, error) {
	requestURL := "vdc/recovery-status/devices/" + ssPodName + "." +
		objStoreName + "-ss." + ssPodNameSpace + ".svc.cluster.local/levels/" + level
	req := client.Request{
		Method:      http.MethodGet,
		Path:        requestURL,
		ContentType: client.ContentTypeJSON,
		Params:      params,
	}
	rebuildInfo := &model.RebuildInfo{}
	err := b.Client.MakeRemoteCall(req, rebuildInfo)
	if err != nil {
		return nil, err
	}

	return rebuildInfo, nil
}
