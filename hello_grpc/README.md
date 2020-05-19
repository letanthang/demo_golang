grpcurl -plaintext localhost:50051 list
grpcurl -plaintext -d '{"name":"thang1"}' localhost:50051 hellogrpc.Greeter/SayHello