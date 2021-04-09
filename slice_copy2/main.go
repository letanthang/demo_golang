package main

import "fmt"

func main() {

	slice1 := []int{5, 9, 3}

	slice2 := slice1

	slice2[0] = 1000

	fmt.Printf("slice1=%+v, slice2=%+v", slice1, slice2)
}
