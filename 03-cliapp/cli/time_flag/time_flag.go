package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		// 指定时间格式的Flag，Layout为格式，timezone为时区（默认UTC）
		Flags: []cli.Flag{
			&cli.TimestampFlag{Name: "meeting", Layout: "2006-01-02 15:04:05", Timezone: time.Local}, // 使用本地时区
		},
		Action: func(ctx *cli.Context) error {
			fmt.Printf("%s", ctx.Timestamp("meeting").String()) // 用 Timestamp
			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
	// 运行时参数 -meeting "2023-07-14 00:00:10"，时间中间有空哥，需要加上引号
	// 输出： 2023-07-14 00:00:10 +0800 CST
}
