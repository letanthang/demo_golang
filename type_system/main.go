package main

import (
	"fmt"
	"github.com/letanthang/demo_golang/type_system/student"
	"strconv"
)

func main() {
	fmt.Println("Hello nc student!")

	// 	Int

	aInt := 1
	fmt.Println(aInt)

	var bInt int
	bInt = 2
	fmt.Printf("bInt = %d\n", bInt)
	// String

	aString := "Hello abc"
	fmt.Printf("aString = %s\n", aString)

	// Array/Slice
	aSlice := []string{"a", "b", "c"}
	fmt.Println(aSlice)

	bSlice := []int{1, 4, 6, 8, 8}
	fmt.Println(bSlice)
	fmt.Printf("%v\n", bSlice)

	for _, v := range bSlice {
		if v > 5 {
			fmt.Printf("%d ", v)
		}
	}

	aArray := [2]int{1, 2}
	fmt.Println(aArray)

	// Map
	aMap := map[string]int{"age": 1000, "level": 5}
	fmt.Println(aMap)

	// Struct

	// vic := model.Student{
	// 	FirstName: "Victor",
	// 	LastName:  "Nguyen",
	// 	Age:       100,
	// 	Email:     "victornguyen@gmail.com",
	// }

	vic := student.Student{"Victor", "Nguyen", 100, "hihihi"}

	fmt.Println(vic)
	fmt.Printf("%+v", vic)
	// Interface
	var i interface{}
	i = vic
	fmt.Println(i)

	// Channel

	ch := make(chan int, 2)
	ch <- 10 // write to channel
	ch <- 12 // write to channel

	fmt.Println(<-ch) //read from channel
	fmt.Println(<-ch) //read from channel

	// Function

	// function receiver

	// pointer

	fmt.Println(vic.GetFullName())
	vic.SetEmail("hehehe")
	SetEmail(&vic, "hahaha")

	fmt.Println(vic.Email)

	// int <-> string

	cInt := 100

	cString := fmt.Sprintf("%d hello", cInt)

	cString = strconv.Itoa(cInt)
	fmt.Println(cString)
	// "100"

	// parse inputJson to struct

	// inputJson := `[
	// 	{"first_name": "Victor", "last_name": "Nguyen", "age": 100, "class_name":"golang"},
	// 	{"first_name": "Anh", "last_name": "Dinh", "age":200, "class_name":"golang"}
	// 	]`

	// fmt.Println(inputJson)
}

func SetEmail(s *student.Student, email string) {
	s.Email = email
}
