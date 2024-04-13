package p2p

// HandshakeFunc HandShakeFunc
type HandshakeFunc func(Peer) error

func NOPHandshakeFunc(Peer) error { return nil }
