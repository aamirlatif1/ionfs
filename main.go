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
		Decoder:       p2p.StringDecoder{},
	}
	tr := p2p.NewTCPTransport(tcpOpts)

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Hello, World!")

	select {}

}
