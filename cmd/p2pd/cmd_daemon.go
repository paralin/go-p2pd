package main

import (
	"github.com/paralin/go-p2pd/daemon"
	"github.com/urfave/cli"
)

var cliDaemonConfig *daemon.Config

func init() {
	cliDaemonConfig = daemon.DefaultConfig()
	CliCommands = append(CliCommands, cli.Command{
		Name:  "daemon",
		Usage: "starts the p2pd daemon",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "listen-api",
				Usage:       "API listen multiaddr.",
				Value:       cliDaemonConfig.ApiListen,
				Destination: &cliDaemonConfig.ApiListen,
			},
			cli.StringFlag{
				Name:        "data-path",
				Usage:       "Path (directory) to store daemon database.",
				Value:       cliDaemonConfig.DataPath,
				Destination: &cliDaemonConfig.DataPath,
			},
		},
		Action: func(c *cli.Context) error {
			d, err := daemon.NewDaemon(*cliDaemonConfig)
			if err != nil {
				return err
			}

			return d.Run()
		},
	})
}
