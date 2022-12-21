package main

import "fmt"

func main() {
	str := "Ahello world"
	fmt.Println(str[0]) // => ascii 65
	fmt.Println(str[0:1]) // => A  | but not work for utf8
	fmt.Println([]rune(str)[0])           // => ascii 65
	fmt.Println(string([]rune(str)[0]))   // => A
	fmt.Println(string([]rune(str)[0:5])) // => Ahell

	//ascii code to string
	aInt := 65
	r := rune(aInt)
	aString := string(r)
	fmt.Println(aString)

	// thang practise 12/2022
	aStr := "Á I Love you"
	fmt.Println(string([]rune(aStr)[0]))

	t := aStr[0]
	fmt.Println(string(t))
}
