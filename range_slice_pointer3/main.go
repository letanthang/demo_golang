package main

import "fmt"

// Student :
type Student struct {
	Name  string
	Age   int
	Class string
}

func main() {
	slice := []*Student{{"Thang", 36, "2H"}, {"Tram", 29, "2A"}}
	slice2 := make([]*Student, 2)
	fmt.Println(slice[0], slice[1])

	// v is just a copy!!!
	// but copy of a pointer will work in case of reference
	for i, v := range slice {
		slice2[i] = v
		v.Class = "3B"
		slice[i].Age = 100
	}
	fmt.Println(slice[0], slice[1])

	fmt.Println(slice2[0], slice2[1])
}
