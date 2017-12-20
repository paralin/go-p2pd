package main

import (
	"github.com/paralin/go-p2pd/ctl"
	"github.com/urfave/cli"
)

func init() {
	CliCommands = append(CliCommands, cli.Command{
		Name:        "ctl",
		Usage:       "Ctl contains all control commands.",
		Before:      ctl.CtlBefore,
		Flags:       ctl.CtlFlags,
		Subcommands: ctl.CtlCommands,
	})
}
