package main

import "fmt"

// Student :
type Student struct {
	Name  string
	Age   int
	Class string
}

func main() {
	slice := []Student{{"Thang", 36, "2H"}, {"Tram", 29, "2A"}}
	pointerSlice := make([]*Student, 2)
	pointerSlice1 := make([]*Student, 2)
	// v is just a copy!!!
	for i, v := range slice {
		// &v do not work
		pointerSlice[i] = &v
		pointerSlice1[i] = &slice[i]
	}
	fmt.Println(pointerSlice[0], pointerSlice[1])
	fmt.Println(pointerSlice1[0], pointerSlice1[1])
}
