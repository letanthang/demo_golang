# install protoc
linux: 
$ apt install -y protobuf-compiler

mac: 
$ brew install protobuf
# install golang plugin
$ go get google.golang.org/protobuf/cmd/protoc-gen-go \
         google.golang.org/grpc/cmd/protoc-gen-go-grpc
# update PATH
$ export PATH="$PATH:$(go env GOPATH)/bin"

# update *.proto - add "option go_package"
option go_package = "./hellogrpc";

# generate code from *.proto
protoc --go_out=. --go-grpc_out=. hellogrpc/hellogrpc.proto

# test service discovery
grpcurl -plaintext localhost:50051 list
grpcurl -plaintext -d '{"name":"thang1"}' localhost:50051 hellogrpc.Greeter/SayHellos