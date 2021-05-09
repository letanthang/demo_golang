//use protoc to gen go files
//go install google.golang.org/protobuf/cmd/protoc-gen-go
protoc --go_out=. echo.proto

//or
protobuf_sample> protoc --go_out=. proto/echo/echo.proto