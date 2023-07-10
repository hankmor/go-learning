package foobarbaz

import (
	"context"
	"errors"
)

type Foo struct {
	X int
}

type Bar struct {
	X int
}

type Baz struct {
	X int
}

// ProvideFoo 返回 Foo.
func ProvideFoo() Foo {
	return Foo{X: 42}
}

// ProvideBar 返回 Bar: a negative Foo.
func ProvideBar(foo Foo) Bar {
	return Bar{X: -foo.X}
}

// ProvideBaz 返回 Bar 和 error
func ProvideBaz(ctx context.Context, bar Bar) (Baz, error) {
	if bar.X == 0 {
		return Baz{}, errors.New("cannot provide baz when bar is zero")
	}
	return Baz{X: bar.X}, nil
}

// Other 一个新的结构
type Other struct {
}

func NewOther() Other {
	return Other{}
}
