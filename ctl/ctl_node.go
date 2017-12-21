package ctl

import (
	"github.com/urfave/cli"
)

var CtlNodeCommands []cli.Command

func init() {
	initNodeAdd()
	initNodeStart()
	initNodeListen()
	initNodeStatus()
	CtlCommands = append(CtlCommands, cli.Command{
		Name:        "node",
		Usage:       "Contains all node-related subcommands.",
		Subcommands: CtlNodeCommands,
	})
}
