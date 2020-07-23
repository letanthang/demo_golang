
main.go  
...
func main() {  
    e := InitializeEvent()  

    e.Start()  
}  
wire.go  
//+build wireinject   
func InitializeEvent() Event {  
    wire.Build(NewEvent, NewGreeter, NewMessage)  
    return Event{}  
}  

install the tool with:  
$ go get github.com/google/wire/cmd/wire  

run wire to genarate wire_gen.go  
$ wire  


