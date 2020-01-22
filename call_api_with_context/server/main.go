package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", decorateRequestID(handleFunc))
	time.AfterFunc(time.Millisecond*100, func() {
		fmt.Println("Server listen at 9090")
	})

	log.Fatalf("error: %v", http.ListenAndServe(":9090", nil))
}

func decorateRequestID(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		requestID := 1000000 + rand.Intn(1000000)
		ctx = context.WithValue(ctx, "requestID", requestID)
		req := r.WithContext(ctx)
		h(w, req)
	}
}

func handleFunc(w http.ResponseWriter, r *http.Request) {
	reqCtx := r.Context()
	requestID, ok := reqCtx.Value("requestID").(int)
	if !ok {
		log.Fatal("Cannot reead request id from ctx")
	}
	log.Println("/ request id:", requestID)
	select {
	case <-reqCtx.Done():
		log.Println(reqCtx.Err().Error())
	case <-time.After(time.Second * 3):
		log.Println("output hello")
		w.Write([]byte("hello"))
	}
}
