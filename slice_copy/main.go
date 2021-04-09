// underline slice is a struct contain a pointer reference to memory address
// copy array means copy pointer point to the same memory address
// so change slice - the copied slice change too
package main

import "fmt"

func main() {

	aSlice := []int{5, 9, 3}

	change(aSlice)
	fmt.Printf("%+v", aSlice)
}

func change(slice []int) {
	slice[0] = 1000
	fmt.Printf("%+v", slice)
}
