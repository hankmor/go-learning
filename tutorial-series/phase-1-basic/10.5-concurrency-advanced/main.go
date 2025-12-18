package main

import (
	"fmt"
	"sync"
	"time"
)

// ==================== 1. sync.WaitGroup 示例 ====================

func workerWaitGroup(id int, wg *sync.WaitGroup) {
	defer wg.Done() // 函数结束时调用 Done()

	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func demoWaitGroup() {
	fmt.Println("\n=== WaitGroup Demo ===")
	var wg sync.WaitGroup

	// 启动 5 个 worker
	for i := 1; i <= 5; i++ {
		wg.Add(1) // 每启动一个 goroutine，计数器 +1
		go workerWaitGroup(i, &wg)
	}

	wg.Wait() // 等待所有 worker 完成
	fmt.Println("All workers completed")
}

// ==================== 2. sync.Mutex 示例 ====================

type SafeCounter struct {
	mu    sync.Mutex
	count int
}

func (c *SafeCounter) Inc() {
	c.mu.Lock() // 加锁
	c.count++
	c.mu.Unlock() // 解锁
}

func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func demoMutex() {
	fmt.Println("\n=== Mutex Demo ===")
	counter := SafeCounter{}
	var wg sync.WaitGroup

	// 启动 1000 个 goroutine 同时增加计数器
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Inc()
		}()
	}

	wg.Wait()
	fmt.Println("Final count:", counter.Value()) // 输出：1000
}

// ==================== 3. sync.RWMutex 示例 ====================

type Cache struct {
	mu   sync.RWMutex
	data map[string]string
}

func (c *Cache) Get(key string) string {
	c.mu.RLock() // 读锁
	defer c.mu.RUnlock()
	return c.data[key]
}

func (c *Cache) Set(key, value string) {
	c.mu.Lock() // 写锁
	defer c.mu.Unlock()
	c.data[key] = value
}

func demoRWMutex() {
	fmt.Println("\n=== RWMutex Demo ===")
	cache := Cache{data: make(map[string]string)}

	// 写入数据
	cache.Set("name", "Go")

	// 多个 goroutine 同时读取（不会阻塞）
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Reader %d: %s\n", id, cache.Get("name"))
		}(i)
	}

	wg.Wait()
}

// ==================== 4. sync.Once 示例 ====================

var (
	instance *Singleton
	once     sync.Once
)

type Singleton struct {
	data string
}

func GetInstance() *Singleton {
	once.Do(func() {
		fmt.Println("Creating singleton instance")
		instance = &Singleton{data: "singleton"}
	})
	return instance
}

func demoOnce() {
	fmt.Println("\n=== Once Demo ===")
	var wg sync.WaitGroup

	// 多个 goroutine 同时获取实例
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			s := GetInstance()
			fmt.Println(s.data)
		}()
	}

	wg.Wait()
	// "Creating singleton instance" 只会打印一次
}

// ==================== 5. sync.Cond 示例 ====================

func demoCond() {
	fmt.Println("\n=== Cond Demo ===")
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	ready := false

	// 等待者
	go func() {
		mu.Lock()
		for !ready {
			cond.Wait() // 等待条件满足
		}
		fmt.Println("Condition met, proceeding...")
		mu.Unlock()
	}()

	// 通知者
	time.Sleep(time.Second)
	mu.Lock()
	ready = true
	cond.Signal() // 通知一个等待的 goroutine
	mu.Unlock()

	time.Sleep(time.Second)
}

// ==================== 主函数 ====================

func main() {
	fmt.Println("Go 并发进阶示例")

	// 运行各个示例
	demoWaitGroup()
	demoMutex()
	demoRWMutex()
	demoOnce()
	demoCond()

	fmt.Println("\n所有示例运行完毕")
}
