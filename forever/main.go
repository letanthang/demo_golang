package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	defer fmt.Println("exiting")

	fmt.Println("cpu num:", runtime.NumCPU())
	forever := make(chan int, 1)

	// this goroutine will keep the program run forever
	go func() {
		i := 1
		for {
			fmt.Println("hello", i)
			i++
			time.Sleep(time.Second)
		}
	}()
	fmt.Println("awaiting signal")
	<-forever
}
