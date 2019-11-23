package main

import (
	"fmt"
)

func main() {

	//map

	aMap := map[string]int{"a": 1, "b": 3, "d": 5}

	for i, v := range aMap {
		fmt.Printf("index at %s has value %d \n", i, v)
	}

	//map[string]interface{}

	// aStudent := student.Student{FirstName: "Phuoc", LastName: "Nguyen", Age: 100}

	// var m map[string]interface{}

	// m["FirstName"] = "Phuoc"
	// m["Age"] = 100

	// channel

}
