package node

import (
	"context"
	"crypto/rand"
	"errors"
	"sync"

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

	mtx    sync.RWMutex
	net    *swarm.Network
	host   *bhost.BasicHost
	priv   crypto.PrivKey
	pub    crypto.PubKey
	laddrs []ma.Multiaddr
}

// nodeFromKeys builds a node from the specified private and public keypair.
func nodeFromKeys(id string, priv crypto.PrivKey, pub crypto.PubKey) (*Node, error) {
	pid, err := peer.IDFromPublicKey(pub)
	if err != nil {
		return nil, err
	}

	ps := pstore.NewPeerstore()
	ps.AddPrivKey(pid, priv)
	ps.AddPubKey(pid, pub)

	return &Node{
		id:        id,
		peerStore: ps,
		peerId:    pid,
		priv:      priv,
		pub:       pub,
	}, nil
}

// NewNode builds a new node from scratch, generating an identity.
func NewNode(id string) (*Node, error) {
	priv, pub, err := crypto.GenerateEd25519Key(rand.Reader)
	if err != nil {
		return nil, err
	}

	return nodeFromKeys(id, priv, pub)
}

// FromSpec builds a node from its spec.
func FromSpec(spec *NodeSpec) (*Node, error) {
	if err := spec.Validate(); err != nil {
		return nil, err
	}

	privKey, err := spec.UnmarshalPrivKey()
	if err != nil {
		return nil, err
	}

	return nodeFromKeys(spec.ID, privKey, privKey.GetPublic())
}

// GetListenAddrs returns the addresses the node has been started with.
func (n *Node) GetListenAddrs() []ma.Multiaddr {
	n.mtx.RLock()
	defer n.mtx.RUnlock()

	return n.laddrs
}

// StartWithAddrs starts the node.
func (n *Node) StartWithAddrs(ctx context.Context, listenAddrs []ma.Multiaddr) error {
	n.mtx.Lock()
	defer n.mtx.Unlock()

	if n.host != nil {
		return errors.New("node already started")
	}

	n.laddrs = listenAddrs
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
	return NewNodeSpec(n.id, n.priv)
}

// GetPeerMultiaddr returns the p2p multiaddr representing this peer.
func (n *Node) GetPeerMultiaddr() ma.Multiaddr {
	maddr, err := ma.NewMultiaddr("/p2p/" + string(n.peerId))
	if err != nil {
		panic(err)
	}
	return maddr
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
