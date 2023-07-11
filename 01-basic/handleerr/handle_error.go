package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	badProcessErr()
	fmt.Println("======")
	betterProcessErr()
}

func badProcessErr() {
	var err error
	err = fn1()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = fn2()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = fn3()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func betterProcessErr() {
	h := &handler{}
	h.handle(fn1)
	h.handle(fn2)
	h.handle(fn3)
	if h.err != nil {
		fmt.Println(h.err)
	}
}

type handler struct {
	err error
}

func (h *handler) handle(fn func() error) {
	if h.err != nil {
		return
	}
	h.err = fn()
}

func fn1() error {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(10)
	if i%3 == 0 {
		return fmt.Errorf("error1")
	}
	return nil
}

func fn2() error {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(10)
	if i%3 == 0 {
		return fmt.Errorf("error2")
	}
	return nil
}

func fn3() error {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(10)
	if i%3 == 0 {
		return fmt.Errorf("error3")
	}
	return nil
}
