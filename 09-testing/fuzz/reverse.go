package main

import (
	"errors"
	"unicode/utf8"
)

func Reverse(s string) (string, error) {
	// 非utf8编码的字符，不能翻转，如：\xe4，翻转后乱码
	if !utf8.ValidString(s) {
		return s, errors.New("input is not valid UTF-8")
	}
	// bs := []byte(s) // 只能处理英文字符，中文字符乱码，问题所在
	bs := []rune(s) // 将字符串转为 rune 切片，它可以正确解码单个UTF-8字符
	for i, j := 0, len(bs)-1; i < len(bs)/2; i, j = i+1, j-1 {
		bs[i], bs[j] = bs[j], bs[i]
	}
	return string(bs), nil
}
