package p2p_test

import (
	"testing"

	"github.com/aamirlatif1/ionfs/internal/p2p"
	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr:    ":3000",
		HandshakeFunc: p2p.NOOPHandshake,
		Decoder:       p2p.GOBDecoder{},
	}

	tr := p2p.NewTCPTransport(tcpOpts)
	err := tr.ListenAndAccept()

	assert.Equal(t, tcpOpts.ListenAddr, tr.ListenAddr)
	assert.NoError(t, err)
}
