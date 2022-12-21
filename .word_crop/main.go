package main

import (
	"fmt"
	"strings"
)

func main() {
	result := Solution("super doc", 4)
	fmt.Println(result)
}

func Solution(message string, K int) string {
	// Implement your solution here
	if len(message) <= K {
		return message
	}

	str := message[0 : K+1]
	lastChar := str[K : K+1]
	if lastChar != " " || len(str)+3 > K {
		str = CutTheLastWord(str, K) + "..."
	}

	return str
}

func CutTheLastWord(str string, K int) string {
	if str == "" {
		return ""
	}
	str = strings.TrimRight(str, " ")
	for str[len(str)-1:] != " " && len(str) > 0 {
		if len(str) == 1 {
			str = ""
			break
		} else {
			str = str[0 : len(str)-1]
		}
	}

	if len(str)+3 <= K {
		return str
	} else {
		return CutTheLastWord(str, K)
	}
}
