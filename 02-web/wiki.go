package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

const FILE_EXT = ".txt"

// wiki页面，存储数据结构
type Page struct {
	Title string // 标题
	Body  []byte // 内容，类型为slice
}

// 存储文件
// func save(p *Page) error {
func (p *Page) save() error {
	// 文件名，后缀为txt
	filename := p.Title + FILE_EXT
	// 写入文件，成功则无任何异常，否则返回异常信息
	// 0600参数：八进制，表示当前用户有用创建文件的读写权限
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// 加载页面
func loadPage(title string) (*Page, error) {
	filename := title + FILE_EXT
	// 读取文件内容，第二个返回值为异常信息
	body, err := ioutil.ReadFile(filename)
	// 有异常，则返回异常信息
	if err != nil {
		return nil, err
	}
	// 返回Page
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, req *http.Request) {
	// URL必须以/view/开始，后边为文件名称，所以去掉/view/之后为文件名
	title := req.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	// 存在错误，则前台返回错误信息
	if err != nil {
		// 忽略错误信息
		fmt.Fprintf(w, "<div style='color: red;'>%s</div>", "访问的文件不存在！")
		log.Printf("Handle request error, file may not exists: %v", err.Error())
	} else {
		// 渲染为Html内容，忽略错误信息
		fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
	}
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// 测试方法

func TestSaveAndLoad(t *testing.T) {
	content := "This is a sample page."
	p1 := &Page{Title: "TestPage", Body: []byte(content)}
	err := p1.save()
	if err != nil {
		t.Fatalf("Save file \"TestPage.txt\" failed: %v", err)
	}
	p2, _ := loadPage("TestPage")
	// 内容的byte切片转换为string
	ret := string(p2.Body)
	if ret != content {
		t.Fatalf("Unexpected content %v, want %v:", ret, content)
	}
}
