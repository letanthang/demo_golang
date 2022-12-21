// underline slice is a struct contain a pointer reference to memory address
// copy array means copy pointer point to the same memory address
// so change slice - the copied slice change too
package main

import "fmt"

func main() {

	aSlice := []int{5, 9, 3}

	change(aSlice)
	fmt.Printf("%+v\n", aSlice)
	appendSlice(aSlice)
	fmt.Printf("%+v\n", aSlice)
}

func change(slice []int) {
	slice[0] = 1000
	fmt.Printf("%+v\n", slice)
}
func appendSlice(slice []int) {
	slice = append(slice, 100)
	fmt.Printf("%+v\n", slice)
}
