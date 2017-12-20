package daemon

import (
	"github.com/Sirupsen/logrus"
	"github.com/multiformats/go-multiaddr"
	"github.com/pkg/errors"
)

// Config contains the daemon configuration.
type Config struct {
	// ApiListen contains the multiaddr the API listens to.
	ApiListen string
	// DataPath contains the path to a directory which contains system data.
	DataPath string
	// Log if set will be used as the root for logging.
	Log *logrus.Entry
}

// DefaultConfig builds the default configuration.
func DefaultConfig() *Config {
	return &Config{
		ApiListen: "/ip4/127.0.0.1/tcp/4050",
		DataPath:  "/var/lib/p2pd",
	}
}

// Validate checks the config for validity.
func (c *Config) Validate() error {
	if _, err := c.ParseListenAddress(); err != nil {
		return errors.WithMessage(err, "parse listen address as multiaddr")
	}
	if c.DataPath == "" {
		return errors.New("data path must be specified")
	}
	return nil
}

// ParseListenAddress parses the listen address to a multiaddr.
func (c *Config) ParseListenAddress() (multiaddr.Multiaddr, error) {
	return multiaddr.NewMultiaddr(c.ApiListen)
}
