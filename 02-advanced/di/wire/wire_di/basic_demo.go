package main

//go:generate go get github.com/google/wire/cmd/wire@latest
//go:generate wire
func main() {
	// 使用wire实现依赖注入
	e := InitializeEvent()
	e.Start()
}
