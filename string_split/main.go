package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "hello-how-are-you-?"
	a := strings.Split(s, "-")
	fmt.Printf("v :%+v len:%d cap:%d", a, len(a), cap(a))
}
