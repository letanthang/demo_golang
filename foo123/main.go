package main

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

type Student struct {
	FirstName string
	LastName  string
	Age       int
	Grade     float32
	ClassName string
}

func main() {
	slice := []int{0,1,2}

	pas 

}

func getfullname(firstname, lastname string) string {
	return "Hello! I am " + firstname + " " + lastname

}

// cách 2
func getfullname2(firstname, lastname string) string {
	msg := fmt.Sprintf("Hello world! I am %s %s %d", firstname, lastname, 2023)
	return msg
}
func getDays(day []string) string {
	return day[0] + "," + day[1] + "," + day[2]
}

func passByValue2(a *[]int) {
	*a = append(*a, 100)
}
