package goroutine

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	count int64
	sg    sync.WaitGroup
)

func Countn() {
	/*
		n 个 goroutine，每个对 count 加 m 次，结果应该为 n * m，而实际上是小于这个值
	*/

	var n = 1000
	var m = 100
	sg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer sg.Done()
			for i := 0; i < m; i++ {
				count++ // ++操作不是原子操作，同java一样，先从内存读取，再加1，再写回内存，三步不是原子操作
				runtime.Gosched()
			}
		}()
	}

	sg.Wait()
	fmt.Println("expected count: ", n*m, "actual count: ", count) // expected count:  100000 actual count:  96069
}

func Count() {
	sg.Add(2)

	go IncErr()
	go IncErr()

	// 等待
	sg.Wait()
	fmt.Println("count: ", count) // 执行 IncErr 可能输出2，也可能输出4
}

func IncErr() {
	defer sg.Done()
	for i := 0; i < 2; i++ {
		value := count    // 赋值到临时变量
		runtime.Gosched() // 当前 goroutine 从线程退出，并放回到队列，切换 gorountine 是的竞态条件更明显
		value++           // 临时变量加1
		count = value     // 赋值给count
	}
}

// SyncCountn 改进 Countn，可以使用同步
func SyncCountn() {
	var lock sync.Mutex
	var n = 1000
	var m = 100
	sg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer sg.Done()
			for i := 0; i < m; i++ {
				lock.Lock()   // 加锁
				count++       // 执行++操作，现在符合原子性了
				lock.Unlock() // 解锁
				runtime.Gosched()
			}
		}()
	}

	sg.Wait()
	fmt.Println("expected count: ", n*m, "actual count: ", count)
	// expected count:  100000 actual count:  100000
}

// var atomicCount sync.Atom

// AtomicCountn 除了使用sync包的mutex，还可以使用 atomic 包，类似java 的 atomic
func AtomicCountn() {
	var n = 1000
	var m = 100
	sg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer sg.Done()
			for i := 0; i < m; i++ {
				atomic.AddInt64(&count, 1) // 不使用 count++，而是改为原子的增加方法
				runtime.Gosched()
			}
		}()
	}

	sg.Wait()
	fmt.Println("expected count: ", n*m, "actual count: ", count)
}
