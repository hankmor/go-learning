package main

import (
	// 先导入matcher包，不使用会报错，所以需要加上_，目的是为了调用 init 方法注册匹配器
	_ "goinaction/c02/matcher"
	"goinaction/c02/search"
)

// 主程序入口
func main() {
	search.Run("president")
}
