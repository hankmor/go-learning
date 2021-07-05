package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// ==================
// 最基本问候方法
// ==================

/*
Hello 函数，参数name为string类型，返回值为string
*/
func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v, Welcome!", name)
	return message
}

// ==================
// 带异常信息的问候
// ==================

/*
Hello1 函数，参数name为string，返回两个结果：string和一个error信息，如果name为空字符串，返回错误信息
*/
func Hello1(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	return fmt.Sprintf("Hi, %v, Welcome!", name), nil
}

// ==================
// 随机的问候语
// ==================

/*
RandomHello 能够随机返回一句问候语句
*/
func RandomHello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	return fmt.Sprintf(randomFormat(), name), nil
}

// init 初始化方法，程序启动时自动执行
func init() {
	fmt.Println("初始化随机数...")
	// 设置随机数种子，每次随机数生成不同
	rand.Seed(time.Now().UnixNano())
}

// randomFormat 随机返回问候语句，使用slice类型存储问候语句，并使用rand来制定随机返回
// 小写字母开头的方法，只能包内可见，外部不能访问
func randomFormat() string {
	// slice类型，类似数组，但是并不指定长度，其长度可以自动变化
	formats := []string{
		"Hi, %v, welcome!",
		"Greate to see you, %v",
		"Hail, %v! Well met!",
	}

	// rand生成随机数
	return formats[rand.Intn(len(formats))]
}

// ==================
// 问候多个人
// ==================

// Hellos 方法传入数组类型的多个名称，返回存储了问候信息的map类型
func Hellos(names []string) (map[string]string, error) {
	// 创建map
	messages := make(map[string]string)
	// 遍历名称数组，两个变量，第一个为数组下边，第一个为数组元素，这里用不到下标，所以使用"_"符号忽略下标值
	for _, name := range names {
		// 调用RandomHello方法随机生成问候语
		message, err := RandomHello(name)
		// 有异常，直接返回nil和错误信息
		if err != nil {
			return nil, err
		}
		// 给map赋值
		messages[name] = message
	}
	return messages, nil
}
