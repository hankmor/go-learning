package math

import (
	"testing"
)

// 测试函数签名：func TestXxx(t *testing.T)
func TestAdd(t *testing.T) {
    got := Add(1, 2)
    want := 3

    if got != want {
        // t.Errorf 输出错误信息，测试继续执行
        // t.Fatalf 输出错误信息，并立即终止当前测试函数
        t.Errorf("Add(1, 2) = %d; want %d", got, want)
    }
}

func TestAddTableDriven(t *testing.T) {
    // 1. 定义测试用例结构
    tests := []struct {
        name string // 用例名称
        a, b int    // 输入参数
        want int    // 期望结果
    }{
        {"Positive", 1, 2, 3},
        {"Negative", -1, -2, -3},
        {"Mixed", -1, 1, 0},
        {"Zero", 0, 0, 0},
    }

    // 2. 遍历执行
    for _, tt := range tests {
        // t.Run 启动子测试，方便定位具体的错误用例
        t.Run(tt.name, func(t *testing.T) {
            got := Add(tt.a, tt.b)
            if got != tt.want {
                t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
            }
        })
    }
}

// 这是一个基准测试
// 目标是运行函数 b.N 次，看看平均耗时
func BenchmarkAdd(b *testing.B) {
    // b.N 是 Go 框架自动调整的一个足够大的数（例如 1000, 1000000...）
    // 以确保运行时间足够长，从而获得稳定的平均值
    for i := 0; i < b.N; i++ {
        Add(1, 2)
    }
}
