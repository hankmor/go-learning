package command

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

func ExitCmd() *cli.Command {
	return &cli.Command{
		Name: "exit",
		Flags: []cli.Flag{
			&cli.IntFlag{Name: "code", Aliases: []string{"c"}, Value: 0, DefaultText: "0", Usage: "custom exit `CODE`"}, // DefaultText 可以用于再帮助文档中显示不同于默认值的信息
		},
		Action: func(ctx *cli.Context) error {
			code := ctx.Int("code") // 退出码
			fmt.Println("code", code)
			if ctx.Bool("i") {
				fmt.Println("this is verbose info")
			}
			return cli.Exit("app exit with code", code) // 自定义退出消息和退出码
		},
	}
}
