package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name   string  `json:"name"`
	Age    int     `json:"age,string"`
	Height float32 `json:"height"`
}

func main() {
	str := `{"name": "Thang", "age": "100", "height": 1.6}`
	var aPerson Person
	err := json.Unmarshal([]byte(str), &aPerson)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("struct: %+v", aPerson)
}
