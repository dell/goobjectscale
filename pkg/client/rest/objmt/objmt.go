package objmt

import (
	"fmt"
	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/model"
	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/rest/client"
	"net/http"
)

// Objmt is a REST implementation of the Objmt interface
type Objmt struct {
	Client *client.Client
}

func (o *Objmt) GetAccountBillingInfo(ids []string, params map[string]string) (*model.AccountBillingInfoList, error) {
	// TODO prepare request body with IDs
	req := client.Request{
		Method:      http.MethodPost,
		Path:        fmt.Sprintf("/object/mt/account/info"),
		ContentType: client.ContentTypeXML,
		//Body:		 &body,
		Params: params,
	}
	ret := &model.AccountBillingInfoList{}
	err := o.Client.MakeRemoteCall(req, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (o *Objmt) GetAccountBillingSample(ids []string, params map[string]string) (*model.AccountBillingSampleList, error) {
	// TODO prepare request body with IDs
	req := client.Request{
		Method:      http.MethodPost,
		Path:        fmt.Sprintf("/object/mt/account/sample"),
		ContentType: client.ContentTypeXML,
		//Body:		 &body,
		Params: params,
	}
	ret := &model.AccountBillingSampleList{}
	err := o.Client.MakeRemoteCall(req, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (o *Objmt) GetBucketBillingInfo(account string, ids []string, params map[string]string) (*model.BucketBillingInfoList, error) {
	// TODO prepare request body with IDs
	req := client.Request{
		Method:      http.MethodPost,
		Path:        fmt.Sprintf("/object/mt/account/%s/bucket/info", account),
		ContentType: client.ContentTypeXML,
		//Body:		 &body,
		Params: params,
	}
	ret := &model.BucketBillingInfoList{}
	err := o.Client.MakeRemoteCall(req, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (o *Objmt) GetBucketBillingSample(account string, ids []string, params map[string]string) (*model.BucketBillingSampleList, error) {
	// TODO prepare request body with IDs
	req := client.Request{
		Method:      http.MethodPost,
		Path:        fmt.Sprintf("/object/mt/account/%s/bucket/sample", account),
		ContentType: client.ContentTypeXML,
		//Body:		 &body,
		Params: params,
	}
	ret := &model.BucketBillingSampleList{}
	err := o.Client.MakeRemoteCall(req, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (o *Objmt) GetBucketBillingPerf(account string, ids []string, params map[string]string) (*model.BucketPerfDataList, error) {
	// TODO prepare request body with IDs
	req := client.Request{
		Method:      http.MethodPost,
		Path:        fmt.Sprintf("/object/mt/account/%s/bucket/perf", account),
		ContentType: client.ContentTypeXML,
		//Body:		 &body,
		Params: params,
	}
	ret := &model.BucketPerfDataList{}
	err := o.Client.MakeRemoteCall(req, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (o *Objmt) GetReplicationInfo(account string, replicationPairs [][]string, params map[string]string) (*model.BucketReplicationInfoList, error) {
	// TODO prepare request body with IDs
	req := client.Request{
		Method:      http.MethodPost,
		Path:        fmt.Sprintf("/object/mt/account/%s/replication/info", account),
		ContentType: client.ContentTypeXML,
		//Body:		 &body,
		Params: params,
	}
	ret := &model.BucketReplicationInfoList{}
	err := o.Client.MakeRemoteCall(req, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (o *Objmt) GetReplicationSample(account string, replicationPairs [][]string, params map[string]string) (*model.BucketReplicationSampleList, error) {
	// TODO prepare request body with IDs
	req := client.Request{
		Method:      http.MethodPost,
		Path:        fmt.Sprintf("/object/mt/account/%s/replication/sample", account),
		ContentType: client.ContentTypeXML,
		//Body:		 &body,
		Params: params,
	}
	ret := &model.BucketReplicationSampleList{}
	err := o.Client.MakeRemoteCall(req, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (o *Objmt) GetStoreBillingInfo(params map[string]string) (*model.StoreBillingInfoList, error) {
	req := client.Request{
		Method:      http.MethodGet,
		Path:        fmt.Sprintf("/object/mt/store/info"),
		ContentType: client.ContentTypeXML,
		Params:      params,
	}
	ret := &model.StoreBillingInfoList{}
	err := o.Client.MakeRemoteCall(req, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (o *Objmt) GetStoreBillingSample(params map[string]string) (*model.StoreBillingSampleList, error) {
	req := client.Request{
		Method:      http.MethodGet,
		Path:        fmt.Sprintf("/object/mt/store/sample"),
		ContentType: client.ContentTypeXML,
		Params:      params,
	}
	ret := &model.StoreBillingSampleList{}
	err := o.Client.MakeRemoteCall(req, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (o *Objmt) GetStoreReplicationData(ids []string, params map[string]string) (*model.StoreReplicationDataList, error) {
	// TODO prepare request body with IDs
	req := client.Request{
		Method:      http.MethodPost,
		Path:        fmt.Sprintf("/object/mt/store/replication"),
		ContentType: client.ContentTypeXML,
		//Body:		 &body,
		Params: params,
	}
	ret := &model.StoreReplicationDataList{}
	err := o.Client.MakeRemoteCall(req, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
