package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

func main() {
	j := 0
	defer func(i int) {
		fmt.Println("exiting at:", time.Now(), j)
	}(j)

	fmt.Println("cpu num:", runtime.NumCPU())
	forever := make(chan os.Signal, 1)
	signal.Notify(forever, syscall.SIGINT, syscall.SIGTERM)
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	go func() {
		for {
			j++
			fmt.Println("worldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworld")
			time.Sleep(time.Millisecond * 300)
		}
	}()

	// this goroutine will keep the program run forever
	go printHello(ctx)

	fmt.Println("awaiting signal")
	<-forever
	fmt.Println("j=", j)
}

func printHello(ctx context.Context) {
	i := 1

	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("hello", i)
			i++
			// time.Sleep(time.Microsecond * 200)
		}
	}

}
