package control

import (
	"bytes"
	"fmt"
)

// NodeListenResponse is a response that contains a node listen result.
type NodeListenResponse interface {
	// GetNodePeerId returns the node peer id.
	GetNodePeerId() string
	// GetNodeListenAddrs returns the list of addrs that the node is now listening on.
	GetNodeListenAddrs() []string
}

// NodeListenResponseString makes a readable string of the node listen response.
func NodeListenResponseString(resp NodeListenResponse) string {
	p2pIdStr := "/p2p/" + resp.GetNodePeerId()
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "Swarm listening on\n")
	for _, addr := range resp.GetNodeListenAddrs() {
		fmt.Fprintf(&buf, "\t%s%s\n", addr, p2pIdStr)
	}
	return (&buf).String()
}
