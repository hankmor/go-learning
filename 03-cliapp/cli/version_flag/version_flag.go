package main

import (
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	// 修改默认的版本查询Flag（-v, --version） 为 `--print-version, -V`
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "print-version",
		Aliases: []string{"V"},
		Usage:   "print only the version",
	}

	app := &cli.App{
		Name:    "partay",
		Version: "v19.99.0",
	}
	app.Run(os.Args)
}
