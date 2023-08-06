package wire_demo

import "fmt"

type Message string

type Greeter struct {
	Msg Message
}

type Event struct {
	Greeter Greeter
}

func NewMessage(name string) Message {
	return Message(fmt.Sprintf("hello, %s!", name))
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
