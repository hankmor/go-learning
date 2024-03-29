// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"wire_demo"
)

// Injectors from wire.go:

// 注入器
func InitializeEvent(name string) wire_demo.Event {
	message := wire_demo.NewMessage(name)
	greeter := wire_demo.NewGreeter(message)
	event := wire_demo.NewEvent(greeter)
	return event
}
