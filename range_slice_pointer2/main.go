package main

import "fmt"

// Student :
type Student struct {
	Name  string
	Age   int
	Class Class
}

// Class :
type Class struct {
	ID   int
	Name string
}

func main() {
	slice := []Student{{"Thang", 36, Class{1, "2A"}}, {"Tram", 29, Class{1, "2B"}}}

	// v is just a copy!!!
	for i, v := range slice {
		// v.Name or v.Class.Name do not work
		v.Name = "haha"
		v.Class.Name = "hehe"
		slice[i].Class.ID = 100
	}

	fmt.Println(slice)
}
