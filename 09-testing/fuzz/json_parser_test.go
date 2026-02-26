package main

import (
	"strings"
	"testing"
)

// TestParseJSON 单元测试
func TestParseJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{"简单对象", `{"a":1}`, false},
		{"简单数组", `[1,2,3]`, false},
		{"嵌套对象", `{"a":{"b":{"c":1}}}`, false},
		{"空对象", `{}`, false},
		{"空数组", `[]`, false},
		{"字符串值", `{"name":"老墨"}`, false},
		{"布尔值", `{"active":true}`, false},
		{"null值", `{"value":null}`, false},
		{"无效JSON", `{invalid}`, true},
		{"未闭合", `{"a":1`, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ParseJSON(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestParseJSONWithMaxDepth 测试最大深度限制
func TestParseJSONWithMaxDepth(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		maxDepth int
		wantErr  bool
	}{
		{"深度2-限制3", `{"a":{"b":1}}`, 3, false},
		{"深度3-限制3", `{"a":{"b":{"c":1}}}`, 3, false},
		{"深度4-限制3", `{"a":{"b":{"c":{"d":1}}}}`, 3, true},
		{"数组嵌套", `[[[1]]]`, 3, false},
		{"混合嵌套", `{"a":[{"b":1}]}`, 3, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ParseJSON(tt.input, MaxDepth(tt.maxDepth))
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// FuzzParseJSON 模糊测试 - 发现意外的 panic 和边界情况
func FuzzParseJSON(f *testing.F) {
	// 添加种子语料库
	f.Add(`{"a":1}`)
	f.Add(`[1,2,3]`)
	f.Add(`{"name":"老墨"}`)
	f.Add(`{"nested":{"value":true}}`)
	f.Add(`[]`)
	f.Add(`{}`)
	f.Add(`null`)
	f.Add(`"string"`)
	f.Add(`123`)
	f.Add(`true`)

	f.Fuzz(func(t *testing.T, input string) {
		// 模糊测试的目标：确保不会 panic
		// 设置最大嵌套深度，防止 OOM
		_, err := ParseJSON(input, MaxDepth(100))

		// 不关心是否返回错误，只要不 panic 就行
		// 如果输入是有效 JSON，应该能解析
		// 如果输入是无效 JSON，应该返回错误而不是 panic
		_ = err
	})
}

// FuzzValidateJSON 模糊测试 - 验证 JSON 合法性
func FuzzValidateJSON(f *testing.F) {
	// 添加种子语料库
	f.Add(`{"a":1}`)
	f.Add(`[1,2,3]`)
	f.Add(`invalid`)
	f.Add(`{"unclosed":`)

	f.Fuzz(func(t *testing.T, input string) {
		// 确保验证函数不会 panic
		err := ValidateJSON(input)
		_ = err
	})
}

// FuzzCalculateDepth 模糊测试 - 测试深度计算
func FuzzCalculateDepth(f *testing.F) {
	f.Add(`{}`)
	f.Add(`{"a":{"b":1}}`)
	f.Add(`[[[1]]]`)

	f.Fuzz(func(t *testing.T, input string) {
		depth := calculateDepth(input)

		// 深度不应该是负数
		if depth < 0 {
			t.Errorf("calculateDepth() returned negative depth: %d", depth)
		}

		// 深度不应该超过输入长度
		if depth > len(input) {
			t.Errorf("calculateDepth() returned depth %d > input length %d", depth, len(input))
		}
	})
}

// BenchmarkParseJSON 基准测试
func BenchmarkParseJSON(b *testing.B) {
	input := `{"name":"老墨","age":30,"skills":["Go","Python","JavaScript"]}`

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ParseJSON(input)
	}
}

// BenchmarkParseJSONDeep 基准测试 - 深层嵌套
func BenchmarkParseJSONDeep(b *testing.B) {
	// 生成深度为 50 的嵌套 JSON
	var sb strings.Builder
	for i := 0; i < 50; i++ {
		sb.WriteString(`{"a":`)
	}
	sb.WriteString(`1`)
	for i := 0; i < 50; i++ {
		sb.WriteString(`}`)
	}
	input := sb.String()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ParseJSON(input, MaxDepth(100))
	}
}
