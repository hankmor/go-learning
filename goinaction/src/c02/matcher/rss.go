package matcher

import (
	"c02/search"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
)

// 基于 rss 的匹配器，用于下载 rss 的 xml 然后搜索其中的 title 和 description 内容
//
// 要解析 xml 文件，需要引入 xml 包，类似 json 的解析方式
// 示例rss：https://feeds.npr.org/1001/rss.xml

func init() {
	var rssMatcher rssMatcher
	search.Register("rss", rssMatcher)
}

// item xml中的 item 节点结构
type item struct {
	XmlName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Link        string   `xml:"link"`
}

// image Xml 中的 image 节点结构
type image struct {
	XmlName xml.Name `xml:"image"`
	Url     string   `xml:"url"`
	Title   string   `xml:"title"`
	Link    string   `xml:"link"`
}

// channel rss 的 channel 节点结构
type channel struct {
	XmlName     xml.Name `xml:"channel"`
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
	Description string   `xml:"description"`
	Language    string   `xml:"language"`
	Image       image    `xml:"image"`
	Item        []item   `xml:"item"` // item 有多个
}

// rss xml 的 rss 节点结构
type rss struct {
	XmlName xml.Name `xml:"rss"`
	Channel channel  `xml:"channel"`
}

// rssMatcher 定义 rss 匹配器，实现了 search.Matcher 接口
type rssMatcher struct {
}

// Search 实现 search.Matcher 接口
// 实现逻辑：解析 rss 的xml文件，然后匹配 item 的 title 和 description，看是否匹配搜索关键字 key，能匹配则输出则返回结果
func (matcher rssMatcher) Search(key string, feed *search.Feed) ([]*search.Result, error) {
	log.Printf("Searching in FeedType: %s, FeedSite: %s, FeedUrl: %s \n", feed.Type, feed.Site, feed.Link)
	// 解析xml
	rss, err := matcher.parse(feed)
	if err != nil {
		return nil, err
	}

	// 循环搜索，结果存储在 results 变量中
	var results []*search.Result
	for _, item := range rss.Channel.Item {
		// 搜索 title
		title := item.Title
		match, err := regexp.MatchString(key, title)
		// 有错误，继续搜索
		if err != nil {
			log.Println("Search in Title error: ", err)
			continue
		}
		if match {
			results = append(results, &search.Result{
				Field:   "title",
				Content: title,
			})
		}

		// 搜索 description
		desc := item.Description
		match, err = regexp.MatchString(key, desc)
		if err != nil {
			log.Println("Search in Description error: ", err)
			continue
		}
		if match {
			results = append(results, &search.Result{
				Field:   "description",
				Content: desc,
			})
		}
	}
	// 返回结果
	return results, nil
}

// parse 解析 xml
func (matcher rssMatcher) parse(feed *search.Feed) (*rss, error) {
	// 判断 link 不为空
	url := feed.Link
	if url == "" {
		return nil, errors.New("rss feed URL must not be empty")
	}

	// 发送http请求获取xml数据内容
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	// 关闭 response
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("close responseBody error: ", err)
		}
	}(resp.Body)

	// http 请求失败，没有返回 200，直接返回
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Response Status Error: %d\n", resp.StatusCode)
	}

	// 解析xml，如果有错误，不处理错误而是直接返回错误
	var rss rss
	err = xml.NewDecoder(resp.Body).Decode(&rss)
	return &rss, err
}
