package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	result := make(map[string]int)
	a := strings.Fields(s)
	for _, v := range a {
		result[v] = result[v] + 1
	}
	return result
}

func main() {
	wc.Test(WordCount)
}
