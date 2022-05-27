package http

import (
	"bytes"
	"encoding/json"
	"errors"

	"io/ioutil"
	"net/http"
)

// Request represents the http request header
type Request struct {
	Url      string
	Method   string
	Data     map[string]string
	Header   map[string]string
	Query    map[string]string
	Response interface{}
}

// init initialize and return new request
func (request *Request) init(method string, url string, data map[string]string) (*http.Request, error) {
	if len(data) == 0 {
		return http.NewRequest(method, url, nil)
	} else {
		json, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		return http.NewRequest(method, url, bytes.NewBuffer(json))
	}
}

// Send initialize new request
func (request *Request) Send() error {
	client := &http.Client{}

	req, err := request.init(request.Method, request.Url, request.Data)

	if err != nil {
		return err
	}

	// Set headers
	for h := range request.Header {
		req.Header.Add(h, request.Header[h])
	}

	// // Set query string
	query := req.URL.Query()
	for q := range request.Query {
		query.Add(q, request.Query[q])
	}
	req.URL.RawQuery = query.Encode()

	// Do request
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	// Execute after function returns
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}

	// Get response data type byte array
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Convert response data from byte array to object or to string if error is not nil.
	err = json.Unmarshal(bodyBytes, &request.Response)

	if err != nil {
		return err
	}

	return nil
}
