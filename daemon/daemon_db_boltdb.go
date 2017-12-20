package daemon

import (
	"github.com/asdine/storm"
	"github.com/paralin/go-p2pd/node"
	"github.com/pkg/errors"
	"path"
)

// daemonDatabaseBolt manages the daemon's on-disk database via boltdb.
type daemonDatabaseBolt struct {
	*Daemon
	db *storm.DB
}

// buildDaemonDatabaseBolt builds a new daemon database with bolt.
func buildDaemonDatabaseBolt(d *Daemon) (daemonDatabase, error) {
	dbPath := path.Join(d.config.DataPath, "daemon-boltdb.db")
	db, err := storm.Open(dbPath)
	if err != nil {
		return nil, errors.Errorf("unable to open db at %s: %v", dbPath, err.Error())
	}

	for _, typ := range daemonDatabaseTypes {
		if err := db.Init(typ); err != nil {
			return nil, err
		}
	}

	return &daemonDatabaseBolt{Daemon: d, db: db}, nil
}

// SaveNode creates or updates a node spec.
func (d *daemonDatabaseBolt) SaveNode(spec *node.NodeSpec) error {
	if err := spec.Validate(); err != nil {
		return err
	}

	return d.db.Save(spec)
}

// GetNode returns a node spec by ID.
// If not found, returns nil, nil, not an error.
func (d *daemonDatabaseBolt) GetNode(id string) (*node.NodeSpec, error) {
	spec := &node.NodeSpec{}
	if err := d.db.One("ID", id, spec); err != nil {
		if err == storm.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return spec, nil
}

// ListNodes returns a list of all node IDs.
func (d *daemonDatabaseBolt) ListNodes() (specs []*node.NodeSpec, err error) {
	err = d.db.All(&specs)
	return
}

// Close closes the database.
func (d *daemonDatabaseBolt) Close() error {
	return d.db.Close()
}

func init() {
	registerDatabaseImpl("boltdb", func(d *Daemon) (daemonDatabase, error) {
		return buildDaemonDatabaseBolt(d)
	})
}
