package main

import "fmt"

func main() {
	str := "Ahello world"
	fmt.Println(str[0])                   // => ascii 65
	fmt.Println([]rune(str)[0])           // => ascii 65
	fmt.Println(string([]rune(str)[0]))   // => A
	fmt.Println(string([]rune(str)[0:5])) // => Ahell

	//ascii code to string
	aInt := 65
	r := rune(aInt)
	aString := string(r)
	fmt.Println(aString)
}
