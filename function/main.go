package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type RequestOption func(req *http.Request)

type RequestHeader struct {
	SendDate          string `json:"sendDate"`          // Date and time the request was sent.
	SeqNo             string `json:"seqNo"`             // Specify a number that is unique for all requests in the same contract ID.
	OrgID             string `json:"orgId"`             // ID representing the contract, - (fixed value)
	APIGroupID        string `json:"apiGroupId"`        // ID representing a group of API users, (fixed value).
	APIClientID       string `json:"apiClientId"`       // ID of the API user, set to a value paid out by the subscriber in advance (fixed value).
	APIClientPassword string `json:"apiClientPassword"` // Password for API user. Set the value paid out by the subscriber in advance (fixed value).
	DataCount         int    `json:"dataCount"`         // Number of data to be processed.
}

func main() {

	// request := http.Request{}

	// header := RequestHeader{
	// 	OrgID: "4eachtech",
	// }


}

func (s *RequestHeader) toHTTPHeader() RequestOption {
	data, err := json.Marshal(s)
	if err != nil {
		log.Fatalln("marshal header: %w", err)
	}

	m := make(map[string]interface{})

	err = json.Unmarshal(data, &m)
	if err != nil {
		log.Fatalln("unmarshal header: %w", err)
	}

	return func(req *http.Request) {
		for k, v := range m {
			req.Header.Set(k, fmt.Sprintf("%v", v))
		}
	}
}
