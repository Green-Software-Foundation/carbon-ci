package watttime

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

//INITIALIZE AND RETURN NEW REQUEST
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

// INITIALIZE NEW REQUEST
func httpRequest(headers httpRequestType) error {
	client := &http.Client{}

	req, err := initRequest(headers.Method, headers.Url, headers.Data)

	if err != nil {
		return err
	}

	// SET HEADERS/S
	for h := range headers.Header {
		req.Header.Add(h, headers.Header[h])
	}

	// SET QUERY STRING
	query := req.URL.Query()
	for q := range headers.Query {
		query.Add(q, headers.Query[q])
	}
	req.URL.RawQuery = query.Encode()

	// DO REQUEST
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	// EXECUTE AFTER FUNCTION RETURNS
	defer resp.Body.Close()

	// GET RESPONSE DATA
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// CONVERT RESPONSE DATA FROM BYTE ARRAY TO OBJECT
	err = json.Unmarshal(bodyBytes, &headers.Response)
	if err != nil {
		return errors.New("===== HTTP RESPONSE =====" + "\n" + string(bodyBytes))
	}

	return nil
}
