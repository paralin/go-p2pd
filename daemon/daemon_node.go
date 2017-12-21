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
	specStopped := 0
	for _, nod := range nodeList {
		nle := nod.LogFields(le)

		nle.WithField("state", nod.State.String()).Debug("loading node")
		if nod.State != node.NodeSpecState_NODE_SPEC_STATE_STARTED {
			specStopped++
			continue
		}

		nle.Debug("starting node")
		if _, err := d.StartNode(nod); err != nil {
			nle.WithError(err).Error("unable to start node")
			nod.State = node.NodeSpecState_NODE_SPEC_STATE_STOPPED
			if err := d.daemonDb.SaveNode(nod); err != nil {
				return err
			}
			continue
		}
		started++
		nle = nod.LogFields(le)
		nle.Info("started node")
	}

	if started == 0 && (len(nodeList)-specStopped) > 0 {
		return errors.Errorf("couldn't start any of the %d registered nodes", len(nodeList)-specStopped)
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

	n, err := node.FromSpec(spec, d.log)
	if err != nil {
		return nil, err
	}

	var addrs []ma.Multiaddr
	var addrsStrs []string
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
		addrsStrs = append(addrsStrs, maddr.String())
	}

	actNode, loaded := d.runningNodes.LoadOrStore(n.GetId(), n)
	if loaded {
		n.Close()
		return actNode.(*node.Node), nil
	}

	if err := n.StartWithAddrs(d.ctx, addrs); err != nil {
		d.runningNodes.Delete(spec.ID)
		spec.State = node.NodeSpecState_NODE_SPEC_STATE_STOPPED
		_ = d.daemonDb.SaveNode(spec)
		d.log.WithError(err).WithField("addrs", addrs).Error("unable to start node")
		return nil, err
	}

	spec.State = node.NodeSpecState_NODE_SPEC_STATE_STARTED
	spec.Addrs = addrsStrs
	if err := d.daemonDb.SaveNode(spec); err != nil {
		d.log.WithError(err).Error("unable to save updated node state")
	}

	return n, nil
}

// flushNodeSpec flushes a node state spec to the db.
func (d *Daemon) flushNodeSpec(id string) error {
	nodInter, ok := d.runningNodes.Load(id)
	if !ok {
		return errors.Errorf("node with id not found: %s", id)
	}

	nod := nodInter.(*node.Node)
	nodSpec, err := nod.BuildSpec()
	if err != nil {
		return err
	}

	return d.daemonDb.SaveNode(nodSpec)
}
