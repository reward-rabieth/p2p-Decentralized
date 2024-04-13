package p2p

import "net"

// Message holds any arbitary data is  that is being sent over the
//
//	transport between two nodes in the network
type Message struct {
	From    net.Addr
	Payload []byte
}
