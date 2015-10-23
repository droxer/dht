package kademlia

import (
	"net"
)

const (
	IDBytes = 20
)

type nodeID [IDBytes]byte

func (n *nodeID) XOR(other nodeID) (ret nodeID) {
	for i := 0; i < IDBytes; i++ {
		ret[i] = n[i] ^ other[i]
	}
	return
}

type bucket struct {
	contacts [K]*contact
}

type contact struct {
	id   nodeID
	ip   net.IP
	port uint16
}

type Peer struct {
	ID nodeID
}

func NewPeer() *Peer {
	return nil
}
