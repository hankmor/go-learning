package main

import (
	"fmt"
	"time"
)

func main() {
	// 当前时间
	now := time.Now()
	fmt.Printf("now: %v\n", now) // now: 2022-05-31 14:01:58.68345 +0800 CST m=+0.000064237

	// 计算

	// 加上2个小时后
	tm := now.Add(2 * time.Hour)
	fmt.Printf("time: %v\n", tm) // time: 2022-05-31 16:01:58.68345 +0800 CST m=+7200.000064237

	d := tm.Sub(now)
	fmt.Printf("time: %v\n", d.Hours()) // 2
	// 相同
	println(tm.Equal(now))  // false
	println(now.After(tm))  // false
	println(now.Before(tm)) // true

	tm = tm.AddDate(0, 1, 0)
	fmt.Printf("time: %v\n", tm) // time: 2022-07-01 16:01:58.68345 +0800 CST

	// 格式化
	layout := "2006-01-02 15:04:05.000 -0700" // 数字不能更改，含义看format.go的Layout文档
	println("format: ", tm.Format(layout))
	// 解析
	tm, _ = time.Parse(layout, tm.Format(layout))
	fmt.Printf("parse: %v\n", tm)

	// during

	sec := 10
	println(time.Duration(sec) * time.Second) // 10秒对应的纳秒值
	hour := 2
	println(time.Duration(hour) * time.Hour) // 2小时的纳秒数

	d1 := time.Duration(60) * time.Second
	d2 := time.Duration(1) * time.Minute
	println(d1 - d2) // 0
}
