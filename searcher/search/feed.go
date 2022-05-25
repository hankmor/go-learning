package search

import (
	"encoding/json"
	"log"
	"os"
)

// 数据提供者，通过读取并解析 data/data.json 文件类为 search 提供数据源信息

// jsonFile 文件存储位置常量
const jsonFile = "searcher/src/c02/data/data.json"

// Feed 提供的数据结构
type Feed struct {
	// 映射json文件的结构到具体的属性上，方便json模块可以直接解析
	Site string `json:"site"`
	Link string `json:"link"`
	Type string `json:"type"`
}

// ParseJson 解析json文件到 Feed 结构的切片中
func ParseJson() ([]*Feed, error) {
	// 读取文件
	file, err := os.Open(jsonFile)
	// 如果读取文件失败，直接返回nil和错误信息
	if err != nil {
		return nil, err
	}
	// 读取文件必须要关闭资源
	// 这里在一个闭包里边处理错误，如果关闭文件出现错误，直接打印错误信息
	// 也可以不处理错误，直接关闭文件：defer file.Close()
	// defer的作用：函数正常或异常返回时都会执行，保证资源的正常关闭，此外，打开资源和关闭资源的代码可以就近编写，增加代码可读性
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	// 开始解析json文件
	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)
	return feeds, err
}
