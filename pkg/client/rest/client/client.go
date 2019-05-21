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
	"strings"

	"github.com/emcecs/objectscale-management-go-sdk/pkg/client/model"
)

// Client is a REST client that communicates with the ObjectScale management API
type Client struct {
	// Username is the user name used to authenticate against the API
	Username string `json:"username"`

	// Password is the password used to authenticate against the API
	Password string `json:"password"`

	// Endpoint is the URL of the management API
	Endpoint string `json:"endpoint"`

	token       string
	HTTPClient  *http.Client
	authRetries int
}

func (c *Client) login() error {
	u, err := url.Parse(c.Endpoint)
	if err != nil {
		return err
	}
	u.Path = "/login"
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return err
	}
	req.SetBasicAuth(c.Username, c.Password)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	if err = handleResponse(resp); err != nil {
		return err
	}
	c.token = resp.Header.Get("X-SDS-AUTH-TOKEN")
	if c.token == "" {
		return fmt.Errorf("server error: login failed")
	} else {
		c.authRetries = 0
	}
	return nil
}

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
			if  err != nil {
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
	return c.token != ""
}

// MakeRemoteCall executes an API request against the client endpoint
func (c *Client) MakeRemoteCall(r Request) (*http.Response, error) {
	var (
		obj []byte
		err error
		q   = url.Values{}
	)
	switch r.ContentType {
	case ContentTypeXML:
		obj, err = xml.Marshal(r.Body)
	case ContentTypeJSON:
		obj, err = json.Marshal(r.Body)
	default:
		return nil, errors.New("invalid content-type")
	}
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(c.Endpoint)
	if err != nil {
		return nil, err
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
		return nil, err
	}
	if !c.isLoggedIn() {
		if err = c.login(); err != nil {
			return nil, err
		}
	}
	req.Header.Add("Content-Type", r.ContentType)
	req.Header.Add("Accept", "application/xml")
	req.Header.Add("X-SDS-AUTH-TOKEN", c.token)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 401 {
		if c.authRetries < AuthRetriesMax {
			c.authRetries += 1
			c.token = ""
			return c.MakeRemoteCall(r)
		} else {
			return nil, errors.New(strings.ToLower(resp.Status))
		}
	}
	return resp, nil
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

