package gowiki

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
)

const SavePath = "gowiki/wiki/"

// 全局变量，一次解析模板，避免多次重复解析
var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

// 正则表达式验证 url 输入
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9\\-_]+)$")

func init() {
	_, err := os.Stat(SavePath)
	if err != nil {
		// 错误报告的是文件不存在，创建文件夹
		if os.IsNotExist(err) {
			os.Mkdir(SavePath, os.ModePerm)
		}
	}
}

// Page 页面
type Page struct {
	Title string
	Body  []byte
}

// 保存文件
func (p *Page) save() error {
	filename := SavePath + p.Title + ".txt"
	// 0600: 建文件时应仅对当前用户具有读写权限
	return os.WriteFile(filename, p.Body, 0600)
}

// 读取页面
func loadPage(title string) (*Page, error) {
	filename := SavePath + title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// 读取：处理 /view/ 开头的url，读取文件到页面
func viewHandler(w http.ResponseWriter, r *http.Request) {
	// URL中的文件名
	title, _ := getTitle(w, r)
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

// 修改
func editHandler(w http.ResponseWriter, r *http.Request) {
	title, _ := getTitle(w, r)
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	// 模板
	renderTemplate(w, "edit", p)
}

// 保存
func saveHandler(w http.ResponseWriter, r *http.Request) {
	title, _ := getTitle(w, r)
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

// 渲染模板
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	// 性能低，每次都要解析模板
	// t, handleerr := template.ParseFiles(tmpl + ".html")
	// if handleerr != nil {
	//	http.Error(w, handleerr.Error(), http.StatusInternalServerError)
	//	return
	// }
	// handleerr = t.Execute(w, p)
	// if handleerr != nil {
	//	http.Error(w, handleerr.Error(), http.StatusInternalServerError)
	// }
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// 获取标题，验证输入的合法性
func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("invalid Page Title")
	}
	return m[2], nil // 标题是第二个子表达式
}

// 使用闭包优化

func viewHandler1(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}
func editHandler1(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}
func saveHandler1(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

// 参数为func，返回一个func
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		// m[2] 为 title
		fn(w, r, m[2])
	}
}

func Run() {
	// test
	// p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	// handleerr := p1.save()
	// if handleerr != nil {
	//	return
	// }
	// p2, _ := loadPage("TestPage")
	// fmt.Println(string(p2.Body))

	// 访问页面时URL: /view/文件名
	// http.HandleFunc("/view/", viewHandler)
	// http.HandleFunc("/edit/", editHandler)
	// http.HandleFunc("/save/", saveHandler)

	// 使用闭包优化后，也可以使用匿名函数
	// http.HandleFunc("/view/", func(w http.ResponseWriter, r *http.Request) {
	//	title, _ := getTitle(w, r)
	//	p, handleerr := loadPage(title)
	//	if handleerr != nil {
	//		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
	//		return
	//	}
	//	renderTemplate(w, "view", p)
	// })
	http.HandleFunc("/view/", makeHandler(viewHandler1))
	http.HandleFunc("/edit/", makeHandler(editHandler1))
	http.HandleFunc("/save/", makeHandler(saveHandler1))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
