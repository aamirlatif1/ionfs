package p2p

import "net"

// TCPPeer represents a remote node over a TCP connection.
type TCPPeer struct {
	conn net.Conn

	// if we dial and retrieve a connection, it's outbound
	// if we accept a connection, it's inbound
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

func (p *TCPPeer) Close() error {
	return p.conn.Close()
}
