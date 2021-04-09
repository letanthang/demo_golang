package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

var slice1 = []Student{{"Hang", 10}, {"Trung", 20}}

func main() {

	fmt.Printf("%+v\n", slice1)

	wrong()
	right1()
	right2()
}
func wrong() {
	result := []*Student{}

	for _, v := range slice1 {
		result = append(result, &v)
	}
	fmt.Printf("%+v, [%v, %v]\n", result, *result[0], *result[1]) // same address-> wrong

}

func right1() {
	result := []*Student{}

	for i := range slice1 {
		result = append(result, &slice1[i])
	}
	fmt.Printf("%+v, [%v, %v]\n", result, *result[0], *result[1])

}

func right2() {
	result := []*Student{}

	for _, v := range slice1 {
		v := v
		result = append(result, &v)
	}
	fmt.Printf("%+v, [%v, %v]\n", result, *result[0], *result[1])

}
