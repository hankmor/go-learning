package greetings

import (
	"errors"
	"fmt"
)

/*
Hello 函数，参数name为string类型，返回值为string
*/
func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v, Welcome!", name)
	return message
}

/*
Hello1 函数，参数name为string，返回两个结果：string和一个error信息，如果name为空字符串，返回错误信息
*/
func Hello1(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	return fmt.Sprintf("Hi, %v, Welcome!", name), nil
}
