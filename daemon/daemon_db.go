package daemon

import (
	"github.com/paralin/go-p2pd/node"
)

// daemonDatabase is a database implementation.
type daemonDatabase interface {
	// SaveNode creates or updates a node spec.
	SaveNode(spec *node.NodeSpec) error
	// GetNode returns a node spec by ID.
	// If not found, returns nil, nil, not an error.
	GetNode(id string) (*node.NodeSpec, error)
	// ListNodes returns a list of all node specs.
	ListNodes() ([]*node.NodeSpec, error)
	// Close closes the database.
	Close() error
}

// daemonDatabaseCtor constructs a database implementation.
type daemonDatabaseCtor func(d *Daemon) (daemonDatabase, error)

// daemonDatabaseImpls maps implementation ID to constructor.
var daemonDatabaseImpls = make(map[string]daemonDatabaseCtor)

// daemonDatabaseTypes contains a list of types the db will store.
var daemonDatabaseTypes = []interface{}{&node.NodeSpec{}}

// registerDatabaseImpl registers a compiled database implementation.
func registerDatabaseImpl(id string, ctor daemonDatabaseCtor) {
	daemonDatabaseImpls[id] = ctor
}
