package main

import (
	"github.con/reward-rabieth/decentralized/p2p"
	"log"
)

func main() {
	tcpOpts := p2p.TCPTransportOpts{
		ListenAdr:     ":3000",
		Decoder:       p2p.DefaultDecoder{},
		HandshakeFunc: p2p.NOPHandshakeFunc,
	}
	tr := p2p.NewTCPTransport(tcpOpts)
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	select {}

}
