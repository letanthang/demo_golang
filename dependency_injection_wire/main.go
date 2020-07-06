package main

import "fmt"

type Message string

type Greeter struct {
	Message Message // <- adding a Message field
}

func (g Greeter) Greet() Message {
	return g.Message
}

type Event struct {
	Greeter Greeter
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

func NewMessage() Message {
	return Message("Hi there!")
}
func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}

func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
}

func main() {
	e := InitializeEvent()

	e.Start()
}
