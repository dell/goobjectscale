package testutils

import "net/http"

// IncrCapture increments the integer value of a capture field in the map
func IncrCapture(captures map[string]interface{}, name string) {
	currentVal, ok := captures[name].(int)
	if !ok || captures[name] == nil {
		captures[name] = 0
	}
	captures[name] = currentVal + 1
}

// RoundTripFunc is a transport mock that makes a fake HTTP response locally
type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip mocks an http request and returns an http response
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

//NewTestClient returns *http.Client with Transport replaced to avoid making real calls
func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}
