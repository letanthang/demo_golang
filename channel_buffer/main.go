package main

import "fmt"

func main() {
	ch := make(chan int, 2) // buffer channel with capacity 2

	// len <= capacity -> no sleep
	ch <- 1
	ch <- 5

	// len > capacity -> routine sleep
	ch <- 6

	fmt.Println("Read channel")
	fmt.Println(<-ch)
	fmt.Println("Program end")
}
