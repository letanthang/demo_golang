package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Le Tan Thang"
	a := strings.Fields(s)
	fmt.Printf("value:%+v, len:%d, cap:%d\n", a, len(a), cap(a))
	fmt.Println(strings.Join(a, "-"))
}
