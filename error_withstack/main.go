package main

import (
	"fmt"

	"emperror.dev/errors"
)

func main() {
	// Todo : not finish!
	err := errors.New("db error")
	err2 := errors.WithMessage(err, "get user 1")
	err3 := errors.WithStack(err)
	fmt.Println("err = ",err,", err2 = ",err2,", err3 = ", err3)

}
