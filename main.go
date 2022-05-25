package main

import (
	"github.com/koobyte/go-learning/basic"
	"github.com/koobyte/go-learning/fuzz"
	"github.com/koobyte/go-learning/generic"
	"github.com/koobyte/go-learning/gin"
	"github.com/koobyte/go-learning/gowiki"
	"github.com/koobyte/go-learning/howto"
	"github.com/koobyte/go-learning/json"
	"github.com/koobyte/go-learning/oop"
	"github.com/koobyte/go-learning/searcher"
)

func main() {
	// 入门
	howto.Run()
	// 基本语法
	basic.Run()
	// 字符串
	fuzz.Run()
	// 面向对象
	oop.Run()
	// 泛型
	generic.Run()
	// json处理
	json.Run()
	// gowiki应用
	gowiki.Run()
	// rss搜索订阅应用
	searcher.Run()
	// gin框架示例
	gin.Run()
}
