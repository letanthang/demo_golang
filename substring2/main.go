package main

import "fmt"

func main() {
	value := "cat;dog"

	s1 := value[:3]
	fmt.Println(s1)
	// Take substring from index 4 to length of string.
	substring := value[4:len(value)]
	fmt.Println(substring)

	//print value without ";"
	fmt.Println(value[:3] + value[4:])
}
