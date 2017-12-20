package node

// NodeSpecState is a desired state of a node in the db.
type NodeSpecState uint32

const (
	// NodeSpecState_STOPPED indicates the node is inactive.
	NodeSpecState_STOPPED NodeSpecState = iota
	// NodeSpecState_STARTED indicates the node should be running.
	NodeSpecState_STARTED
)

// String attempts to find a string value for the NodeSpecState.
func (n NodeSpecState) String() string {
	switch n {
	case NodeSpecState_STARTED:
		return "STARTED"
	case NodeSpecState_STOPPED:
		return "STOPPED"
	}
	return "UNKNOWN"
}
