package search

import (
	"log"
	"sync"
)

// 创建一个存储 Matcher 的map，由 Register 方法注册到其中
var matchers = make(map[string]Matcher)

// Run 搜索器启动方法，搜索给定参数的内容
func Run(key string) {
	// 读取json数据源文件
	feeds, err := ParseJson()
	if err != nil {
		log.Fatal("Parse json file error: ", err)
	}

	// 创建结果通道，启动 goroutine 后将其传入并接收搜索结果
	results := make(chan *Result)

	// WaitGroup 是一个信号量，当 goroutine 并行执行时通过它来计数
	// WaitGroup.wait 方法阻塞，知道计数器减为0
	// 每完成一个并行任务需要调用WaitGroup.Done方法是计数器减1
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(feeds)) // 初始化信号量长度

	// 遍历feeds切片，返回的第一个参数为元素下标，这里忽略
	for _, feed := range feeds {
		// 获取 Matcher
		matcher, exists := matchers[feed.Type]
		// 不存在，则使用默认的 Matcher
		if !exists {
			log.Println("Not found matcher: ", feed.Type, "Using \"default\" instead but can not search anything.")
			matcher = matchers["default"]
		}

		// 并行执行，创建一个匿名函数，将当前 matcher 和 feed 作为参数传入
		// 其实这里是一个闭包，不采用匿名函数传参的方式，读到的matcher 和 feed 可能只有最后一个
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, key, feed, results)
			waitGroup.Done()
		}(matcher, feed)
	}

	// 创建并执行一个匿名函数，开启一个并行任务来检测所有搜索任务是否都完成
	go func() {
		// 阻塞知道waitGroup的计数器为0
		waitGroup.Wait()

		// 关闭通道
		close(results)
	}()

	// 显示结果，results通道有数据则输出，没有输出则阻塞，直到该通道关闭
	Display(results)
}

func Register(feedType string, matcher Matcher) {
	existsMatcher, exists := matchers[feedType]
	if exists {
		log.Fatal(feedType, existsMatcher, "Matcher already exists.")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}
