package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

const (
	// 文件后缀
	FileExt = ".txt"
	// 文件目录
	FileDir = "02-web/data/"
	// 模板目录
	TemplateDir = "02-web/wiki_v2/tmpl/"
)

// 运用模板缓存，在程序初始化时就加载模板，而不是每次调用方法再去加载一次
// Must为一个panic，保证必须成功加载模板，否则程序退出
var templates = template.Must(template.ParseFiles(TemplateDir+"create.html", TemplateDir+"edit.html", TemplateDir+"list.html", TemplateDir+"view.html"))

// 验证url中的文件title，只能为字母和数字的组合
var validPath = regexp.MustCompile("^/(edit|view)/([a-zA-Z0-9]+)$")

// 验证保存时的Title
var validTitle = regexp.MustCompile("^([a-zA-Z0-9]+)$")

// wiki页面，存储数据结构
type Page struct {
	Title string // 标题
	Body  []byte // 内容，类型为slice
}

// 存储文件
// func save(p *Page) error {
func (p *Page) save() error {
	// 文件名，后缀为txt
	filename := FileDir + p.Title + FileExt
	// 写入文件，成功则无任何异常，否则返回异常信息
	// 0600参数：八进制，表示当前用户有用创建文件的读写权限
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// 加载页面
func loadPage(title string) (*Page, error) {
	filename := FileDir + title + FileExt
	// 读取文件内容，第二个返回值为异常信息
	body, err := ioutil.ReadFile(filename)
	// 有异常，则返回异常信息
	if err != nil {
		return nil, err
	}
	// 返回Page
	return &Page{Title: title, Body: body}, nil
}

// ============================================
// 增改查http请求处理器
// 优化wikie_v2，使用函数参数和闭包来减少重复代码
// ============================================

// 页面渲染请求处理器
func viewHandler(w http.ResponseWriter, req *http.Request, title string) {
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
	// 验证title
	m := validTitle.FindStringSubmatch(title)
	if m == nil {
		http.Error(w, "Invalid Title", http.StatusNotAcceptable)
		return
	}
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
func editHandler(w http.ResponseWriter, req *http.Request, title string) {
	p, err := loadPage(title)
	// 文件不存在，则创建一个Page
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

// 列表页面
func listHandler(w http.ResponseWriter, r *http.Request) {
	fs, _ := ioutil.ReadDir(FileDir)
	// 创建slice，长度与fs一致
	ps := make([]Page, len(fs))
	// 模板循环遍历
	for i, fi := range fs {
		title := fi.Name()[:len(fi.Name())-len(FileExt)]
		// 给slice赋值
		ps[i] = Page{title, nil}
	}
	renderTemplate(w, "list", ps)
}

// 新增页面
func createHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "create", nil)
}

// 渲染html页面
func renderTemplate(w http.ResponseWriter, tmpl string, p interface{}) {
	// 调用全局templates.ExecuteTemplate方法，执行某一个模板
	// 注意解析后模板名称不带目录
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		// 有异常，返回500
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// 创建Handler，返回一个HandlerFunc
// 使用函数作为参数，内部闭包调用
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	// 闭包，返回函数，内部调用fn参数
	return func(w http.ResponseWriter, r *http.Request) {
		// 获取title
		m := validPath.FindStringSubmatch(r.URL.Path)
		// 没有匹配，说明文件规则不合法，直接返回404
		if m == nil {
			http.NotFound(w, r)
			return
		}
		// 调用fn
		fn(w, r, m[2])
	}
}

func main() {
	// 列表
	http.HandleFunc("/", listHandler)
	// 新增
	http.HandleFunc("/create", createHandler)
	// 查看，创建handler
	http.HandleFunc("/view/", makeHandler(viewHandler))
	// 保存
	http.HandleFunc("/save", saveHandler)
	// 编辑，创建handler
	http.HandleFunc("/edit/", makeHandler(editHandler))
	// 启动服务器
	log.Fatal(http.ListenAndServe(":8080", nil))
}
