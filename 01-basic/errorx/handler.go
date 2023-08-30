package main

import (
	"fmt"
	"math/rand"
	"reflect"
)

type ErrHandler struct {
	err error
}

type Processor func() error
type TwoTupleProcessor func() (any, error)
type ThreeTupleProcessor func() (any, any, error)
type FourTupleProcessor func() (any, any, any, error)
type FiveTupleProcessor func() (any, any, any, any, error)

func NewHandler() *ErrHandler {
	return &ErrHandler{}
}

func (h *ErrHandler) Handle(p Processor) {
	if h.err == nil {
		err := p()
		h.err = err
	}
}

// 返回的结果还需要断言，不如不返回结果

func (h *ErrHandler) Handle2(p TwoTupleProcessor) any {
	rs := getFuncReturns(p)
	if h.err == nil {
		r, err := p()
		h.err = err
		return r
	} else {
		return reflect.Zero(rs[0]).Interface()
	}
}

func getFuncReturns(f any) []reflect.Type {
	ft := reflect.TypeOf(f)
	if ft.Kind() != reflect.Func {
		return []reflect.Type{}
	}
	var ts = make([]reflect.Type, ft.NumOut())
	for i := 0; i < ft.NumOut(); i++ {
		ts[i] = ft.Out(i)
	}
	return ts
}

// main

type user struct {
	name string
}

func main() {
	var err error
	var s string
	h := NewHandler()
	h.Handle(func() error {
		return f1(0, "haha")
	})
	fmt.Println(h.err)
	h.Handle(func() error {
		s, err = f2(0, 1, "a", "b")
		return err
	})
	fmt.Println(s, h.err)
	v := h.Handle2(func() (any, error) {
		return f2(0, 1, "a", "b")
	})
	fmt.Println(v, h.err)
	if v1 := h.Handle2(func() (any, error) {
		return f3("a")
	}); v1 != nil {
		fmt.Println(v1.(*user))
	}
	fmt.Println(h.err)
	//var x any
	//u := x.(*user) // panic, nil
	//fmt.Println(u)
	u := reflect.Zero(reflect.TypeOf(&user{})).Interface()
	fmt.Println(u.(*user))

	if h.err != nil {
		fmt.Println("has error")
	} else {
		fmt.Println("ok")
	}
}

func f1(p1 int, p2 string) error {
	if rand.Intn(100)%3 == 0 {
		return fmt.Errorf("test error")
	}
	return nil
}

func f2(p0, p1 int, p2, p3 string) (string, error) {
	if rand.Intn(100)%3 == 0 {
		return "test", fmt.Errorf("test error")
	}
	return "demo", nil
}

func f3(name string) (*user, error) {
	if rand.Intn(100)%3 == 0 {
		return nil, fmt.Errorf("test error")
	}
	return &user{name: "hank"}, nil
}
