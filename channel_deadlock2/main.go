package main

import "fmt"

func main() {
	ch := make(chan int) //unbuffer channel

	// write to unbuffer channel cause current routine sleep
	ch <- 1

	fmt.Println("program end")
}
