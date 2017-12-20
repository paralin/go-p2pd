package daemon

import (
	"github.com/libp2p/go-maddr-filter"
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
