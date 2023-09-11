package main

import (
	"fmt"
)

type Student struct {
	FirstName string
	LastName  string
	Age       int
	Grade     float32
	ClassName string
}

func main() {

	ten := "Sang"
	ho := "Nguyen"
	msg := getfullname(ten, ho)
	fmt.Println(msg)
	msg1 := getfullname2(ten, ho)
	fmt.Println(msg1)

	//khai báo slice
	days := []string{"Monday", "Tuesday", "Wednesday"}
	msg2 := getDays(days)
	fmt.Println(msg2)
	manh := Student{
		FirstName: "Manh",
		LastName:  "Nguyen",
		Age:       26,
		Grade:     9,
		ClassName: "HHM",
	}
	fmt.Println(manh)
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
