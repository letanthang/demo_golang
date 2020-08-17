package main

import (
	"fmt"

	"github.com/letanthang/demo_golang/mongo_search/db"
	"github.com/letanthang/demo_golang/mongo_search/types"
)

func main() {
	students, err := db.SearchStudent(types.StudentSearchReq{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(students)
}
