package ctl

import (
	"context"
	"fmt"

	"github.com/paralin/go-p2pd/control"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

func initNodeStatus() {
	CtlNodeCommands = append(CtlNodeCommands, cli.Command{
		Name:   "status",
		Usage:  "Status checks node's status.",
		Action: ctlNodeStatus,
	})
}

// ctlNodeStatus commands a node to report its status.
func ctlNodeStatus(c *cli.Context) error {
	ctx, ctxCancel := context.WithCancel(context.Background())
	defer ctxCancel()

	args := c.Args()
	if len(args) != 1 {
		return errors.New("expected 1 argument with node id")
	}

	nodeId := args[0]
	req := &control.StatusNodeRequest{
		NodeId: nodeId,
	}

	// do some early validation
	if err := req.Validate(); err != nil {
		return errors.Errorf("args invalid: %v", err.Error())
	}

	client := control.NewControlServiceClient(cliApiConn)
	resp, err := client.StatusNode(ctx, req)
	if err != nil {
		return err
	}

	fmt.Printf("Node status\n\t/p2p/%s\n", resp.GetNodePeerId())
	if len(resp.GetNodeListenAddrs()) > 0 {
		fmt.Printf(control.NodeListenResponseString(control.NodeListenResponse(resp)))
	}

	return nil
}
