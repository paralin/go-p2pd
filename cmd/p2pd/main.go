package main

import (
	"github.com/paralin/go-p2pd/meta"
	"github.com/urfave/cli"
)

// CliCommands contains all the root-level commands.
var CliCommands []cli.Command

// CliFlags contains all the root-level flags.
var CliFlags []cli.Flag

func main() {
	app := cli.NewApp()
	app.Name = "p2pd"
	app.Usage = "p2pd daemon and cli."
	app.Authors = meta.Authors
	app.Commands = CliCommands
	app.Flags = CliFlags
	app.HideVersion = true
	app.RunAndExitOnError()
}
