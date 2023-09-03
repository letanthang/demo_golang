package main

import (
	"log"
	"os"
)

type closeFunc func() error

func main() {
	closeFile := logToFile()
	defer closeFile()
	
	log.Println("This is a test log entry")
}

func logToFile() closeFunc{
	f, err := os.OpenFile("testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	log.SetOutput(f)

	return func() error {
		return f.Close()
	}
}
