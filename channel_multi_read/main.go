package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	c1 := make(chan int, 1)
	c2 := make(chan bool, 1)

	go func() {
		ms := 100 + rand.Intn(300-100)
		fmt.Println("sleep in", ms, "ms")
		time.Sleep(time.Duration(ms) * time.Millisecond)
		// c1 <- 100
		close(c1)
	}()

	go func() {
		ms := 100 + rand.Intn(300-100)
		fmt.Println("sleep in", ms, "ms")
		time.Sleep(time.Duration(ms) * time.Millisecond)
		close(c2)
	}()
	go func() {
		select {
		case <-c1:
			fmt.Println("c1: receive from routine")
		case flag := <-c2:
			fmt.Println("c2: receive from routine", flag)
		}
	}()

	select {
	case <-c1:
		fmt.Println("c1: receive")
	case flag := <-c2:
		fmt.Println("c2: receive", flag)
	}

	fmt.Println("Program end")
}
