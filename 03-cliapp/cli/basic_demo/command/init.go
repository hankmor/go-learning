package command

import "github.com/urfave/cli/v2"

var cmds cli.Commands

func Commands() cli.Commands {
	cmds = append(cmds, HelloCmd(), ExitCmd())
	return cmds
}
