package main

import "context"

import "fmt"

type req int

const requestID req = 9999

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "requestName", "hello-request")

	ctx = context.WithValue(ctx, requestID, 123934939)

	//print context value

	name, ok := ctx.Value("requestName").(string)
	if !ok {
		fmt.Println("Cannot find context value")
	} else {
		fmt.Println(name)
	}

	reqID, ok := ctx.Value(requestID).(int)
	if !ok {
		fmt.Println("Cannot find context value")
	} else {
		fmt.Println(reqID)
	}

	reqID0, ok := ctx.Value(9999).(int)
	if !ok {
		fmt.Println("Cannot find context value")
	} else {
		fmt.Println(reqID0)
	}

}
