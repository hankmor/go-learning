// 构建约束

//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"wire_demo"
)

// 注入器
func InitializeEvent() wire_demo.Event {
	//wire.Build(wire_demo.NewEvent, wire_demo.NewGreeter) //  no provider found for wire_demo.Message
	wire.Build(wire_demo.NewEvent, wire_demo.NewGreeter, wire_demo.NewMessage)
	return wire_demo.Event{}
}
