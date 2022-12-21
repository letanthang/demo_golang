package main

import "fmt"

func main() {
	a := []int{4, 2, 1, 9, 5}

	fmt.Println(a[:1])
	fmt.Println(a[2:])
	a = append(a[:1], a[2:]...)
	fmt.Printf("%v", a)
}
