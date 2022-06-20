package main

import (
	"errors"
	"myrpc"
	"net"
	"net/http"
	"net/rpc"
)

/*
RPC 必须符合方法签名格式：
func (t *T) MethodName(argType T1, replyType *T2) error
*/

type Arith int

// Multiply ret必须为指针，否则会报方法不能找到的错误
func (a *Arith) Multiply(p *myrpc.Param, ret *int) error {
	*ret = p.A * p.B
	return nil
}

func (a *Arith) Divide(p *myrpc.Param, ret *myrpc.Quotient) error {
	if p.B == 0 {
		return errors.New("divide by zero")
	}
	ret.Quo = p.A / p.B
	ret.Rem = p.A % p.B
	return nil
}

func main() {
	// 注册rpc
	ar := new(Arith)
	err := rpc.Register(ar)
	if err != nil {
		panic(err)
	}
	rpc.HandleHTTP() // 注册http处理
	listener, err1 := net.Listen("tcp", ":1234")
	if err1 != nil {
		panic(err1)
	}
	err = http.Serve(listener, nil)
	if err != nil {
		panic(err)
	}
}
