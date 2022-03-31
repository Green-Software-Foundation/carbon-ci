package watttime

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// initRequest initialize and return new request
func initRequest(method string, url string, data map[string]string) (*http.Request, error) {
	if len(method) == 0 {
		return http.NewRequest(method, url, nil)
	} else {
		json, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		return http.NewRequest(method, url, bytes.NewBuffer(json))
	}
}

// httpRequest initialize new request
func httpRequest(headers httpRequestType) error {
	client := &http.Client{}

	req, err := initRequest(headers.Method, headers.Url, headers.Data)

	if err != nil {
		return err
	}

	// Set headers
	for h := range headers.Header {
		req.Header.Add(h, headers.Header[h])
	}

	// Set query string
	query := req.URL.Query()
	for q := range headers.Query {
		query.Add(q, headers.Query[q])
	}
	req.URL.RawQuery = query.Encode()

	// Do request
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	// Execute after function returns
	defer resp.Body.Close()

	// Get response data type byte array
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Convert response data from byte array to object or to string if error is not nil.
	err = json.Unmarshal(bodyBytes, &headers.Response)
	if err != nil {
		return errors.New("===== HTTP RESPONSE =====" + "\n" + string(bodyBytes))
	}

	return nil
}
