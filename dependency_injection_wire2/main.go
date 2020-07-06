package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

type Message string

type Greeter struct {
	Grumpy  bool
	Message Message // <- adding a Message field
}

func (g Greeter) Greet() Message {
	if g.Grumpy {
		return "Go away!"
	}
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
	var grumpy bool
	if time.Now().Unix()%2 == 0 {
		grumpy = true
	}
	return Greeter{Message: m, Grumpy: grumpy}
}

func NewEvent(g Greeter) (Event, error) {
	return Event{Greeter: g}, errors.New("could not create event: event greeter is grumpy")
}

func main() {
	e, err := InitializeEvent()
	if err != nil {
		fmt.Printf("failed to create event: %s\n", err)
		os.Exit(2)
	}

	e.Start()
}
