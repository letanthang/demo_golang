package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handleFunc)
	time.AfterFunc(time.Millisecond*100, func() {
		fmt.Println("Server listen at 9090")
	})

	log.Fatalf("error: %v", http.ListenAndServe(":9090", nil))
}

func handleFunc(w http.ResponseWriter, r *http.Request) {
	select {
	case <-time.After(time.Second * 1):
		inputJSON := `[
		{"first_name": "Victor", "last_name": "Nguyen", "age": 100, "class_name":"golang"},
		{"first_name": "Anh", "last_name": "Dinh", "age":200, "class_name":"golang"}
		]`
		log.Println("output json string")
		w.Write([]byte(inputJSON))
	}
}
