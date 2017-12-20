package ctl

import (
	"context"
	"net"
	"time"

	ma "github.com/multiformats/go-multiaddr"
	mnet "github.com/multiformats/go-multiaddr-net"
	"github.com/pkg/errors"
)

// MultiaddrDialer is a GRPC dialer that accepts multi-addr addresses.
var MultiaddrDialer = func(addr string, timeout time.Duration) (net.Conn, error) {
	maddr, err := ma.NewMultiaddr(addr)
	if err != nil {
		return nil, errors.Wrap(err, "parse multiaddr")
	}

	ctx := context.Background()
	if timeout != 0 {
		ctxSub, ctxSubCancel := context.WithTimeout(ctx, timeout)
		defer ctxSubCancel()
		ctx = ctxSub
	}

	dialer := mnet.Dialer{}
	conn, err := dialer.DialContext(ctx, maddr)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
