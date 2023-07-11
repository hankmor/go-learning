package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	cs := []Conn{{Id: 1}, {Id: 2}, {Id: 3}}
	fmt.Println(Query(cs, "a"))
	//Query(cs, "b")
	//Query(cs, "c")
}

type Conn struct {
	Id int
}

func (c *Conn) DoQuery(query string) Result {
	time.Sleep(2 * time.Second)
	return Result{Content: strconv.Itoa(c.Id) + " - " + query}
}

type Result struct {
	Content string
}

func Query(conns []Conn, query string) Result {
	ch := make(chan Result)
	for _, conn := range conns {
		go func(c Conn) {
			select {
			case ch <- c.DoQuery(query): // 无法立即发送结果时，走 default 分支，可以实现非阻塞
			default: // 不阻塞
				fmt.Println("default")
			}
		}(conn)
	}
	return <-ch // 阻塞
}
