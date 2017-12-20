package daemon

import (
	ma "github.com/multiformats/go-multiaddr"
	"github.com/paralin/go-p2pd/node"
	"github.com/pkg/errors"
)

// loadStartDbNodes loads and starts all nodes in the database.
// Typically called at initialization.
func (d *Daemon) loadStartDbNodes() error {
	le := d.log.WithField("proc", "start-db-nodes")

	le.Debug("loading nodes from database")
	nodeList, err := d.daemonDb.ListNodes()
	if err != nil {
		return err
	}

	started := 0
	for _, nod := range nodeList {
		nle := le.
			WithField("id", nod.ID).
			WithField("state", nod.State.String())

		nle.Debug("loading node")
		if nod.State != node.NodeSpecState_STARTED {
			continue
		}

		nle.Debug("starting node")
		if _, err := d.StartNode(nod); err != nil {
			nle.WithError(err).Error("unable to start node")
			nod.State = node.NodeSpecState_STOPPED
			if err := d.daemonDb.SaveNode(nod); err != nil {
				return err
			}
			continue
		}
		started++
		nle.Info("started node")
	}

	if started == 0 && len(nodeList) > 0 {
		return errors.Errorf("couldn't start any of the %d registered nodes", len(nodeList))
	}

	return nil
}

// StartNode attempts to start a node from a spec.
func (d *Daemon) StartNode(spec *node.NodeSpec) (*node.Node, error) {
	if err := spec.Validate(); err != nil {
		return nil, err
	}

	if exist, loaded := d.runningNodes.Load(spec.ID); loaded {
		return exist.(*node.Node), nil
	}

	n, err := node.FromSpec(spec)
	if err != nil {
		return nil, err
	}

	var addrs []ma.Multiaddr
	for ai, addr := range spec.Addrs {
		maddr, err := ma.NewMultiaddr(addr)
		if err != nil {
			return nil, errors.Errorf("addrs[%d]: invalid addr: %v", ai, err.Error())
		}
		if d.addrFilters.AddrBlocked(maddr) {
			d.log.WithField("addr", maddr.String()).Debug("ignoring blocked address")
			continue
		}
		addrs = append(addrs, maddr)
	}

	actNode, loaded := d.runningNodes.LoadOrStore(n.GetId(), n)
	if loaded {
		n.Close()
		return actNode.(*node.Node), nil
	}

	if err := n.StartWithAddrs(d.ctx, addrs); err != nil {
		d.runningNodes.Delete(spec.ID)
		spec.State = node.NodeSpecState_STOPPED
		_ = d.daemonDb.SaveNode(spec)
		return nil, err
	}

	return n, nil
}
