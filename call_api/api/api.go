package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// Get ...
func Get(url string, data, headers map[string]interface{}) (ApiResponse, error) {

	// create request
	req, _ := http.NewRequest("GET", url, nil)

	// create query
	query := req.URL.Query()
	for key, value := range data {
		query.Add(key, fmt.Sprintf("%v", value))
	}

	req.URL.RawQuery = query.Encode()

	for key, value := range headers {
		req.Header.Set(key, fmt.Sprintf("%v", value))
	}
	// call request
	timeout := time.Duration(30 * time.Second)
	client := http.Client{Timeout: timeout}
	res, err := client.Do(req)
	response := ApiResponse{}
	if err != nil {
		return response, err
	}
	defer res.Body.Close()
	response.Status = res.Status
	response.StatusCode = res.StatusCode
	response.Body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return response, err
	}
	return response, nil
}

func Post(url string, data, headers map[string]interface{}) (ApiResponse, error) {
	response := ApiResponse{}
	method := "POST"
	payloadStr, _ := json.Marshal(data)
	payload := strings.NewReader(string(payloadStr))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return response, err
	}

	for key, val := range headers {
		req.Header.Add(key, fmt.Sprintf("%v", val))
	}

	res, err := client.Do(req)
	if err != nil {
		return response, err
	}
	defer res.Body.Close()
	response.Status = res.Status
	response.StatusCode = res.StatusCode
	response.Body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return response, err
	}
	return response, nil
}

func Put(url string, data, headers map[string]interface{}) (ApiResponse, error) {
	response := ApiResponse{}
	method := "PUT"
	payloadStr, _ := json.Marshal(data)
	payload := strings.NewReader(string(payloadStr))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return response, err
	}

	for key, val := range headers {
		req.Header.Add(key, fmt.Sprintf("%v", val))
	}

	if username, ok := headers["Username"]; ok {
		if password, ok := headers["Password"]; ok {
			req.SetBasicAuth(username.(string), password.(string))
		}
	}

	res, err := client.Do(req)
	if err != nil {
		return response, err
	}
	defer res.Body.Close()
	response.Status = res.Status
	response.StatusCode = res.StatusCode
	response.Body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return response, err
	}
	return response, nil
}
