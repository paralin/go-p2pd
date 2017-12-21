package node

import (
	"context"
	"crypto/rand"
	"errors"
	"sync"

	"github.com/Sirupsen/logrus"
	crypto "github.com/libp2p/go-libp2p-crypto"
	peer "github.com/libp2p/go-libp2p-peer"
	pstore "github.com/libp2p/go-libp2p-peerstore"
	swarm "github.com/libp2p/go-libp2p-swarm"
	bhost "github.com/libp2p/go-libp2p/p2p/host/basic"
	ma "github.com/multiformats/go-multiaddr"
)

// Node contains an active libp2p node.
type Node struct {
	id        string
	peerId    peer.ID
	peerStore pstore.Peerstore
	le        *logrus.Entry

	mtx    sync.RWMutex
	net    *swarm.Network
	host   *bhost.BasicHost
	priv   crypto.PrivKey
	pub    crypto.PubKey
	laddrs []ma.Multiaddr
}

// nodeFromKeys builds a node from the specified private and public keypair.
func nodeFromKeys(id string, priv crypto.PrivKey, pub crypto.PubKey, le *logrus.Entry, laddrs []ma.Multiaddr) (*Node, error) {
	pid, err := peer.IDFromPublicKey(pub)
	if err != nil {
		return nil, err
	}

	le = le.WithField("id", id).
		WithField("state", NodeSpecState_NODE_SPEC_STATE_STOPPED)

	ps := pstore.NewPeerstore()
	ps.AddPrivKey(pid, priv)
	ps.AddPubKey(pid, pub)

	return &Node{
		id:        id,
		le:        le,
		peerStore: ps,
		peerId:    pid,
		priv:      priv,
		pub:       pub,
		laddrs:    laddrs,
	}, nil
}

// NewNode builds a new node from scratch, generating an identity.
func NewNode(id string, le *logrus.Entry) (*Node, error) {
	priv, pub, err := crypto.GenerateEd25519Key(rand.Reader)
	if err != nil {
		return nil, err
	}

	return nodeFromKeys(id, priv, pub, le, nil)
}

// FromSpec builds a node from its spec.
func FromSpec(spec *NodeSpec, le *logrus.Entry) (*Node, error) {
	if err := spec.Validate(); err != nil {
		return nil, err
	}

	privKey, err := spec.UnmarshalPrivKey()
	if err != nil {
		return nil, err
	}

	var laddrs []ma.Multiaddr
	for _, addr := range spec.Addrs {
		maddr, err := ma.NewMultiaddr(addr)
		if err != nil {
			le.WithError(err).Warn("invalid listen address in db")
			continue
		}
		laddrs = append(laddrs, maddr)
	}

	return nodeFromKeys(spec.ID, privKey, privKey.GetPublic(), le, laddrs)
}

// StartWithAddrs starts the node.
func (n *Node) StartWithAddrs(ctx context.Context, listenAddrs []ma.Multiaddr) error {
	n.mtx.Lock()
	defer n.mtx.Unlock()

	if n.host != nil {
		return errors.New("node already started")
	}

	n.le.Debug("starting node")

	// merge listen addrs lists
	for i := 0; i < len(listenAddrs); i++ {
		for _, el := range n.laddrs {
			if el.Equal(listenAddrs[i]) {
				listenAddrs[i] = listenAddrs[len(listenAddrs)-1]
				listenAddrs[len(listenAddrs)-1] = nil
				listenAddrs = listenAddrs[:len(listenAddrs)-1]
				i--
				break
			}
		}
	}
	n.laddrs = append(n.laddrs, listenAddrs...)

	bnet, err := swarm.NewNetwork(ctx, listenAddrs, n.peerId, n.peerStore, nil)
	if err != nil {
		return err
	}

	host, err := bhost.NewHost(ctx, bnet, &bhost.HostOpts{})
	if err != nil {
		return err
	}

	n.net = bnet
	n.host = host
	n.le = n.le.WithField("state", NodeSpecState_NODE_SPEC_STATE_STARTED.String())
	n.le.Info("started node")
	for _, addr := range n.laddrs {
		n.le.WithField("addr", addr.String()).Debug("swarm listening on addr")
	}
	return nil
}

// GetId returns the node ID.
func (n *Node) GetId() string {
	return n.id
}

// GetPeerId returns the peer ID of the node.
func (n *Node) GetPeerId() peer.ID {
	return n.peerId
}

// BuildSpec returns the specification for this node.
func (n *Node) BuildSpec() (*NodeSpec, error) {
	state := NodeSpecState_NODE_SPEC_STATE_STOPPED
	if n.net != nil {
		state = NodeSpecState_NODE_SPEC_STATE_STARTED
	}

	var addrs []string
	for _, laddr := range n.laddrs {
		addrs = append(addrs, laddr.String())
	}

	return NewNodeSpec(n.id, n.priv, state, addrs)
}

// Close closes the node.
func (n *Node) Close() (closeErr error) {
	n.mtx.Lock()
	defer n.mtx.Unlock()

	if n.host != nil {
		closeErr = n.host.Close()
		n.host = nil
	}
	if n.net != nil {
		err := n.net.Close()
		if err != nil && closeErr == nil {
			closeErr = err
		}
		n.net = nil
	}
	return
}
