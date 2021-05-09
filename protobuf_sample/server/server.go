package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"app/protobuf_sample/proto/echo"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/mux"
)

func Echo(resp http.ResponseWriter, req *http.Request) {
	contentLength := req.ContentLength
	fmt.Printf("Content Length Received: %v\n", contentLength)

	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalf("Unable to read message from request: %v", err)
	}
	request := &echo.EchoRequest{}

	proto.Unmarshal(data, request)
	name := request.GetName()

	result := &echo.EchoResponse{Message: "Hello " + name}
	response, err := proto.Marshal(result)
	if err != nil {
		log.Fatalf("Unable to marshal response: %v", err)
	}
	resp.Write(response)
}

func main() {
	fmt.Println("Start the api server...")
	r := mux.NewRouter()
	r.HandleFunc("/echo", Echo).Methods("POST")

	server := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 2 * time.Second,
		ReadTimeout:  2 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
