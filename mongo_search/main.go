package main

import (
	"fmt"

	"app/mongo_search/db"
	"app/mongo_search/types"
)

func main() {
	students, err := db.SearchStudent(types.StudentSearchReq{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(students)
}
