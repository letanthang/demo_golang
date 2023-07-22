package main

import (
	"fmt"
)

func readChannel(ch chan int) {
	a := <-ch
	b := <-ch
	fmt.Println("Read from channel: ", a, b)
}

func declare (ch chan int){
	ch <- 14
	ch <- 20
}

func main() {
	defer func() {
		fmt.Println("Program end")
	}()
	ch := make(chan int)
	go declare(ch)
	readChannel(ch)
}
