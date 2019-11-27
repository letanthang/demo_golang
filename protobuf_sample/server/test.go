package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
	"github.com/letanthang/demo_golang/protobuf_sample/proto/echo"
)

func test() {
	req := &echo.EchoRequest{Name: "Hello"}
	data, err := proto.Marshal(req)

	if err != nil {
		log.Fatalf("Error while marshalling the object: %v", err)
	}

	res := &echo.EchoRequest{}
	err = proto.Unmarshal(data, res)
	if err != nil {
		log.Fatalf("error while unmarshalling the object: %v", err)
	}
	fmt.Printf("value from the unmarshalled data is %v", res.GetName())

}
