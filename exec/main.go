package main

import (
	"bufio"
	"bytes"
	"context"
	"log"
	"os"
	"os/exec"
	"time"
)

// LogWriter Custom Log output type that includes a context for logging
type LogWriter struct {
	ctx context.Context
}

func (c LogWriter) Write(p []byte) (int, error) {

	scanner := bufio.NewScanner(bytes.NewReader(p))

	// log.Infof(c.ctx, string(p))
	log.Println("thanglog", c.ctx, string(p))

	if err := scanner.Err(); err != nil {
		log.Printf("Error scanning input: %v", err)
	}
	return len(p), nil
}

func main() {

	ctx, cancel := context.WithTimeout(context.TODO(), 6*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "yarn", "install")
	cmd.Dir = "project"
	cmd.Stdout = LogWriter{ctx: ctx}
	cmd.Stderr = LogWriter{ctx: ctx}
	if err := cmd.Run(); err != nil {
		// yarn install command should be interrupted after 60 sec
		log.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}
