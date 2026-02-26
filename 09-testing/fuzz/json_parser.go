package main

import (
	"encoding/json"
	"fmt"
)

// ParseOption 解析选项
type ParseOption func(*parseConfig)

type parseConfig struct {
	maxDepth int
}

// MaxDepth 设置最大嵌套深度
func MaxDepth(depth int) ParseOption {
	return func(c *parseConfig) {
		c.maxDepth = depth
	}
}

// ParseJSON 解析 JSON 字符串，支持设置最大嵌套深度
func ParseJSON(input string, opts ...ParseOption) (interface{}, error) {
	cfg := &parseConfig{
		maxDepth: 100, // 默认最大深度
	}
	for _, opt := range opts {
		opt(cfg)
	}

	// 检查嵌套深度
	depth := calculateDepth(input)
	if depth > cfg.maxDepth {
		return nil, fmt.Errorf("JSON 嵌套深度 %d 超过最大限制 %d", depth, cfg.maxDepth)
	}

	var result interface{}
	err := json.Unmarshal([]byte(input), &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// calculateDepth 计算 JSON 字符串的嵌套深度
func calculateDepth(input string) int {
	maxDepth := 0
	currentDepth := 0

	for _, ch := range input {
		switch ch {
		case '{', '[':
			currentDepth++
			if currentDepth > maxDepth {
				maxDepth = currentDepth
			}
		case '}', ']':
			currentDepth--
		}
	}

	return maxDepth
}

// ValidateJSON 验证 JSON 字符串是否合法
func ValidateJSON(input string) error {
	var js interface{}
	return json.Unmarshal([]byte(input), &js)
}
