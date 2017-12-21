package node

import (
	ma "github.com/multiformats/go-multiaddr"
)

// GetPeerMultiaddr returns the p2p multiaddr representing this peer.
func (n *Node) GetPeerMultiaddr() ma.Multiaddr {
	maddr, err := ma.NewMultiaddr("/p2p/" + string(n.peerId))
	if err != nil {
		panic(err)
	}
	return maddr
}

// GetListenAddrs returns the addresses the node has been started with.
func (n *Node) GetListenAddrs() []ma.Multiaddr {
	n.mtx.RLock()
	defer n.mtx.RUnlock()

	return n.laddrs
}

// AddListenAddr adds a listening address to the node.
func (n *Node) AddListenAddr(maddr ma.Multiaddr) error {
	n.mtx.Lock()
	defer n.mtx.Unlock()

	for _, eaddr := range n.laddrs {
		if eaddr.Equal(maddr) {
			return nil
		}
	}

	n.le.WithField("addr", maddr.String()).Info("listening on new addr")
	n.laddrs = append(n.laddrs, maddr)
	return n.net.Listen(maddr)
}
