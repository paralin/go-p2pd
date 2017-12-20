package main

import (
	"github.com/paralin/go-p2pd/ctl"
	"github.com/paralin/go-p2pd/meta"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "p2pdctl"
	app.Usage = "p2pd control cli"
	app.Authors = meta.Authors
	app.Commands = ctl.CtlCommands
	app.Before = ctl.CtlBefore
	app.Flags = ctl.CtlFlags
	app.HideVersion = true
	app.RunAndExitOnError()
}
