package main

import "fmt"

type Status string

const (
	New        Status = "New"
	WaitToPick        = "WTP"
	Picking           = "Picking"
	Picked            = "Picked"
	Storing           = "Storing"
	Delivering        = "Delivering"
	Delivered         = "Delivered"
)

func main() {
	var orderStatus Status
	orderStatus = Picking
	fmt.Println(orderStatus)
	if orderStatus == Picked {
		fmt.Println("What??? Picked really?")
	}

	if orderStatus == "Picking" {
		fmt.Println("yeah Picking")
	}
}
