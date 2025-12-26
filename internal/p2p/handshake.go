package p2p

// HandshakeFunc defines the function signature for handshake functions.
type HandshakeFunc func(Peer) error

func NOOPHandshake(Peer) error {
	return nil
}
