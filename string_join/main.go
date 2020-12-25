package main

import (
	"fmt"
	"strings"
)

func main() {
	slice := []string{"hehe", "hoho", "hihi"}
	s := strings.Join(slice, "-")
	fmt.Println(s)
}
