package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("http://localhost:9090")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	io.Copy(os.Stdout, res.Body)

}
