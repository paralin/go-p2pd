package ctl

import (
	"time"

	"github.com/pkg/errors"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
)

// CtlCommands contains top level commands.
// note: see cli_conn.go
var CtlCommands []cli.Command

// CtlFlags contains top level flags.
// note: see cli_conn.go
var CtlFlags []cli.Flag

// cliApiConn is the cli GRPC connection to the api.
var cliApiConn *grpc.ClientConn

// cliDialTimeout is the time to wait for a connection in the CLI.
var cliDialTimeout = time.Second * 10

// CtlBefore contains the before func to run.
var CtlBefore cli.BeforeFunc = func(c *cli.Context) (err error) {
	// build cli connection
	cliApiConn, err = grpc.Dial(
		cliConnArgs.ConnMultiaddr,
		grpc.WithDialer(MultiaddrDialer),
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithTimeout(cliDialTimeout),
	)
	if err != nil {
		err = errors.WithMessage(err, "dial api server at "+cliConnArgs.ConnMultiaddr)
	}
	return
}
