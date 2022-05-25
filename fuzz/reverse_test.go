package fuzz

import (
	"testing"
	"unicode/utf8"
)

/*
测试命令：go test
测试并输出详细信息：go test -v
如果只想运行某一个测试方法，执行：go test -v -run="FuzzReverse"，-run后对应的是一个正则表达式

运行任何随机生成的字符串输入，进行模糊测试，执行：go test -v -fuzz=Fuzz
如果测试时发生故障，导致问题的输入会被写入语料库文件中，下次运行就可以使用
错误信息：
	--- FAIL: FuzzReverse (0.00s)
		--- FAIL: FuzzReverse/d97214ce235bfcf5f4cc06763db1c2e6f45fb2ca1cb41d4f19dda599e5798692 (0.00s)
	        reverse_test.go:57: Before: "\xe4", after: "�"
运行失败的测试用力：go test -v -run=FuzzReverse/d97214ce235bfcf5f4cc06763db1c2e6f45fb2ca1cb41d4f19dda599e5798692
由于 "\xe4" 不是utf8字符，所以 Reverse 方法不能处理，增加判断即可修复这个 bug

默认情况下，go test -fuzz=Fuzz 会一直执行，CTRL + C停止，可以通过 -fuzztime 30s 指定时间：go test -fuzz=Fuzz -fuzztime 30s
*/

// 单元测试，测试方法以 Test 开头
func TestReverse(t *testing.T) {
	type Case struct {
		in, want string
	}

	// 准备测试用例
	cases := []Case{
		// 顺序赋值
		{"Hello World!", "!dlroW olleH"},
		{" ", " "},
		{"", ""},
		{"123456", "654321"},
		// 中文测试
		{"你好，中国", "国中，好你"},
		// Emoji测试
		{"你好，😄", "😄，好你"},
	}

	for _, c := range cases {
		r, _ := Reverse(c.in)
		if r != c.want {
			t.Errorf("Reverse: %q, want: %q", r, c.want)
		}
	}
}

// 模糊测试：可以为您的代码提供输入，并且可以识别您提出的测试用例没有达到的边缘用例。
// 模糊测试以 Fuzz 开头，无法控制输入，但是可以通过一些方式验证输入与输出的正确性，比如这里的翻转两次结果与输入相同
func FuzzReverse(f *testing.F) {
	// 测试用力
	testcases := []string{"Hello, world", " ", "!12345", "你好", "哈，😁"}
	for _, tc := range testcases {
		f.Add(tc) // 添加种子语料库
	}
	// 执行测试
	f.Fuzz(func(t *testing.T, orig string) {
		rev, rErr := Reverse(orig)
		if rErr != nil {
			return // 出错则跳过测试
		}
		doubleRev, drErr := Reverse(rev)
		if drErr != nil {
			t.Skip() // 跳过
		}
		t.Logf("Number of runes: orig=%d, rev=%d, doubleRev=%d", utf8.RuneCountInString(orig), utf8.RuneCountInString(rev), utf8.RuneCountInString(doubleRev))
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}
