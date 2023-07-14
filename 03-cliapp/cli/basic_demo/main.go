package main

import (
	"cli_demo/cmd"
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
)

var Verbose bool

func main() {
	cliApp := cli.NewApp()
	cliApp.Name = "demo-cli"
	cliApp.Usage = "cli usage demo"
	cliApp.Version = "0.0.1"
	// 系统命令
	cliApp.Commands = cmd.Commands()
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
