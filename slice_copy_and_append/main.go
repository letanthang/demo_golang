package main

import "fmt"

func main() {

	aSlice := []int{5, 9, 3}
	fmt.Printf("%+v\n", aSlice)
	change(aSlice)
	fmt.Printf("%+v\n", aSlice)
}

func change(slice []int) {
	slice = append(slice, 10) // this line break the pointer so input slice won't change!!!!
	// this behavior is supper normal to a struct pointer 
	slice[0] = 1000
}
