package main

import "wire_demo"

func main() {
	// 没有依赖注入，需要单独创建实例并传递参数
	msg := wire_demo.NewMessage("hank") // 创建msg
	g := wire_demo.NewGreeter(msg)      // 创建问候者
	e := wire_demo.NewEvent(g)          // 创建问候事件
	e.Start()
}
