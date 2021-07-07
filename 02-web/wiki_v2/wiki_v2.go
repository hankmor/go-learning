package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

const FILE_EXT = ".txt"
const FILE_DIR = "02-web/data/"
const TEMPLATE_DIR = "02-web/wiki_v2/tmpl/"

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
		log.Printf("Handle request error, file may not exists: %v", err.Error())
	} else {
		// 使用html/template包渲染Html内容，忽略错误信息
		renderTemplate(w, "view", p)
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
	renderTemplate(w, "edit", p)
}

// 列表页面
func listHandler(w http.ResponseWriter, r *http.Request) {
	// fs, _ := ioutil.ReadDir(FILE_DIR)
	// TODO 模板循环遍历？？
	// renderTemplate(w, "edit.html", p)
}

// 新增页面
func createHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "create", nil)
}

// 渲染html页面
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(TEMPLATE_DIR + tmpl + ".html")
	t.Execute(w, p)
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
