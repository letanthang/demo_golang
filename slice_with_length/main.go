package main

import "fmt"

func main() {
	slice := make([]int, 4)
	fmt.Println(len(slice))
	slice = append(slice, 10)
	fmt.Println(len(slice))

	// if you want to set length of slice, u must assign to each slice[i]
}
