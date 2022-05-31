package basic

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
)

func ErrMain() {
	/*
		Go 内置了 error 接口，如下：
		type error interface {
			Error() string
		}
	*/
	// fmt.Println("error demo")
	// errorDemo()
	//
	// // defer 可以保证在函数正常或异常返回时都能够执行，常用来清理资源
	// fmt.Println("defer demo")
	// err := deferDemo()
	// if err != nil {
	// 	log.Println("err:", err)
	// }
	//
	// fmt.Println("panic and recover")
	// panicAndRecover()

	println("====")
	// var wg sync.WaitGroup
	// wg.Add(10)
	for i := 0; i < 10; i++ {
		// makeErr()
		// 不能捕获
		// if err := recover(); err != nil {
		// 	continue
		// }

		// 使用匿名函数执行
		// go func() {
		func() {
			defer func() {
				// wg.Done()
				if err := recover(); err != nil {
					fmt.Printf("error: %v\n", err)
				}
			}()
			makeErr()
		}()
	}
	// wg.Wait()
}

func makeErr() {
	switch rand.Intn(3) {
	case 0:
		panic("zero...")
	case 1:
		println("number is 1")
	case 2:
		println("number is 2")
	}
}

// GO语言引入了 panic() 和 recover() 两个函数来报告和处理运行时错误
// 当函数内调用 panic() 时，函数立即终止执行，但是 defer 语句仍然能够执行
// 此时，会逐层向上执行panic流程，直至所属的goroutine中所有正在执行的函数被终止
//
// recover()函数用于终止错误处理流程。一般情况下，recover()应该在一个使用defer
// 关键字的函数中执行以有效截取错误处理流程。如果没有在发生异常的goroutine中明确调用恢复
// 过程（使用recover关键字），会导致该goroutine所属的进程打印异常信息后直接退出。
func panicAndRecover() {
	invokeRecoverWhenNoError()
	funcA()
}

func invokeRecoverWhenNoError() {
	fmt.Println("invokeRecoverWhenNoError")
	defer func() {
		if r := recover(); r != nil {
			log.Printf("invokeRecoverWhenNoError Runtime error caught: %v", r)
		}
	}()
}

func funcA() {
	fmt.Println("func a")
	defer fmt.Println("func a defer")
	funcB()
	defer func() {
		if r := recover(); r != nil {
			log.Printf("func a Runtime error caught: %v", r)
		}
	}()
	panic("func a panic")
}

func funcB() {
	fmt.Println("func b")
	defer fmt.Println("func b defer")
	funcC()
	defer func() {
		if r := recover(); r != nil {
			log.Printf("func b Runtime error caught: %v", r)
		}
	}()
	panic("func b panic")
}

func funcC() {
	fmt.Println("func c")
	defer fmt.Println("func c defer")
	defer func() {
		if r := recover(); r != nil {
			log.Printf("func c Runtime error caught: %v", r)
		}
	}()
	panic("func c panic")
}

func errorDemo() {
	pathErr := PathError{
		Op:   "CREATE",
		Path: "/usr/home",
		Err:  errors.New("path error"),
	}
	err := pathErr.Error()
	fmt.Println(err)
}

// PathError 申明一个错误结构类型
type PathError struct {
	Op   string
	Path string
	Err  error
}

// Error 打印错误方法
func (e *PathError) Error() string {
	return e.Op + " " + e.Path + " " + e.Err.Error()
}

func deferDemo() error {
	fmt.Println("enter function")
	// 返回前执行
	// defer语句的调用是遵照
	// 先进后出的原则，即最后一个defer语句将最先被执行
	defer fmt.Println("defer1...")
	defer fmt.Println("defer2...")

	fmt.Println("before return")
	// log.Fatal("make a fatal")
	return errors.New("make a error")
	/*
		enter function
		before return
		defer2...
		defer1...
	*/
}
