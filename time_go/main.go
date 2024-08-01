package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

const timeLayout = "20060102T15:04:05MST"

var (
	VST = time.FixedZone("Asia/Ho_Chi_Minh", 7*60*60)
	JST = time.FixedZone("Asia/Tokyo", 9*60*60)
)

type Student struct {
	FirstName   string
	Birthday    time.Time
	GraduateDay time.Time
}

func main() {

	// createTime()
	parseTimeWithTimezone()

}

func createTime() {
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

func parseTimeWithTimezone() {
	stringTime := "20240801T12:20:05JST"
	t, err := time.ParseInLocation(timeLayout, stringTime, JST)
	if err != nil {
		fmt.Println(err)
		return
	}

	// utc:= t.UTC()
	// loc, err := time.LoadLocation("Asia/Ho_Chi_Minh")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	fmt.Println(t.UTC())
	fmt.Println(time.Now().UTC())
}
