package main

import "fmt"

func main() {
	msg := "Lê Thắng"
	msgWithoutSpace := msg[0:2] + msg[3:]
	fmt.Println(msg)
	fmt.Println(msgWithoutSpace)

	runeSlice := []rune(msg)
	runeSliceWithoutSpace := append(runeSlice[0:2],runeSlice[3:]...)
	fmt.Println(string(runeSliceWithoutSpace))
}
