package main

import "fmt"

// Go接口都是非侵入性的，一个类（struct）只要实现了接口的所有方法，就说这个类实现了该接口，其他无任何限制
// 鸭子理论：只要它看起来像鸭子，那它可能就是鸭子

// IFile 定义接口
type IFile interface {
	Read(buf []byte) (n int, err error)
	Write(buf []byte) (n int, err error)
	Close() error
}

// IReader 读取接口
type IReader interface {
	Read(buf []byte) (n int, err error)
}

// IWriter 写入接口
type IWriter interface {
	Write(buf []byte) (n int, err error)
}

// ICloser 关闭接口
type ICloser interface {
	Close() error
}

// Printer 打印接口
type Printer interface {
	Print(content string)
}

// File File类，实现了接口的方法，未实现Printer
type File struct {
}

func (file *File) Read(buf []byte) (n int, err error) {
	fmt.Println("Read Not Implementation")
	return 0, nil
}
func (file *File) Write(buf []byte) (n int, err error) {
	fmt.Println("Write Not Implementation")
	return 0, nil
}
func (file *File) Close() (err error) {
	fmt.Println("Close Not Implementation")
	return nil
}

func main() {
	var file *File = new(File)
	fmt.Printf("%T \n", file)
	file.Read(nil)
	file.Write(nil)
	file.Close()
	var file1 IFile = file
	fmt.Printf("%T \n", file1, file1)
	file.Read(nil)
	file.Write(nil)
	file.Close()
	var file2 IReader = file
	fmt.Printf("%T \n", file2)
	file2.Read(nil)
	var file3 IWriter = file
	fmt.Printf("%T \n", file3)
	file3.Write(nil)
	var file4 ICloser = file
	fmt.Printf("%T \n", file4)
	file4.Close()

	// File未实现Print接口，编译失败
	// var file5 Printer = file

	// 接口查询
	closer, ok := file4.(ICloser)
	fmt.Printf("%T, ok = %v \n", closer, ok) // *main.File, ok = true

	// 类型查询，实例.(type) 可以查询实例的类型，只能用在switch语句中
	switch t := file1.(type) {
	case ICloser:
		fmt.Printf("type is %T \n", t) // type is *main.File
	case IReader:
		fmt.Println("type is IReader")
	case IWriter:
		fmt.Println("type is IWriter")
	}

	// 类型查询示例
	typeQuery(100, "helloGo", file2, new(Printer))
	/*
		arg 0 = 100, type is int
		arg 1 = helloGo type is string
		Read Not Implementation
		arg 3 = 0xc000010230 type is *main.Printer
	*/
}

func typeQuery(args ...interface{}) {
	for idx, arg := range args {
		switch t := arg.(type) {
		case int:
			fmt.Printf("arg %d = %v, type is int \n", idx, arg)
		case string:
			fmt.Printf("arg %d = %v type is string \n", idx, arg)
		default:
			if v, ok := arg.(IReader); ok {
				v.Read(nil)
			} else {
				fmt.Printf("arg %d = %v type is %T \n", idx, arg, t)
			}
		}
	}
}
