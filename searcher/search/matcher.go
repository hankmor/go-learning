package search

import (
	"fmt"
	"log"
)

// 匹配器，定义了搜索的结果和搜索的接口标准，并可以显示搜索的结果内容、
// 对于一个数据源 Feed ，匹配器都可以并发地通过 goroutine 执行

// init方法时go定义的规范，保证其能够在main函数执行前调用
func init() {
	// 申明一个默认匹配器变量，并注册
	var defaultMatcher defaultMatcher
	Register("default", defaultMatcher)
}

// Result 定义搜索的结果结构
type Result struct {
	Field   string // 结果所在的字段名称，包括 Title 和 Description，分别表示 在"标题"和"描述"中搜索到
	Content string // 内容，Field 为 Title 是显示完整的标题，为 Description 时显示完整的描述信息
}

// Matcher 匹配器接口，定义标准的 Search 方法
type Matcher interface {
	// Search 搜索方法
	//
	// key: 搜索的关键字，feed: 从哪个数据源中搜索
	// 返回 Result 的切片和错误信息
	Search(key string, feed *Feed) ([]*Result, error)
}

// Match 匹配逻辑，可以并发的执行，需要将结果输出到能够处理并发数据共享的通道 channel 中，然后显示结果的 Display
// 方法可以从通道中读取数据并输出
//
// matcher: 匹配器, key: 搜索关键字, feed: 数据源, results: 数据结果通道，将匹配器搜索的结果输出到通道中
func Match(matcher Matcher, key string, feed *Feed, results chan<- *Result) {
	rets, err := matcher.Search(key, feed)
	if err != nil {
		// 打印错误信息，然后返回，这里由于是并发执行，不能 log.Fatal()
		log.Println(err)
		return
	}

	// 将搜索的结果输出到通道中，然后 Display 方法就可以获取到通道中的内容并显示
	for _, ret := range rets {
		results <- ret
	}
}

// Display 显示搜索到的结果信息，方法参数为 通道中的结果数据
func Display(results chan *Result) {
	// 遍历通道中的结果，并打印
	//
	// 如果通道中没有写入任何数据，此时将阻塞
	// 如果通道关闭，则该方法将返回
	for result := range results {
		// 按照 Field: \n Content \n\n 的格式显示
		fmt.Printf("%s: \n%s\n\n", result.Field, result.Content)
	}
}

// 默认的匹配器
type defaultMatcher struct {
}

// Search 默认的匹配器下的搜索方法实现：直接返回nil
func (matcher defaultMatcher) Search(key string, feed *Feed) ([]*Result, error) {
	return nil, nil
}
