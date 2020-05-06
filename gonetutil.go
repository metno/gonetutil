package gonetutil

import (
	"net"
	"os"
	"strings"
)

// Fqdn - Get fully qualified domain name
func Fqdn() string {
	// There seems not to be a native method in go.
	// This is taken from here: https://github.com/Showmax/go-fqdn (Apache Licence)
	// Moved to a local repo to guarantee availability

	hostname, err := os.Hostname()
	if err != nil {
		return "unknown"
	}

	addrs, err := net.LookupIP(hostname)
	if err != nil {
		return hostname
	}

	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			ip, err := ipv4.MarshalText()
			if err != nil {
				return hostname
			}
			hosts, err := net.LookupAddr(string(ip))
			if err != nil || len(hosts) == 0 {
				return hostname
			}
			fqdn := hosts[0]
			return strings.TrimSuffix(fqdn, ".") // return fqdn without trailing dot
		}
	}
	return hostname
}
