package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func Sqrt(x float64) (float64, error) {
	var err error
	var ret float64
	if x < 0 {
		err = ErrNegativeSqrt(x)
	} else {
		ret = math.Sqrt(x)
	}
	return ret, err
}

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot sqrt negative number: %f", float64(e))
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
