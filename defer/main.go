package main

import (
	"errors"
	"fmt"
)

func main() {
	Foo()
}

func Foo() (err error) {
	//!!!!! we cannot defer with error checking like this!!!!!
	defer func(err error) {
		if err == nil {
			fmt.Println("world")
		}
	}(err)

	if true {
		err = errors.New("Bar error")
		return
	}

	fmt.Println("hello")
	return nil
}
