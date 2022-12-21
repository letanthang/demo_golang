package main

import "fmt"

func main() {
	str := "ervervige"

	fmt.Println(Solution(str))

}

func Solution(S string) int {
	str := S
	m := map[byte]int{}
	for i := 0; i < len(str); i++ {
		if v, ok := m[str[i]]; !ok {
			m[str[i]] = 1
		} else {
			m[str[i]] = v + 1
		}
	}
	fmt.Println(m)

	//find odd element

	oddCount := 0
	for _, v := range m {
		if v%2 != 0 {
			oddCount++
		}
	}

	if oddCount <= 1 {
		return 0
	} else {
		return oddCount - 1
	}

	return oddCount

}
