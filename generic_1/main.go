package main

import "fmt"

func main() {
	intSlice := []int{1, 5, 10, 2}
	fmt.Println(Include(intSlice, 8))
	fmt.Println(Include(intSlice, 2))
	fmt.Println(Include([]string{"d", "a", "x"}, "x"))

	fmt.Println(Max(intSlice))

	myMap := map[string]any{"K": "Kotlin", "J": "Java", "M": 1000000}
	fmt.Println(Keys(myMap))
}

func Include[T comparable](slice []T, ele T) bool {
	for _, v := range slice {
		if ele == v {
			return true
		}
	}
	return false
}

type Number interface {
	int | float64
}

func Max[T Number](slice []T) T {
	result := slice[0]
	for _, v := range slice {
		if v > result {
			result = v
		}
	}
	return result
}

func Keys[K comparable, V any](m map[K]V) []K {
	l := len(m)
	result := make([]K, l)
	i := 0
	for k := range m {
		result[i] = k
		i++
	}
	return result
}