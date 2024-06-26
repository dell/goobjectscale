// Copyright © 2023 Dell Inc. or its subsidiaries. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//      http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package objmt

import (
	"context"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/dell/goobjectscale/pkg/client/model"
	"github.com/dell/goobjectscale/pkg/client/rest/client"
)

type bucketIDsReqBody struct {
	XMLName xml.Name `xml:"bucket_list"`
	IDs     []string `xml:"id"`
}

type accountIDsReqBody struct {
	XMLName xml.Name `xml:"account_list"`
	IDs     []string `xml:"id"`
}

type storeIDsReqBody struct {
	XMLName xml.Name `xml:"store_list"`
	IDs     []string `xml:"id"`
}

type replicationPairsReqBody struct {
	XMLName      xml.Name         `xml:"replication_list"`
	Replications []replicationIDs `xml:"replication"`
}

type replicationIDs struct {
	XMLName xml.Name `xml:"replication"`
	Src     string   `xml:"src"`
	Dest    string   `xml:"dest"`
}

func newReplicationIDs(ids [][]string) *replicationPairsReqBody {
	ret := &replicationPairsReqBody{}

	ret.Replications = []replicationIDs{}
	for _, id := range ids {
		ret.Replications = append(ret.Replications, replicationIDs{Src: id[0], Dest: id[1]})
	}

	return ret
}

// Objmt is a REST implementation of the Objmt interface.
type Objmt struct {
	Client client.RemoteCaller
}

// GetAccountBillingInfo returns billing info metrics for defined accounts.
func (o *Objmt) GetAccountBillingInfo(ctx context.Context, ids []string, params map[string]string) (*model.AccountBillingInfoList, error) {
	req := client.Request{
		Method:      http.MethodPost,
		Path:        "/object/mt/account/info",
		ContentType: client.ContentTypeXML,
		Body:        accountIDsReqBody{IDs: ids},
		Params:      params,
	}

	ret := &model.AccountBillingInfoList{}

	err := o.Client.MakeRemoteCall(ctx, req, ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetAccountBillingSample returns billing sample (time-window) metrics for defined accounts.
func (o *Objmt) GetAccountBillingSample(ctx context.Context, ids []string, params map[string]string) (*model.AccountBillingSampleList, error) {
	req := client.Request{
		Method:      http.MethodPost,
		Path:        "/object/mt/account/sample",
		ContentType: client.ContentTypeXML,
		Body:        accountIDsReqBody{IDs: ids},
		Params:      params,
	}

	ret := &model.AccountBillingSampleList{}

	err := o.Client.MakeRemoteCall(ctx, req, ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetBucketBillingInfo returns billing info metrics for defined buckets and account.
func (o *Objmt) GetBucketBillingInfo(ctx context.Context, account string, ids []string, params map[string]string) (*model.BucketBillingInfoList, error) {
	// TODO prepare request body with IDs
	req := client.Request{
		Method:      http.MethodPost,
		Path:        fmt.Sprintf("/object/mt/account/%s/bucket/info", account),
		ContentType: client.ContentTypeXML,
		Body:        bucketIDsReqBody{IDs: ids},
		Params:      params,
	}

	ret := &model.BucketBillingInfoList{}

	err := o.Client.MakeRemoteCall(ctx, req, ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetBucketBillingSample returns billing sample (time-window) metrics for defined buckets and account.
func (o *Objmt) GetBucketBillingSample(ctx context.Context, account string, ids []string, params map[string]string) (*model.BucketBillingSampleList, error) {
	req := client.Request{
		Method:      http.MethodPost,
		Path:        fmt.Sprintf("/object/mt/account/%s/bucket/sample", account),
		ContentType: client.ContentTypeXML,
		Body:        bucketIDsReqBody{IDs: ids},
		Params:      params,
	}

	ret := &model.BucketBillingSampleList{}

	err := o.Client.MakeRemoteCall(ctx, req, ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetBucketBillingPerf returns performance metrics for defined buckets and account.
func (o *Objmt) GetBucketBillingPerf(ctx context.Context, account string, ids []string, params map[string]string) (*model.BucketPerfDataList, error) {
	req := client.Request{
		Method:      http.MethodPost,
		Path:        fmt.Sprintf("/object/mt/account/%s/bucket/perf", account),
		ContentType: client.ContentTypeXML,
		Body:        bucketIDsReqBody{IDs: ids},
		Params:      params,
	}

	ret := &model.BucketPerfDataList{}

	err := o.Client.MakeRemoteCall(ctx, req, ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetReplicationInfo returns billing info metrics for defined replication pairs and account.
func (o *Objmt) GetReplicationInfo(ctx context.Context, account string, replicationPairs [][]string, params map[string]string) (*model.BucketReplicationInfoList, error) {
	req := client.Request{
		Method:      http.MethodPost,
		Path:        fmt.Sprintf("/object/mt/account/%s/replication/info", account),
		ContentType: client.ContentTypeXML,
		Body:        newReplicationIDs(replicationPairs),
		Params:      params,
	}

	ret := &model.BucketReplicationInfoList{}

	err := o.Client.MakeRemoteCall(ctx, req, ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetReplicationSample returns billing sample (time-window) metrics for defined replication pairs and account.
func (o *Objmt) GetReplicationSample(ctx context.Context, account string, replicationPairs [][]string, params map[string]string) (*model.BucketReplicationSampleList, error) {
	req := client.Request{
		Method:      http.MethodPost,
		Path:        fmt.Sprintf("/object/mt/account/%s/replication/sample", account),
		ContentType: client.ContentTypeXML,
		Body:        newReplicationIDs(replicationPairs),
		Params:      params,
	}

	ret := &model.BucketReplicationSampleList{}

	err := o.Client.MakeRemoteCall(ctx, req, ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetStoreBillingInfo returns billing info metrics for object store.
func (o *Objmt) GetStoreBillingInfo(ctx context.Context, params map[string]string) (*model.StoreBillingInfoList, error) {
	req := client.Request{
		Method:      http.MethodGet,
		Path:        "/object/mt/store/info",
		ContentType: client.ContentTypeXML,
		Params:      params,
	}

	ret := &model.StoreBillingInfoList{}

	err := o.Client.MakeRemoteCall(ctx, req, ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetStoreBillingSample returns billing sample (time-window) metrics for object store.
func (o *Objmt) GetStoreBillingSample(ctx context.Context, params map[string]string) (*model.StoreBillingSampleList, error) {
	req := client.Request{
		Method:      http.MethodGet,
		Path:        "/object/mt/store/sample",
		ContentType: client.ContentTypeXML,
		Params:      params,
	}

	ret := &model.StoreBillingSampleList{}

	err := o.Client.MakeRemoteCall(ctx, req, ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetStoreReplicationData returns CRR metrics for defined object stores.
func (o *Objmt) GetStoreReplicationData(ctx context.Context, ids []string, params map[string]string) (*model.StoreReplicationDataList, error) {
	req := client.Request{
		Method:      http.MethodPost,
		Path:        "/object/mt/store/replication",
		ContentType: client.ContentTypeXML,
		Body:        storeIDsReqBody{IDs: ids},
		Params:      params,
	}

	ret := &model.StoreReplicationDataList{}

	err := o.Client.MakeRemoteCall(ctx, req, ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
