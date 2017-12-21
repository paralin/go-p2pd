package daemon

import (
	"github.com/libp2p/go-maddr-filter"
	ma "github.com/multiformats/go-multiaddr"
	"net"
)

// parseAddrFilters parses a list of maddr filters.
func parseAddrFilters(filters []string) *filter.Filters {
	addrFilters := []string(filters)
	addrFilter := filter.NewFilters()
	for _, s := range addrFilters {
		// note: filters are checked in config validation step.
		_, ipnet, _ := net.ParseCIDR(s)
		addrFilter.AddDialFilter(ipnet)
	}
	return addrFilter
}

// multiAddrsToString converts a list of multiaddrs to a list of strings.
func multiAddrsToString(addrs []ma.Multiaddr) (outAddrs []string) {
	outAddrs = make([]string, len(addrs))
	for i := 0; i < len(addrs); i++ {
		outAddrs[i] = addrs[i].String()
	}
	return
}
