package ctl

import (
	"context"
	"fmt"

	"github.com/paralin/go-p2pd/control"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

func initNodeAdd() {
	CtlNodeCommands = append(CtlNodeCommands, cli.Command{
		Name:   "add",
		Usage:  "Adds a new node to the p2pd instance.",
		Action: ctlNodeAdd,
	})
}

// ctlNodeAdd adds a node to the daemon.
func ctlNodeAdd(c *cli.Context) error {
	ctx, ctxCancel := context.WithCancel(context.Background())
	defer ctxCancel()

	args := c.Args()
	if len(args) != 1 {
		return errors.New("expected 1 argument with id for new node")
	}

	nodeId := args[0]
	client := control.NewControlServiceClient(cliApiConn)
	resp, err := client.CreateNode(ctx, &control.CreateNodeRequest{NodeId: nodeId})
	if err != nil {
		return err
	}

	fmt.Printf("Created node\n\t/p2p/%s\n\n", resp.GetNodePeerId())
	return nil
}
