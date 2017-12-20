package ctl

import (
	"github.com/urfave/cli"
)

// cliConnArgs contains CLI flags for API connection info.
var cliConnArgs struct {
	// ConnMultiaddr is the multi-addr to reach the API on.
	ConnMultiaddr string
}

func init() {
	CtlFlags = append(CtlFlags, cli.StringFlag{
		Name:        "api",
		Usage:       "Multiaddr to contact the API.",
		Value:       "/ip4/127.0.0.1/tcp/4050",
		Destination: &cliConnArgs.ConnMultiaddr,
	})
}
