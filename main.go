package main

import (
	"fmt"
	"log"

	"github.com/aamirlatif1/ionfs/internal/p2p"
)

func main() {

	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr:    ":3000",
		HandshakeFunc: p2p.NOOPHandshake,
		Decoder:       p2p.DefaultDecoder{},
		OnPeer: func(peer p2p.Peer) error {
			fmt.Printf("New peer connected: %+v\n", peer)
			return nil
		},
	}
	tr := p2p.NewTCPTransport(tcpOpts)

	go func() {
		for rpc := range tr.Consume() {
			fmt.Printf("Received RPC from %s: %+v\n", rpc.From, rpc)
		}
	}()

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Hello, World!")

	select {}

}
