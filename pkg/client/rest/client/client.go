package client

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/model"
)

// Client is a REST client that communicates with the ObjectScale management API
type Client struct {
	// ObjectStoreID is used for generating TOPT for retrieving the OS
	// TOKEN to be used when accessing the object store's management API
	ObjectStoreID string `json:"objectStoreID"`

	// Endpoint is the URL of the management API
	Endpoint string `json:"endpoint"`

	// Gateway is the auth endpoint
	Gateway string `json:"gateway"`

	// Namespace is the K8S namespace where client is running
	Namespace string `json:"namespace"`

	// PodName is the K8S pod name where client is running
	PodName string `json:"podName"`

	// Token is OS token returned from fed service API interactions
	Token string `json:"token"`

	objectScaleID string
	HTTPClient    *http.Client
	authRetries   int

	// Should X-EMC-Override header be added into the request
	OverrideHeader bool
}

func (c *Client) getObjectScaleID() error {
	u, err := url.Parse(c.Gateway)
	if err != nil {
		return err
	}
	u.Path = "/fedsvc/objectScaleId"
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return err
	}
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	if err = handleResponse(resp); err != nil {
		return err
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("cannot read response, returned by GET /fedsvc/objectScaleId %w", err)
	}
	c.objectScaleID = string(bodyBytes)
	if c.objectScaleID == "" {
		return fmt.Errorf("server error: unable to get object scale id")
	}
	return nil
}

func (c *Client) login() error {
	if err := c.getObjectScaleID(); err != nil {
		return err
	}
	u, err := url.Parse(c.Gateway)
	if err != nil {
		return err
	}
	u.Path = "/mgmt/serviceLogin"
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return err
	}

	// urn:osc:{ObjectScaleId}:{ObjectStoreId}:service/{ServiceNameId}
	serviceUrn := fmt.Sprintf("urn:osc:%s:%s:service/%s", c.objectScaleID, "", c.PodName)

	// B64-{ObjectScaleId},{ObjectStoreId},{ServiceK8SNamespace},{ServiceNameId}
	userName := fmt.Sprintf("%s,%s,%s,%s", c.objectScaleID, "", c.Namespace, c.PodName)
	userName = base64.StdEncoding.EncodeToString([]byte(userName))
	userName = "B64-" + userName

	fmt.Println("Username: " + userName)

	// HMACSHA256(key, ServiceUrn + Time_factor)
	// time_factor = currentTimeInSeconds (rounded to nearest 30 seconds)
	timeFactor := time.Now().UTC().Round(30 * time.Second).Unix()

	fmt.Println(timeFactor)

	//secret := "p55GJt5RmSiVU2IyCmv6qtrpJzoqsgMJ9M8G8MyPaNJGpk5w5kDH5HzsBokdKsVuhaQBaWNfn45JLrzavA5e5SdUzxnwiCFyGoQtgAsS0RoEbo2jU5JMPcOK47jhC1Yl"
	secret := "Bq0Qy9giIopNApzNaQvxzVIhW4Nh0rdI5PvkZksXzMJZnXknk18VAvB28u1QVEAHNJrervGKqY6Fq5OJMXF88VOT0YgPhgWL18EHzaEMy9mIwPRVatxzfHouSBYHFZ5x"
	data := serviceUrn + strconv.FormatInt(timeFactor, 10)

	fmt.Println("Data to be hashed: " + data)

	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	//password := hex.EncodeToString(h.Sum(nil))
	password := base64.StdEncoding.EncodeToString(h.Sum(nil))
	password = "HARDCODED-" + password

	fmt.Println("Hash Result: " + password)

	req.SetBasicAuth(userName, password)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	if err = handleResponse(resp); err != nil {
		return err
	}
	c.Token = resp.Header.Get("X-SDS-AUTH-TOKEN")
	if c.Token == "" {
		return fmt.Errorf("server error: login failed")
	} else {
		c.authRetries = 0
	}
	return nil
}

// func (c *Client) login() error {
// 	u, err := url.Parse(c.Gateway)
// 	if err != nil {
// 		return err
// 	}
// 	u.Path = "/mgmt/login"
// 	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
// 	if err != nil {
// 		return err
// 	}
// 	req.SetBasicAuth(c.Username, c.Password)
// 	resp, err := c.HTTPClient.Do(req)
// 	if err != nil {
// 		return err
// 	}
// 	if err = handleResponse(resp); err != nil {
// 		return err
// 	}
// 	c.token = resp.Header.Get("X-SDS-AUTH-TOKEN")
// 	if c.token == "" {
// 		return fmt.Errorf("server error: login failed")
// 	} else {
// 		c.authRetries = 0
// 	}
// 	return nil
// }

func handleResponse(resp *http.Response) error {
	if resp.StatusCode > 399 {
		switch resp.Body {
		case nil:
			switch {
			case resp.Status != "":
				return fmt.Errorf("server error: %s", strings.ToLower(resp.Status))
			case resp.StatusCode != 0:
				return fmt.Errorf("server error: status code %d", resp.StatusCode)
			}
		default:
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Printf("server errror: %s", strings.ToLower(resp.Status))
				return err
			}
			apiError := &model.Error{}
			err = xml.Unmarshal(body, apiError)
			if err != nil {
				return errors.New(string(body))
			}
			switch {
			case apiError.Code == 1004:
				return errors.New("server error: not found")
			default:
				return fmt.Errorf("server error: %s", strings.ToLower(apiError.Description))
			}
		}
	}
	return nil
}

func (c *Client) isLoggedIn() bool {
	return c.Token != ""
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
	if !c.isLoggedIn() {
		if err = c.login(); err != nil {
			return err
		}
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
	case resp.StatusCode == http.StatusUnauthorized:
		if c.authRetries < AuthRetriesMax {
			c.authRetries += 1
			c.Token = ""
			return c.MakeRemoteCall(r, into)
		}
		return errors.New(strings.ToLower(resp.Status))
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
	// AuthRetriesMax is the maximum number of times the client will attempt to
	// login before returning an error
	AuthRetriesMax = 3

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
