package main

import (
	"app/call_api/api"
	"app/call_api/model"
	"encoding/json"
	"fmt"
)

func main() {
	url := "http://localhost:9090"
	resp, err := api.Get(url, nil, nil)
	if err != nil {
		panic(err)
	}
	var students []model.Student
	// fmt.Println(string(resp.Body))
	json.Unmarshal(resp.Body, &students)
	fmt.Printf("students: %+v", students)

}
