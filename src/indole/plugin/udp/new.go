package udp

import (
	"io"
	"log"
	"net"
)

// Build ...
func Build(args *Args) io.ReadWriteCloser {
	addr, err := net.ResolveUDPAddr(args.Network, args.Address)
	if err != nil {
		log.Println("plugin", "udp", "New", err)
		return nil
	}

	conn, err := net.ListenUDP(args.Network, addr)
	if err != nil {
		log.Println("vertex", "udp", "new", "New", "net.ListenUDP", err)
	}

	remote, err := net.ResolveUDPAddr(args.Network, args.Remote)

	ret := &UDP{
		conn: conn,
		addr: remote,
	}
	return ret
}

// Args ...
type Args struct {
	Network string `xml:"network,attr"`
	Address string `xml:"address,attr"`
	Remote  string `xml:"remote,attr"`
}
