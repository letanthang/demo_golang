package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//1. timeout & slice advance
	forever := make(chan int)
	defer fmt.Println("Program exit")
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*4)
	defer cancel()

	go func(ctx context.Context) {
		i := 0
		for {
			select {
			case <-ctx.Done():
				fmt.Println("goroutine finish")
				close(forever)
				return
			default:
				fmt.Println("hello", i)
				i++
				time.Sleep(time.Second)

			}
		}

	}(ctx)

	<-forever

}
