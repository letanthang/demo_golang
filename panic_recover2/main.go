package main

import "fmt"

func recoveryFunction() {
	// if recoveryMessage := recover(); recoveryMessage != nil {
	// 	fmt.Println(recoveryMessage)
	// }
	fmt.Println("This is recovery function...")
}

func executePanic() {
	defer recoveryFunction()
	panic("This is Panic Situation")
	fmt.Println("The function executes Completely")
}

func main() {
	executePanic()
	fmt.Println("Main block is executed completely...")
}
