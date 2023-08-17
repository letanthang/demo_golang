package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Println(runtime.GOMAXPROCS(0), "threads")
	slice1 := make([]int, 2, 4)
	slice1[0] = 5
	slice1[1] = 10

	slice2 := slice1
	slice2 = append(slice2, 15)

	fmt.Println(&slice1[0], &slice2[0])
	slice2 = append(slice2, 20, 25)

	fmt.Println(len(slice1), cap(slice1))
	fmt.Println(slice1, slice2)

	fmt.Println(&slice1[0], &slice2[0])


}
