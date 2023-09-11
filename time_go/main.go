package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Student struct {
	FirstName   string
	Birthday    time.Time
	GraduateDay time.Time
}

func main() {
	a := Student{
		FirstName:   "Thang",
		Birthday:    time.Date(1984, 3, 20, 0, 0, 0, 0, time.Local),
		GraduateDay: time.Date(2024, 3, 20, 0, 0, 0, 0, time.Local),
	}

	bs, err := json.Marshal(a)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(bs))

}
