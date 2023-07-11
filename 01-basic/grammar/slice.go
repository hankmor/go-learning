package main

import (
	"bytes"
	"fmt"
)

type path []byte

func (p *path) TruncateAtFinalSlash() {
	fmt.Printf("%x\n", p)
	i := bytes.LastIndex(*p, []byte("/"))
	if i >= 0 {
		*p = (*p)[0:i]
	}
}

func (p path) TruncateAtFinalSlash1() {
	fmt.Printf("%x\n", p)
	i := bytes.LastIndex(p, []byte("/"))
	if i >= 0 {
		p = (p)[0:i] // 调用者复制了一份 p，将 p 重新切片再传给 p 其实无法改变原始切片
	}
}

func main() {
	pathName := path("/usr/bin/tso") // Conversion from string to path.
	pathName.TruncateAtFinalSlash()  // 正常工作
	fmt.Printf("%s\n", pathName)

	pathName1 := path("/usr/bin/tso")
	pathName1.TruncateAtFinalSlash1() // 无法正常工作
	fmt.Printf("%s\n", pathName1)
}
