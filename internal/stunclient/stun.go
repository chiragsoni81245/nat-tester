package stunclient

import (
	"net"
	"time"

	"github.com/pion/stun"
)

func Query(conn *net.UDPConn, server string, timeout time.Duration) (*net.UDPAddr, error) {
	raddr, err := net.ResolveUDPAddr("udp", server)
	if err != nil {
		return nil, err
	}

	msg := stun.MustBuild(stun.TransactionID, stun.BindingRequest)

	_, err = conn.WriteTo(msg.Raw, raddr)
	if err != nil {
		return nil, err
	}

	conn.SetReadDeadline(time.Now().Add(timeout))

	buf := make([]byte, 1500)

	for {
		n, _, err := conn.ReadFromUDP(buf)
		if err != nil {
			return nil, err
		}

		var res stun.Message
		res.Raw = buf[:n]

		if err := res.Decode(); err != nil {
			continue
		}

		var xorAddr stun.XORMappedAddress
		if err := xorAddr.GetFrom(&res); err != nil {
			continue
		}

		return &net.UDPAddr{
			IP:   xorAddr.IP,
			Port: xorAddr.Port,
		}, nil
	}
}
