package main

func main() {
	// 使用wire实现依赖注入
	name := "jason"
	e := InitializeEvent(name)
	e.Start()
}
