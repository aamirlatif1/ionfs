package p2p

import "net"

// RPC represents any arbitrary message sent between peers.
type RPC struct {
	From    net.Addr
	Payload []byte
}
