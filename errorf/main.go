package main

import "fmt"

func main() {
	const name, id = "bueller", 17
	err := fmt.Errorf("users %q (id %d) not found", name, id)
	fmt.Println(err.Error())
	fmt.Printf("%+v", err)
}
