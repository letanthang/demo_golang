package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

type Student struct {
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Age       int    `json:"age"`
	Classname string `json:"class_name"`
}

func main() {
	fmt.Println("======= BAI 1 =======")
	cInt := 100
	c := strconv.Itoa(cInt)
	fmt.Printf("%s %s\n", c, reflect.TypeOf(c))
	d, _ := strconv.Atoi(c)
	fmt.Printf("%d %s\n", d, reflect.TypeOf(d))

	fmt.Println("======= BAI 2 =======")

	student := []Student{}
	inputJson := `[
		{"first_name": "Victor", "last_name": "Nguyen", "age": 100, "class_name":"golang"},
		{"first_name": "Anh", "last_name": "Dinh", "age":200, "class_name":"golang"}
		]`

	err := json.Unmarshal([]byte(inputJson), &student)

	if err != nil {
		fmt.Println(err)
		panic("wrong json format")
	}

	fmt.Printf("%+v\n", student)

	// bai toan nguoc : struct -> inputJson string

	bs, err := json.Marshal(student)
	if err != nil {
		fmt.Println(err)
		return
	}

	jsonString := string(bs)

	fmt.Println(jsonString)

}
