package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Student :
type Student struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	ClassName string `json:"class_name"`
	Age       int    `json:"age"`
}

func main() {
	callAPI3()
}

func callAPI1() {
	res, err := http.Get("http://localhost:9090")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	bs, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bs))
}

func callAPI2() {
	res, err := http.Get("http://localhost:9090")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	bs, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var students []Student

	err = json.Unmarshal(bs, &students)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("result:%+v", students)

}

func callAPI3() {
	res, err := http.Get("http://localhost:9090")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	var students []Student

	decoder := json.NewDecoder(res.Body)
	decoder.Decode(&students)
	fmt.Printf("result3: %+v", students)

}
