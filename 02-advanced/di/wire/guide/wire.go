//go:build wireinject
// +build wireinject

package main

import (
	"context"
	"github.com/google/wire"
	"wire_demo/guide/foobarbaz"
)

// SuperSet 使用NewSet创建集合
var SuperSet = wire.NewSet(foobarbaz.ProvideFoo, foobarbaz.ProvideBar, foobarbaz.ProvideBaz)

var OtherSet = wire.NewSet(foobarbaz.NewOther)

// MegaSet 可以将其他集合加入新的集合中
var MegaSet = wire.NewSet(SuperSet, OtherSet)

func initializeBaz(ctx context.Context) (foobarbaz.Baz, func(), error) {
	panic(wire.Build(MegaSet))
	return foobarbaz.Baz{}, func() {
	}, nil
}
