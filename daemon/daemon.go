package daemon

import (
	"context"
	"os"
	"sync"

	"github.com/Sirupsen/logrus"
	"github.com/libp2p/go-maddr-filter"
	mnet "github.com/multiformats/go-multiaddr-net"
	"github.com/paralin/go-p2pd/control"
	"github.com/paralin/go-p2pd/node"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// Daemon is an instance of the P2PD daemon.
type Daemon struct {
	config       Config
	log          *logrus.Entry
	runningNodes sync.Map
	daemonDb     daemonDatabase
	addrFilters  *filter.Filters

	ctx       context.Context
	ctxCancel context.CancelFunc
}

// NewDaemon builds a new P2PD daemon.
func NewDaemon(conf Config) (*Daemon, error) {
	if err := conf.Validate(); err != nil {
		return nil, errors.WithMessage(err, "validate config")
	}

	log := conf.Log
	if log == nil {
		logger := logrus.New()
		logger.SetLevel(logrus.DebugLevel)
		log = logrus.NewEntry(logger)
	}

	if err := os.MkdirAll(conf.DataPath, 0755); err != nil {
		return nil, err
	}

	var err error
	d := &Daemon{config: conf, log: log}
	d.addrFilters = parseAddrFilters(conf.AddrFilters)
	dbCtor := daemonDatabaseImpls[conf.DataDb]
	d.daemonDb, err = dbCtor(d)
	if err != nil {
		return nil, err
	}
	log.
		WithField("db-type", conf.DataDb).
		WithField("db-path", conf.DataPath).
		Debug("opened database")

	return d, nil
}

// Run manages the daemon and returns any fatal errors.
func (d *Daemon) Run(ctx context.Context) error {
	le := d.log
	d.ctx, d.ctxCancel = context.WithCancel(ctx)
	defer d.Close()

	// Parse our API listen multiaddr.
	apiMultiaddr, err := d.config.ParseListenAddress()
	if err != nil {
		return err
	}

	apiListener, err := mnet.Listen(apiMultiaddr)
	if err != nil {
		return err
	}
	defer apiListener.Close()

	// Attempt to load and activate all registered nodes.
	if err := d.loadStartDbNodes(); err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	control.RegisterControlServiceServer(
		grpcServer,
		newDaemonControlServer(d),
	)

	le.
		WithField("multiaddr", apiMultiaddr).
		Info("api listening at address")

	errCh := make(chan error, 1)
	go func() {
		errCh <- grpcServer.Serve(apiListener.NetListener())
	}()

	return <-errCh
}

// RegisterNode registers an existing node.
func (d *Daemon) RegisterNode(n *node.Node) error {
	id := n.GetId()
	_, loaded := d.runningNodes.LoadOrStore(id, n)
	if loaded {
		return errors.Errorf("node already running with id: %s", id)
	}

	// if the node is not running, allow overwriting it.
	spec, err := n.BuildSpec()
	if err != nil {
		return err
	}
	if err := d.daemonDb.SaveNode(spec); err != nil {
		return err
	}

	return nil
}

// Close closes the daemon, normally done automatically when Run() returns.
func (d *Daemon) Close() {
	d.daemonDb.Close()
	d.ctxCancel()
}
