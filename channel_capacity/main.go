package main

import "fmt"

func main() {
	ch := make(chan int)

	// go func() {
	// ch <- 100
	// }()

	fmt.Println("hello world")

	fmt.Println(<-ch)
}
