package main

import "fmt"

func Includes[T comparable](slice []T, elem T) bool {
	for _, v := range slice {
		if v == elem {
			return true
		}
	}
	return false
}

func Reverse[T any](slice []T) []T {
	result := make([]T, len(slice))
	for i, elem := range slice {
		result[len(slice)-1-i] = elem
	}
	return result
}

func main() {
	intSlice := []int{1, 4, 7, 9}
	fmt.Println(intSlice)
	fmt.Println("slice include", 9, Includes(intSlice, 9))
	fmt.Println("reverse is", Reverse(intSlice))

}
