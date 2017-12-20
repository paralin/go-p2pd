package daemon

import (
	"context"

	"github.com/paralin/go-p2pd/control"
	"github.com/paralin/go-p2pd/node"
	"github.com/pkg/errors"
)

// daemonControlServer implements the control.ControlService service.
type daemonControlServer struct {
	*Daemon
}

// newDaemonControlServer builds a new daemonControlServer
func newDaemonControlServer(daemon *Daemon) control.ControlServiceServer {
	return &daemonControlServer{daemon}
}

// CreateNode creates a new node.
func (d *daemonControlServer) CreateNode(
	ctx context.Context,
	req *control.CreateNodeRequest,
) (*control.CreateNodeResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	_, ok := d.runningNodes.Load(req.GetNodeId())
	if ok {
		return nil, errors.Errorf("node already exists: %s", req.GetNodeId())
	}

	nNode, err := node.NewNode(req.GetNodeId())
	if err != nil {
		return nil, err
	}

	err = d.RegisterNode(nNode)
	if err != nil {
		nNode.Close()
		return nil, err
	}

	return &control.CreateNodeResponse{
		NodePeerId: nNode.GetPeerId().Pretty(),
	}, nil
}

// StartNode starts a node.
func (d *daemonControlServer) StartNode(
	ctx context.Context,
	req *control.StartNodeRequest,
) (*control.StartNodeResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	nodeId := req.GetNodeId()
	spec, err := d.daemonDb.GetNode(nodeId)
	if err != nil {
		return nil, err
	}
	if spec == nil {
		return nil, errors.Errorf("node id not known: %s", nodeId)
	}

	nod, err := d.Daemon.StartNode(spec)
	if err != nil {
		return nil, errors.WithMessage(err, "start node")
	}

	resp := &control.StartNodeResponse{NodePeerId: nod.GetPeerId().Pretty()}
	laddrs := nod.GetListenAddrs()
	for _, addr := range laddrs {
		resp.NodeListenAddrs = append(resp.NodeListenAddrs, addr.String())
	}

	return resp, nil
}
