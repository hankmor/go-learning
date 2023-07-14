package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	// EXAMPLE: 现有帮助模板下添加新信息
	cli.AppHelpTemplate = fmt.Sprintf(`%s

WEBSITE: http://awesometown.example.com

SUPPORT: support@awesometown.example.com

`, cli.AppHelpTemplate)

	// EXAMPLE: 替换默认帮助模板
	cli.AppHelpTemplate = `NAME:
	  {{.Name}} - {{.Usage}}
	USAGE:
	  {{.HelpName}} {{if .VisibleFlags}}[global options]{{end}}{{if .Commands}} command [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}
	  {{if len .Authors}}
	AUTHOR:
	  {{range .Authors}}{{ . }}{{end}}
	  {{end}}{{if .Commands}}
	COMMANDS:
	{{range .Commands}}{{if not .HideHelp}}   {{join .Names ", "}}{{ "\t"}}{{.Usage}}{{ "\n" }}{{end}}{{end}}{{end}}{{if .VisibleFlags}}
	GLOBAL OPTIONS:
	  {{range .VisibleFlags}}{{.}}
	  {{end}}{{end}}{{if .Copyright }}
	COPYRIGHT:
	  {{.Copyright}}
	  {{end}}{{if .Version}}
	VERSION:
	  {{.Version}}
	  {{end}}
	`

	// EXAMPLE: 替换默认 HelpPrinter
	//cli.HelpPrinter = func(w io.Writer, templ string, data interface{}) {
	//	fmt.Println("Ha HA.  I pwnd the help!!1")
	//}

	// 修改默认的 -h,--help 帮助 Flag 为 --haaaaalp, --halp
	cli.HelpFlag = &cli.BoolFlag{
		Name:    "haaaaalp",
		Aliases: []string{"halp"},
		Usage:   "HALP",
		EnvVars: []string{"SHOW_HALP", "HALPPLZ"},
	}

	(&cli.App{}).Run(os.Args)
}
