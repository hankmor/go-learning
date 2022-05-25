package fuzz

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func Reverse(s string) (string, error) {
	// 非utf8编码的字符，不能翻转，如：\xe4，翻转后乱码
	if !utf8.ValidString(s) {
		return s, errors.New("input is not valid UTF-8")
	}
	// bs := []byte(s) // 只能处理英文字符，中文字符乱码
	bs := []rune(s) // 改为rune类型，可以处理单个字符
	for i, j := 0, len(bs)-1; i < len(bs)/2; i, j = i+1, j-1 {
		bs[i], bs[j] = bs[j], bs[i]
	}
	return string(bs), nil
}

func Run() {
	input := "The quick brown fox jumped over the lazy dog"
	rev, _ := Reverse(input)
	doubleRev, _ := Reverse(rev)
	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q\n", rev)
	fmt.Printf("reversed again: %q\n", doubleRev)

	ch, cnErr := Reverse("中国加油")
	em, emErr := Reverse("笑😁")
	fmt.Printf("中文测试：%q, error: %v\n", ch, cnErr)
	fmt.Printf("笑😁：%q, error: %v\n", em, emErr)

	// 测试一个非utf8字符
	un, unErr := Reverse("\xe4")
	fmt.Printf("非 utf8 字符: %q, error: %v\n", un, unErr)
	// 非 utf8 字符: "\xe4", error: input is not valid UTF-8
}
