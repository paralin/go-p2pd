package node

// NodeSpec represents a node specification in the database.
type NodeSpec struct {
	// ID is the user specified slug/ID for this node.
	ID string `storm:"id"`
}
