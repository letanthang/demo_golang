package main

import (
	"fmt"
	"github.com/letanthang/demo_golang/type_system/student"
	"reflect"
)

func main() {

	aStruct := struct {
		Name string
		Age  int
	}{Name: "Thang", Age: 35}

	aType := reflect.TypeOf(aStruct)

	fmt.Printf("aStruct has type %s and kind %s\n", aType.Name(), aType.Kind())

	bStruct := student.Student{FirstName: "Phat", LastName: "Nguyen"}
	bType := reflect.TypeOf(bStruct)
	fmt.Printf("aStruct has type %s and kind %s\n", bType.Name(), bType.Kind())

	fmt.Println(bType.Field(1))
}
