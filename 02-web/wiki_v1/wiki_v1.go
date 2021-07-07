package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

const FILE_EXT = ".txt"
const FILE_DIR = "02-web/files/"

// wiki页面，存储数据结构
type Page struct {
	Title string // 标题
	Body  []byte // 内容，类型为slice
}

// 存储文件
// func save(p *Page) error {
func (p *Page) save() error {
	// 文件名，后缀为txt
	filename := FILE_DIR + p.Title + FILE_EXT
	// 写入文件，成功则无任何异常，否则返回异常信息
	// 0600参数：八进制，表示当前用户有用创建文件的读写权限
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// 加载页面
func loadPage(title string) (*Page, error) {
	filename := FILE_DIR + title + FILE_EXT
	// 读取文件内容，第二个返回值为异常信息
	body, err := ioutil.ReadFile(filename)
	// 有异常，则返回异常信息
	if err != nil {
		return nil, err
	}
	// 返回Page
	return &Page{Title: title, Body: body}, nil
}

// ===========================
// 增改查http请求处理器
// ===========================

// 页面渲染请求处理器
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
		fmt.Fprintf(w, "<a href=\"/list\">返回列表</a> <a href=\"/edit/"+title+"\">编辑</a>"+
			"<h1>%s</h1><div>%s</div>", p.Title, p.Body)
	}
}

// 保存
func saveHandler(w http.ResponseWriter, req *http.Request) {
	title := req.FormValue("title")
	// 这里的body取出来是字符串
	body := req.FormValue("body")
	// 创建Page，body转为[]byte
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		fmt.Fprintf(w, "<div style='color: red;'>%s</div>", "保存文件出现异常！")
		log.Printf("Save file error, title : %s, body: %s: ,error: %v", title, body, err.Error())
	}
	// 保存成功就返回到list页面
	http.Redirect(w, req, "/list", http.StatusFound)
}

// 编辑页面
func editHandler(w http.ResponseWriter, req *http.Request) {
	title := req.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	// 文件不存在，则创建一个Page
	if err != nil {
		p = &Page{Title: title}
	}
	// 显示Html内容，这里是硬编码
	fmt.Fprintf(w, `
		<h1>编辑 %s</h1>
		<form action="/save" method="POST">
			<label>标题</label><input name="title" value="%s"><br>
			<label>内容</label><textarea rows="20" cols="80" name="body">%s</textarea><br>
			<input type="submit" value="保存">
		</form>
	`, p.Title, p.Title, p.Body)
}

// 列表页面
func listHandler(w http.ResponseWriter, r *http.Request) {
	fs, _ := ioutil.ReadDir(FILE_DIR)
	s := "<h1>文件列表</h1> <a href=\"/create\">新增</a> <ul>"
	for _, fi := range fs {
		title := fi.Name()[:len(fi.Name())-len(FILE_EXT)]
		s += "<li>" + title
		s += " <a href=\"/view/" + title + "\">查看</a>"
		s += " <a href=\"/edit/" + title + "\">编辑</a>"
		s += "</li>"
	}
	s += "</ul>"
	fmt.Fprintf(w, s)
}

// 新增页面
func createHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
		<h1>新增</h1>
		<form action="/save" method="POST">
			<label>标题</label><input name="title"><br>
			<label>内容</label><textarea rows="20" cols="80" name="body"></textarea><br>
			<input type="submit" value="保存">
		</form>
	`)
}

func main() {
	// 列表
	http.HandleFunc("/", listHandler)
	// 新增
	http.HandleFunc("/create", createHandler)
	// 查看
	http.HandleFunc("/view/", viewHandler)
	// 保存
	http.HandleFunc("/save", saveHandler)
	// 编辑
	http.HandleFunc("/edit/", editHandler)
	// 启动服务器
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// 测试方法

func testSaveAndLoad(t *testing.T) {
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
