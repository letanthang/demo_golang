package main

import (
	"fmt"
	"sort"
)

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func main() {
	intSlice := []int{5, 2, 3, 2}
	Solution(intSlice)
}

func Solution(A []int) int {

	// lấy đầu ko lấy đuôi
	fmt.Println(A[0:1], A[1:2])

	sort.Sort(sort.Reverse(sort.IntSlice(A)))
	fmt.Println(A)

	min := A[len(A)-1]
	max := A[0]
	for i := min; i < max; i++ {
		current := A[len(A)-1]
		if i == current {
			break
		}
	}
	return 100
}
