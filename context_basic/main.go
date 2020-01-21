package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	// timeout after 2s
	// ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	// cancel after 1s
	// time.AfterFunc(time.Second*1, func() {
	// 	cancel()
	// })
	sleepAndTalk(ctx, "hello")
}

func sleepAndTalk(ctx context.Context, msg string) {
	done := make(chan struct{})
	go func() {
		time.Sleep(time.Second * 3)
		fmt.Println(msg)
		done <- struct{}{}
	}()

	select {
	case <-done:
		return
	case <-ctx.Done():
		fmt.Println("error:", ctx.Err().Error())
	}
}
