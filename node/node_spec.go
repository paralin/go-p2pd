package node

import (
	"errors"

	"github.com/Sirupsen/logrus"
	crypto "github.com/libp2p/go-libp2p-crypto"
	ma "github.com/multiformats/go-multiaddr"
)

// NodeSpec represents a node specification in the database.
type NodeSpec struct {
	// ID is the user specified slug/ID for this node.
	ID string `storm:"id"`
	// PrivKey is the private key for this node.
	PrivKey []byte
	// State is the node state for this node.
	State NodeSpecState
	// Addrs are the list of addresses the node should listen on.
	Addrs []string
}

// NewNodeSpec builds a new NodeSpec.
func NewNodeSpec(id string, privKey crypto.PrivKey, state NodeSpecState, addrs []string) (*NodeSpec, error) {
	privData, err := crypto.MarshalPrivateKey(privKey)
	if err != nil {
		return nil, err
	}

	return &NodeSpec{ID: id, PrivKey: privData, State: state, Addrs: addrs}, nil
}

// Validate checks the spec.
func (s *NodeSpec) Validate() error {
	if s == nil {
		return errors.New("spec cannot be nil")
	}
	if s.ID == "" {
		return errors.New("id cannot be empty")
	}
	if len(s.PrivKey) == 0 {
		return errors.New("priv_key cannot be empty")
	}
	return nil
}

// UnmarshalPrivKey unmarshals the spec private key.
func (s *NodeSpec) UnmarshalPrivKey() (crypto.PrivKey, error) {
	return crypto.UnmarshalPrivateKey(s.PrivKey)
}

// AddAddress adds a listen address.
func (s *NodeSpec) AddAddress(addr ma.Multiaddr) {
	s.Addrs = append(s.Addrs, addr.String())
}

// LogFields adds logging fields to an entry.
func (s *NodeSpec) LogFields(le *logrus.Entry) *logrus.Entry {
	nod := s
	return le.
		WithField("id", nod.ID).
		WithField("state", nod.State.String()).
		WithField("naddrs", len(nod.Addrs))
}
