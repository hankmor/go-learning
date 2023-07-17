package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"math/rand"
	"os"
	"time"
)

var Verbose bool
var weathers = []string{"sunny", "windy", "cloudy", "rainy"}

func main() {
	cliApp := cli.NewApp()
	cliApp.Name = "demo-cli"
	cliApp.Usage = "cli usage demo"
	cliApp.Version = "0.0.1"
	// 系统命令
	cliApp.Commands = []*cli.Command{sayHelloCmd()}
	//cliApp.Commands = command.Commands()
	// 初始化之前调用
	cliApp.Before = func(ctx *cli.Context) error {
		fmt.Println("Before app run ...")
		return nil
	}
	// 全局参数
	cliApp.Flags = append(cliApp.Flags, []cli.Flag{
		&cli.BoolFlag{Name: "i", Usage: "show verbose info", Required: false, Destination: &Verbose}, // destination 可以将设置的参数绑定到变量，后续可以直接使用
	}...)

	err := cliApp.Run(os.Args) // app退出不会调用 os.Exit，所以默认退出代码都是0，可以通过 cli.Exit方法指定退出信息和退出码
	if err != nil {
		fmt.Printf("demo-cli execute error: %v\n", err)
		os.Exit(-1)
	}
}

func sayHelloCmd() *cli.Command {
	return &cli.Command{
		Name:    "hello",        // 命令名称，执行时需要指定
		Aliases: []string{"ho"}, // 命令别名，简化名称
		Usage:   "向您问好，-h 查看更多帮助信息",
		Before: func(context *cli.Context) error {
			fmt.Println("sayHello 命令 Before...")
			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "n", Aliases: []string{"name"}, Usage: "您的姓名 `NAME`", Required: true},
		},
		Subcommands: cli.Commands{
			&cli.Command{
				Name:    "weather",     // 命令名称，执行时需要指定
				Aliases: []string{"w"}, // 命令别名，简化名称
				Usage:   "报告天气情况，-h 查看更多帮助信息",
				Before: func(context *cli.Context) error {
					fmt.Println("sayHello weatherCmd 子命令 Before...")
					return nil
				},
				Flags: []cli.Flag{
					//&cli.StringFlag{Name: "n", Aliases: []string{"name"}, Usage: "您的姓名 `NAME`", Required: true},
				},
				Action: func(ctx *cli.Context) error {
					name := ctx.String("n")
					rd := rand.New(rand.NewSource(time.Now().UnixNano()))
					weatherCmd := weathers[rd.Intn(len(weathers))]
					fmt.Printf("hello %s, today is a %s day!\n", name, weatherCmd)
					return nil
				},
				Category: "weather",
			},
			&cli.Command{
				Name:    "complain-weather", // 命令名称，执行时需要指定
				Aliases: []string{"cw"},     // 命令别名，简化名称
				Usage:   "Complains the weather today",
				Before: func(ctx *cli.Context) error {
					return nil
				},
				Flags: []cli.Flag{},
				Action: func(ctx *cli.Context) error {
					return nil
				},
				Category: "weather",
			},
		}, // 子命令
		Action: func(ctx *cli.Context) error { // 具体命令的执行逻辑
			name := ctx.String("n")
			fmt.Println("hello,", name, "!")
			return nil
		},
	}
}
