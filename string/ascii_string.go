package main

import "fmt"

func main() {
	var s = "我爱你，中国！" // 原字符串
	var b = []byte(s) // byte 切片
	var r = []rune(s) // rune 切片

	for i, v := range r {
		var utf8bs []byte // 每一个中文字符对应的utf8字节
		for j := i * 3; j < (i+1)*3; j++ {
			utf8bs = append(utf8bs, b[j])
		}
		fmt.Printf("%s -> %X -> %X\n", string(v), v, utf8bs)
	}

	/*output:
	我 -> 6211 -> E68891
	爱 -> 7231 -> E788B1
	你 -> 4F60 -> E4BDA0
	， -> FF0C -> EFBC8C
	中 -> 4E2D -> E4B8AD
	国 -> 56FD -> E59BBD
	！ -> FF01 -> EFBC81
	*/
}
