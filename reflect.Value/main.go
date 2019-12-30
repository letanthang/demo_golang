package main

import (
	"fmt"
	"reflect"
)

type Anon reflect.Value

func (a Anon) IsA(typeToAssert reflect.Kind) bool {
	return typeToAssert == reflect.Value(a).Kind()
}

func toAnon(obj interface{}) Anon {
	return Anon(reflect.ValueOf(obj))
}

func main() {
	f := 10.4

	anon := toAnon(f)
	if anon.IsA(reflect.String) {
		fmt.Println("Anon is string")
	} else if anon.IsA(reflect.Int) {
		fmt.Println("Anon is int")
	} else if anon.IsA(reflect.Float32) {
		fmt.Println("Anon is float32")
	} else if anon.IsA(reflect.Float64) {
		fmt.Println("Anon is float64")
	} else {
		fmt.Println("failed")
	}

}
