package control

import (
	ma "github.com/multiformats/go-multiaddr"
	"github.com/pkg/errors"
)

// Validate does cursory validation of the message.
func (r *CreateNodeRequest) Validate() error {
	if r.GetNodeId() == "" {
		return errors.New("node id cannot be empty")
	}
	return nil
}

// Validate does cursory validation of the message.
func (r *StartNodeRequest) Validate() error {
	if r.GetNodeId() == "" {
		return errors.New("node id cannot be empty")
	}
	return nil
}

// Validate does cursory validation of the message.
func (r *ListenNodeRequest) Validate() error {
	if r.GetNodeId() == "" {
		return errors.New("node id cannot be empty")
	}
	if r.GetAddr() == "" {
		return errors.New("addr to listen to cannot be empty")
	}
	if _, err := ma.NewMultiaddr(r.GetAddr()); err != nil {
		return errors.Errorf("addr invalid: %s", err.Error())
	}
	return nil
}

// Validate does cursory validation of the message.
func (r *StatusNodeRequest) Validate() error {
	if r.GetNodeId() == "" {
		return errors.New("node id cannot be empty")
	}
	return nil
}
