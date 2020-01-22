package main

import "net/http"

import "log"

import "fmt"

import "time"

func main() {
	http.HandleFunc("/", handleFunc)
	time.AfterFunc(time.Millisecond*100, func() {
		fmt.Println("Server listen at 9090")
	})

	log.Fatalf("error: %v", http.ListenAndServe(":9090", nil))
}

func handleFunc(w http.ResponseWriter, r *http.Request) {
	select {
	case <-time.After(time.Second * 3):
		log.Println("output hello")
		w.Write([]byte("hello"))
	}
}
