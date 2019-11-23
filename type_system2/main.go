package main

import (
	"encoding/json"
	"fmt"
	"github.com/letanthang/demo_golang/type_system/student"
	"reflect"
	"strconv"
)

func main() {

	//map

	aMap := map[string]int{"a": 1, "b": 3, "d": 5}

	for i, v := range aMap {
		fmt.Printf("index at %s has value %d \n", i, v)
	}

	//map[string]interface{}

	aStudent := student.Student{FirstName: "Phuoc", LastName: "Nguyen", Age: 100}

	bs, _ := json.Marshal(aStudent)

	var m map[string]interface{}

	json.Unmarshal(bs, &m)

	fmt.Printf("%+v", m)

	// channel
	ch := make(chan int)

	go func() {
		ch <- 1
		ch <- 2
		ch <- 8
		ch <- 4
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}

	//array

	// var location [2]float32

	// slice
	// var ids []int

	// Type casting (string <-> byte slice, string - int, int - float, struct - struct, struct <-> map[string]interface{}

	// string - byte slice
	aString := "hello world"

	bs = []byte(aString)

	bString := string(bs)

	fmt.Println(bString)

	// int <-> string
	aInt := 100

	aString = strconv.Itoa(aInt)

	fmt.Println(aString + "mot tram")

	bInt, err := strconv.Atoi(aString)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bInt + 100)

	// int <-> float

	aInt = 300

	var aFloat float32
	aFloat = float32(300)
	type1 := reflect.TypeOf(aFloat)
	fmt.Println(aFloat, "type =", type1.Name(), "kind=", type1.Kind())

	cInt := int(aFloat)
	type1 = reflect.TypeOf(cInt)
	fmt.Println(cInt, "type =", type1.Name(), "kind=", type1.Kind())

	// struct <-> struct

	type Boy struct {
		FirstName string
		LastName  string
		Age       int
		Email     string
	}

	boy := Boy{"Phi", "Nguyen", 300, "phi@gmail.com"}

	type1 = reflect.TypeOf(boy)
	fmt.Println(boy, "type =", type1.Name(), "kind=", type1.Kind())

	bStudent := student.Student(boy)

	fmt.Println(bStudent)

	// type conversion vs type assertion
	//  Type assertion interface{} -> string, int, struct

	var i interface{}

	i = 100

	aInt = i.(int)

	fmt.Println(aInt)

	fmt.Printf("%+v \n", m)

	if dString, ok := m["first_name"].(string); ok {
		fmt.Println(dString)
	} else {
		fmt.Println("cannot assert, wrong type!")
	}

}
