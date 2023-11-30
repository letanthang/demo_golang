package main

import "fmt"

func main() {
	value := "cat;dog"

	fmt.Println(len(value))
	fmt.Println(value[:3])
	// Take substring from index 4 to length of string.
	fmt.Println(value[4:])

	//print value without ";"
	fmt.Println(value[:3] + value[4:])
	fmt.Println(value[1:2], value[3:4])
	fmt.Println(value[0], value[0:1])
}
