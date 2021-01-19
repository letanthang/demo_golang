package main

import "fmt"

func main() {
	c := make(chan string)

	go greet(c) // without this (sub goroutine), main goroutine block forever <=> deadlock
	c <- "World"

	fmt.Println("program end")
}

func greet(c chan string) {
	fmt.Printf("Hello %s\n", <-c)
}
