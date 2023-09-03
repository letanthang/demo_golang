package main

import "fmt"

var WhatIsThe = AnswerToLife()

func AnswerToLife() int { // 1
	return 42
}

func init() { // 2
	WhatIsThe = 0
}

func main() { // 3
	if WhatIsThe == 0 {
		fmt.Println("It's all a lie.")
	}
}
