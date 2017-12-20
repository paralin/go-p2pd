package daemon

import (
	"sync"

	"github.com/Sirupsen/logrus"
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

	return &Daemon{config: conf, log: log}, nil
}

// Run manages the daemon and returns any fatal errors.
func (d *Daemon) Run() error {
	le := d.log

	// Parse our API listen multiaddr.
	apiMultiaddr, err := d.config.ParseListenAddress()
	if err != nil {
		return err
	}

	apiListener, err := mnet.Listen(apiMultiaddr)
	if err != nil {
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
		return errors.Errorf("node already registered with id: %s", id)
	}
	return nil
}
