package main

import (
	"fmt"
	"golang.org/x/sync/singleflight"
	"sync"
	"sync/atomic"
	"time"
)

/*
	singleflight: 限流，多次请求需要等待前边的请求执行完成才能继续请求
*/

var count int32

func getArticle(id int) (article string, err error) {
	// 假设这里会对数据库进行调用, 模拟不同并发下耗时不同
	// 使用原子操作保证并发安全，请求次数越多(count越大)，耗时越久
	atomic.AddInt32(&count, 1)
	time.Sleep(time.Duration(count) * time.Millisecond)
	return fmt.Sprintf("article: %d", id), nil
}

func singleflightGetArticle(sg *singleflight.Group, id int) (string, error) {
	// Do 执行函数, 对同一个 key 多次调用的时候，在第一次调用没有执行完的时候
	// 只会执行一次 fn 其他的调用会阻塞住等待这次调用返回
	// v, handleerr 是传入的 fn 的返回值
	// shared 表示是否真正执行了 fn 返回的结果，还是返回的共享的结果
	v, err, _ := sg.Do(fmt.Sprintf("%d", id), func() (interface{}, error) {
		return getArticle(id)
	})
	return v.(string), err
}

func main() {
	time.AfterFunc(1*time.Second, func() {
		atomic.AddInt32(&count, -count)
	})

	var (
		wg  sync.WaitGroup
		now = time.Now()
		n   = 1000
		// sg  = &singleflight.Group{}
	)

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			// res, _ := singleflightGetArticle(sg, 1)
			res, _ := getArticle(1)
			if res != "article: 1" {
				panic("handleerr")
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("同时发起 %d 次请求，耗时: %s", n, time.Since(now))
	// singleflightGetArticle:
	// 同时发起 1000 次请求，耗时: 3.785056ms
	// getArticle:
	// 同时发起 1000 次请求，耗时: 1.002278214s
}
