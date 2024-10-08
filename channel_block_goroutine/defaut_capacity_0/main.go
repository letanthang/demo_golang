package main

import "fmt"

func main() {
	c := make(chan string)

	// without this sub goroutine, main goroutine block forever
	// because the size exceed capacity will put go routine to sleep

	go greet(c) 
	c <- "World"

	fmt.Println("program end")
}

func greet(c chan string) {
	fmt.Printf("Hello %s\n", <-c)
}
