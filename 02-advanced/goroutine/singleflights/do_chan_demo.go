package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/sync/singleflight"
)

// 除了 Do() 直接返回结果外，singleflight 还提供了

func singleflightGetArticleChan(sg *singleflight.Group, id int) <-chan singleflight.Result {
	// DoChan 函数与 Do 函数类似，只是将 Do 的结果封装为 Result 对象，并返回为一个持有它的 chan
	cr := sg.DoChan(fmt.Sprintf("%d", id), func() (interface{}, error) {
		fmt.Println("getArticle")
		return getArticle(id)
	})
	return cr
}

func doChanDemo() {
	time.AfterFunc(1*time.Second, func() {
		atomic.AddInt32(&count, -count)
	})

	var (
		wg  sync.WaitGroup
		now = time.Now()
		n   = 10
		sg  = &singleflight.Group{}
	)

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			// 直接调用查询方法，请求次数越多，越慢
			// res, err = getArticle(1)
			// 不受请求次数的影响，实际上只会查询一次
			ch := singleflightGetArticleChan(sg, 1)
			r := <-ch
			fmt.Printf("article: %s, shared: %v\n", r.Val, r.Shared)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("同时发起 %d 次请求， 真正查询次数: %d, 耗时: %s", n, count, time.Since(now))
	// singleflightGetArticle:
	// 同时发起 10 次请求，真正查询次数: 1, 耗时: 1.785056ms
	// getArticle:
	// 同时发起 10 次请求，真正查询次数: 10, 耗时: 10.392278s
}
