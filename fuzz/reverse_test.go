package main

import "testing"

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
	}

	for _, c := range cases {
		r := Reverse(c.in)
		if r != c.want {
			t.Errorf("Reverse: %q, want: %q", r, c.want)
		}
	}
}
