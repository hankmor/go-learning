package main

import (
	"fmt"
	"myrpc"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		panic(err)
	}
	var ret int
	// 调用rpc方法，方法名称为 类型.方法名
	err = client.Call("Arith.Multiply", &myrpc.Param{10, 2}, &ret) // 同步调用远程方法
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arith.Multiply: %d\n", ret)

	divCall := client.Go("Arith.Divide", &myrpc.Param{10, 3}, new(myrpc.Quotient), nil) // 异步调用远程方法
	reply := <-divCall.Done
	q := reply.Reply.(*myrpc.Quotient)
	fmt.Printf("Arith.Divide: qua = %d, rem = %d\n", q.Quo, q.Rem)

	/*
		Arith.Multiply: 20
		Arith.Divide: qua = 3, rem = 1
	*/
}
