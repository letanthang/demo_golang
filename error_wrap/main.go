package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func main() {
	err := errors.New("db error")
	err2 := errors.Wrap(err, "get user")
	fmt.Println(err2)
}
