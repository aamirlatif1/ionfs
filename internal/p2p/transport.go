package p2p

// Peer represents a node in the network.
type Peer interface {
}

// Transport handles the communication between the node in the network.
// This can be of the form of TCP, UDP, WebSocket, etc.
type Transport interface {
	ListenAndAccept() error
}
