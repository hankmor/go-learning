package toplvl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
go test -v --run="TestPower" 执行某一个具体的测试方法
go test -v . 顺序执行当前目录的所有测试
go test -v . -shuffle=on 打乱顺序，可以检测依赖关系
*/

func TestPower(t *testing.T) {
	var x, n, want uint = 10, 2, 100
	r := Power(x, n)
	if r != want {
		t.Errorf("test failed, got %d, want %d", r, want)
	}
}

func TestPower2(t *testing.T) {
	var x, n, want uint = 100, 3, 100_00_00
	r := Power(x, n)
	assert.Equal(t, r, want)
}
