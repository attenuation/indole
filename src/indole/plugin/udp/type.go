package udp

import (
	"net"
)

// UDP ...
type UDP struct {
	conn *net.UDPConn
	addr *net.UDPAddr
}

// Close ...
func (thisptr *UDP) Close() error {
	return thisptr.conn.Close()
}

// Read ...
func (thisptr *UDP) Read(p []byte) (n int, err error) {
	for {
		n, addr, err := thisptr.conn.ReadFromUDP(p)
		if addr.IP.Equal(thisptr.addr.IP) && addr.Port == thisptr.addr.Port && addr.Zone == thisptr.addr.Zone {
			return n, err
		}
	}
}

// Write ...
func (thisptr *UDP) Write(p []byte) (n int, err error) {
	return thisptr.conn.WriteToUDP(p, thisptr.addr)
}
