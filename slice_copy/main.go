package main

import "fmt"

func main() {

	aSlice := []int{5, 9, 3}

	change(aSlice)
	fmt.Printf("%+v", aSlice)
}

func change(slice []int) {
	slice = append(slice, 6)
	slice[0] = 10
	fmt.Printf("%+v", slice)
}
