package greetings

import (
	"regexp"
	"testing"
)

/*
TestHello 对greetings.Hello方法进行单元测试
*/
func TestHello(t *testing.T) {
	name := "Sam"
	// 调用方法
	msg := Hello(name)
	// 设定测试条件：正则匹配名称
	wantRegex := regexp.MustCompile(`\b` + name + `\b`)
	// 结果不符合预期
	if !wantRegex.MatchString(msg) {
		t.Fatalf("Hello("+name+") = %q, want match for %#q", msg, wantRegex)
	}
}

// 测试Hello1方法
func TestHello1(t *testing.T) {
	name := "Sam"
	msg, err := Hello1(name)
	wantRegex := regexp.MustCompile(`\b` + name + `\b`)
	if !wantRegex.MatchString(msg) || err != nil {
		t.Fatalf("Hello1("+name+") = %q, %v, want match for %#q", msg, err, wantRegex)
	}
}

// 测试异常情况
func TestHello1Empty(t *testing.T) {
	// 传空的名称
	name := ""
	msg, err := Hello1(name)
	// msg为空字符串或者err为nil，说明没有错误，则测试不通过
	if msg != "" || err == nil {
		t.Fatalf("Hello1("+name+") = %q, %v, want '' and error", msg, err)
	}
}
