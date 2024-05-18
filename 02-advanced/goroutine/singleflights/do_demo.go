package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/sync/singleflight"
)

/*
	singleflight: 抑制重复请求，同时发起的多次相同的请求只有一次会通过，其他请求则等待知道第一次请求通过并直接返回其结果。
    如何才能算重复请求？对于相同的key，如果一个请求在执行且还没有返回，又有其他请求发起，则这些请求就是重复的请求。
*/

var count int32

// 模拟并发请求查询特定 id 的文章，使用 atomic 包保证原子性。
func getArticle(id int) (article string, err error) {
	// 假设这里会对数据库进行调用, 模拟不同并发下耗时不同
	// 使用原子操作保证并发安全，请求次数越多(count越大)，耗时越久
	atomic.AddInt32(&count, 1)
	time.Sleep(time.Duration(count) * time.Millisecond)
	return fmt.Sprintf("article: %d", id), nil
}

// 使用 singleflight 包抑制重复请求，对于相同 id 的 article，只需要真正查询一次数据库，其他请求阻塞等待真正的查询返回后可以直接获得共享数据
func singleflightGetArticle(sg *singleflight.Group, id int) (string, error, bool) {
	// Do 执行函数, 参数为 key 和 fn 函数, 对同一个 key 多次调用的时候，在第一次调用没有执行完的时候
	// 只会执行一次 fn 其他的调用会阻塞住等待这次调用返回
	//
	// fn 返回 interface{} 和 error，也就是说，它内部执行函数时需要返回一个结果和一个 error
	// Do 函数返回 interface{}, error 和 shared 三个值:
	// 1. interface{} 为参数 fn 的第一个返回值
	// 2. error 为参数 fn 的第二个返回值
	// 3. shared 为一个 bool 类型，代表参数 fn 函数的返回值是否共享给了多个调用者，只要有其他调用者获得了共享的数据就返回 true
	v, err, shared := sg.Do(fmt.Sprintf("%d", id), func() (interface{}, error) {
		fmt.Println("getArticle")
		return getArticle(id)
	})
	return v.(string), err, shared
}

func doDemo() {
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
			var res string
			var err error
			var shared bool
			// 直接调用查询方法，请求次数越多，越慢
			// res, err = getArticle(1)
			// 不受请求次数的影响，实际上只会查询一次
			res, err, shared = singleflightGetArticle(sg, 1)
			if err != nil {
				panic(err)
			}
			if res != "article: 1" {
				panic("handle error")
			}
			fmt.Printf("article: %s, shared: %v\n", res, shared)
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
