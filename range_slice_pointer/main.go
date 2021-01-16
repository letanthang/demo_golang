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
	fmt.Println(slice[0], slice[1])
	for i, v := range slice {
		v.Class = "3B"
		slice[i].Age = 100
	}
	fmt.Println(slice[0], slice[1])
}
