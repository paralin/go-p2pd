package ctl

import (
	"context"
	"fmt"

	"github.com/paralin/go-p2pd/control"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

func initNodeListen() {
	CtlNodeCommands = append(CtlNodeCommands, cli.Command{
		Name:   "listen",
		Usage:  "Listen commands a started node to listen on an additional address.",
		Action: ctlNodeListen,
	})
}

// ctlNodeListen commands a node to listen to an addr.
func ctlNodeListen(c *cli.Context) error {
	ctx, ctxCancel := context.WithCancel(context.Background())
	defer ctxCancel()

	args := c.Args()
	if len(args) != 2 {
		return errors.New("expected 2 arguments, example: node_id /ipv4/0.0.0/tcp/4001")
	}

	nodeId := args[0]
	multiAddrStr := args[1]

	req := &control.ListenNodeRequest{
		NodeId: nodeId,
		Addr:   multiAddrStr,
	}

	// do some early validation
	if err := req.Validate(); err != nil {
		return errors.Errorf("args invalid: %v", err.Error())
	}

	client := control.NewControlServiceClient(cliApiConn)
	resp, err := client.ListenNode(ctx, req)
	if err != nil {
		return err
	}

	fmt.Printf("Node listening\n\t/p2p/%s\n", resp.GetNodePeerId())
	if len(resp.GetNodeListenAddrs()) > 0 {
		fmt.Printf(control.NodeListenResponseString(control.NodeListenResponse(resp)))
	}

	return nil
}
