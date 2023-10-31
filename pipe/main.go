package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	pr, pw := io.Pipe()
	

	go func() {
		fmt.Fprint(pw, "Hello world")
		pw.Close()
	}()

	_, err := io.Copy(os.Stdout, pr)
	if err != nil {
		log.Fatal(err)
	}
}
