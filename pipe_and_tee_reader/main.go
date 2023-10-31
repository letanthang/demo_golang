package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {
	resp, err := http.Get("https://api.ipify.org?format=json")
	if err != nil {
		log.Fatal(err)
	}

	pr, pw := io.Pipe()

	wg.Add(1)
	go func() {
		fmt.Println("Writing in pipe writer in goroutine")
		bs, err := io.ReadAll(io.TeeReader(resp.Body, pw))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(bs))
		defer func() {
			resp.Body.Close()
			pw.Close()
			wg.Done()
		}()

	}()

	scanner := bufio.NewScanner(pr)
	for scanner.Scan() {
		fmt.Println(scanner.Text())

	}
	defer pr.Close()

	wg.Wait()
}
