package control

import (
	"errors"
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
