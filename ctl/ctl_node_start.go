package ctl

import (
	"context"
	"fmt"

	"github.com/paralin/go-p2pd/control"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

func initNodeStart() {
	CtlNodeCommands = append(CtlNodeCommands, cli.Command{
		Name:   "start",
		Usage:  "Starts a previously created node.",
		Action: ctlNodeStart,
	})
}

// ctlNodeStart starts a node.
func ctlNodeStart(c *cli.Context) error {
	ctx, ctxCancel := context.WithCancel(context.Background())
	defer ctxCancel()

	args := c.Args()
	if len(args) != 1 {
		return errors.New("expected 1 argument with id of node to start")
	}

	nodeId := args[0]
	client := control.NewControlServiceClient(cliApiConn)
	resp, err := client.StartNode(ctx, &control.StartNodeRequest{NodeId: nodeId})
	if err != nil {
		return err
	}

	fmt.Printf("Started node\n\t/p2p/%s\n", resp.GetNodePeerId())

	if len(resp.GetNodeListenAddrs()) > 0 {
		fmt.Printf("Swarm listening on\n")
		for _, addr := range resp.GetNodeListenAddrs() {
			fmt.Printf("\t%s\n", addr)
		}
	}

	return nil
}
