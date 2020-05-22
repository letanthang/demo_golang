//use protoc to gen go files
//go install google.golang.org/protobuf/cmd/protoc-gen-go
protoc --go_out=. echo.proto