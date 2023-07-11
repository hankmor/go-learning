package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"math/rand"
	"time"
)

var weathers = []string{"sunny", "windy", "cloudy", "rainy"}

//func HelloCmd() *cli.Command {
//	return &cli.Command{
//		Name:    "hello",        // 命令名称，执行时需要指定
//		Aliases: []string{"ho"}, // 命令别名，简化名称
//		Usage:   "向您问好，-h 查看更多帮助信息",
//		Before: func(context *cli.Context) error {
//			fmt.Println("sayHello 命令 Before...")
//			return nil
//		},
//		Flags: []cli.Flag{
//			&cli.StringFlag{Name: "n", Aliases: []string{"name"}, Usage: "您的姓名 `NAME`", Required: true},
//		},
//		Subcommands: cli.Commands{
//			&cli.Command{
//				Name:    "weatherCmd",     // 命令名称，执行时需要指定
//				Aliases: []string{"w"}, // 命令别名，简化名称
//				Usage:   "报告天气情况，-h 查看更多帮助信息",
//				Before: func(context *cli.Context) error {
//					fmt.Println("sayHello weatherCmd 子命令 Before...")
//					return nil
//				},
//				Flags: []cli.Flag{},
//				Action: func(ctx *cli.Context) error {
//					name := ctx.String("n")
//					rd := rand.New(rand.NewSource(time.Now().UnixNano()))
//					weatherCmd := weathers[rd.Intn(len(weathers))]
//					fmt.Printf("hello %s, today is a %s day!", name, weatherCmd)
//					return nil
//				},
//			},
//		}, // 子命令
//		Action: func(ctx *cli.Context) error { // 具体命令的执行逻辑
//			name := ctx.String("n")
//			fmt.Println("hello,", name, "!")
//			return nil
//		},
//	}
//}

var beforeSayHello = func(ctx *cli.Context) error {
	fmt.Println("sayHello 命令 Before...")
	return nil
}

func HelloCmd() *cli.Command {
	return &cli.Command{
		Name:    "hello",        // 命令名称，执行时需要指定
		Aliases: []string{"ho"}, // 命令别名，简化名称
		Usage:   "向您问好，-h 查看更多帮助信息",
		Before:  beforeSayHello,
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "n", Aliases: []string{"name"}, Usage: "您的姓名 `NAME`", Required: true},
		},
		Subcommands: cli.Commands{weatherCmd()}, // 子命令
		Action:      sayHello,                   // 具体命令的执行逻辑
	}
}

func sayHello(ctx *cli.Context) error {
	name := ctx.String("n")
	fmt.Println("hello,", name, "!")
	return nil
}

func weatherCmd() *cli.Command {
	return &cli.Command{
		Name:    "weatherCmd",  // 命令名称，执行时需要指定
		Aliases: []string{"w"}, // 命令别名，简化名称
		Usage:   "报告天气情况，-h 查看更多帮助信息",
		Before: func(context *cli.Context) error {
			fmt.Println("sayHello weatherCmd 子命令 Before...")
			return nil
		},
		Flags:  []cli.Flag{},
		Action: reportWeather,
	}
}

func reportWeather(ctx *cli.Context) error {
	name := ctx.String("n")
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	weather := weathers[rd.Intn(len(weathers))]
	fmt.Printf("hello %s, today is a %s day!", name, weather)
	return nil
}
