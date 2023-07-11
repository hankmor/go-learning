package wire_demo

import "fmt"

type Message string

type Greeter struct {
	Msg Message
}

type Event struct {
	Greeter Greeter
}

func NewMessage() Message {
	return Message("hello, hank!")
}

func NewGreeter(msg Message) Greeter {
	return Greeter{Msg: msg}
}

func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
}

func (g Greeter) Greet() Message {
	return g.Msg
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}
