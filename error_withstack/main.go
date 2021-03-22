package main

import (
	"fmt"

	"emperror.dev/errors"
)

func main() {
	// Todo : not finish!
	err := errors.New("db error")
	err2 := errors.WithStack(err)
	fmt.Println(err2)

}
