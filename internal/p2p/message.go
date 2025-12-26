package p2p

import "net"

// Message represents any arbitrary message sent between peers.
type Message struct {
	From    net.Addr
	Payload []byte
}
