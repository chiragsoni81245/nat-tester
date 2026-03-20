package types

import "net"

type NATType string

const (
	OpenInternet   NATType = "Open Internet"
	FullCone       NATType = "Full Cone NAT (likely)"
	RestrictedCone NATType = "Restricted Cone NAT (likely)"
	PortRestricted NATType = "Port Restricted NAT (likely)"
	SymmetricNAT   NATType = "Symmetric NAT"
	Unknown        NATType = "Unknown"
)

type Result struct {
	Server string
	Addr   *net.UDPAddr
}
