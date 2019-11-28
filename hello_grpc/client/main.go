package main

import (
	"context"
	"log"

	pb "github.com/letanthang/demo_golang/hello_grpc/hellogrpc"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "Thang"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)
	resp, err := c.SayHello(context.TODO(), &pb.HelloRequest{Name: defaultName})
	if err != nil {
		log.Fatalf("cound not send greeting: %v", err)
	}
	log.Printf("Response greeting: %s", resp.Message)
}

// func main() {
// 	conn, err := grpc.Dial(address, grpc.WithInsecure())
// 	if err != nil {
// 		log.Fatalf("did not connect: %v", err)
// 	}

// 	defer conn.Close()
// 	c := pb.NewGreeterClient(conn)

// 	name := defaultName

// 	if len(os.Args) > 1 {
// 		name = os.Args[1]
// 	}

// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
// 	defer cancel()

// 	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})

// 	if err != nil {
// 		log.Fatalf("cound not send greet: %v", err)
// 	}

// 	log.Printf("Response greeting: %s", r.Message)

// }
