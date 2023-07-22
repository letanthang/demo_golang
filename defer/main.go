package main

import (
	"errors"
	"fmt"
)

func main() {
	Foo()
	Foo2()
}

func Foo() (err error) {
	//!!!!! we cannot defer with error checking like this!!!!!
	defer func(err error) {
		if err != nil {
			fmt.Println("world")
		}
	}(err)

	fmt.Println("hello")

	if true {
		err = errors.New("Bar error")
		return
	}

	return nil
}

func Foo2() (err error) {
	// check err the right way
	defer func() {
		if err != nil {
			fmt.Println("world")
		}
	}()

	fmt.Println("hello")

	if true {
		err = errors.New("bar error")
		return
	}

	return nil
}
