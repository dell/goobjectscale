package client

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/model"
)

// Client is a REST client that communicates with the ObjectScale management API
type Client struct {
	// Endpoint is the object control service endpoint
	Endpoint string `json:"endpoint"`

	// Token is OS token returned from fed service API interactions
	Token string `json:"token"`

	HTTPClient *http.Client

	// Should X-EMC-Override header be added into the request
	OverrideHeader bool
}

// MakeRemoteCall executes an API request against the client endpoint, returning
// the object body of the response into a response object
func (c *Client) MakeRemoteCall(r Request, into interface{}) error {
	var (
		obj []byte
		err error
		q   = url.Values{}
	)
	switch r.ContentType {
	case ContentTypeXML:
		obj, err = xml.Marshal(r.Body)
	case ContentTypeJSON:
		if raw, ok := r.Body.(json.RawMessage); ok {
			obj, err = raw.MarshalJSON()
		} else {
			obj, err = json.Marshal(r.Body)
		}
	default:
		return errors.New("invalid content-type")
	}
	if err != nil {
		return err
	}
	u, err := url.Parse(c.Endpoint)
	if err != nil {
		return err
	}
	u.Path = r.Path
	if r.Params != nil {
		for key, value := range r.Params {
			q.Add(key, value)
		}
	}
	u.RawQuery = q.Encode()
	req, err := http.NewRequest(r.Method, u.String(), bytes.NewBuffer(obj))
	if err != nil {
		return err
	}
	req.Header.Add("Accept", r.ContentType)
	req.Header.Add("Content-Type", r.ContentType)
	req.Header.Add("Accept", "application/xml")
	req.Header.Add("X-SDS-AUTH-TOKEN", c.Token)
	if c.OverrideHeader {
		req.Header.Add("X-EMC-Override", "true")
	}
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	var body []byte
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return err
	}

	switch {
	case resp.StatusCode > 399:
		ecsError := &model.Error{}
		switch r.ContentType {
		case ContentTypeXML:
			if err = xml.Unmarshal(body, ecsError); err != nil {
				return err
			}
		case ContentTypeJSON:
			if err = json.Unmarshal(body, ecsError); err != nil {
				return err
			}
		}
		return fmt.Errorf("%s: %s", ecsError.Description, ecsError.Details)
	case into == nil:
		// No errors found, and no response object defined, so just return
		// without error
		return nil
	default:
		if len(body) == 0 {
			return nil
		}
		switch r.ContentType {
		case ContentTypeXML:
			if err = xml.Unmarshal(body, into); err != nil {
				return err
			}
		case ContentTypeJSON:
			if err = json.Unmarshal(body, into); err != nil {
				return err
			}
		}
	}
	return nil
}

const (
	// ContentTypeXML
	ContentTypeXML = "application/xml"

	// ContentTypeJSON
	ContentTypeJSON = "application/json"
)

// Request is an ObjectScale API request wrapper
type Request struct {
	// Method is the method of REST API request
	Method string

	// Path is the path of REST API request
	Path string

	// Body is the body of REST API request
	Body interface{}

	// ContentType is the body of REST API request
	ContentType string

	// Params are the parameters of the REST API request
	Params map[string]string
}
