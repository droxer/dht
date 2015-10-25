package kademlia

import (
	"github.com/droxer/lru"
	"math/rand"
	"net"
)

const IDLength = 20

type nodeID [IDLength]byte

func newID() (ret nodeID) {
	for i := 0; i < IDLength; i++ {
		ret[i] = uint8(rand.Intn(256))
	}
	return
}

func (n *nodeID) XOR(other nodeID) (ret nodeID) {
	for i := 0; i < IDLength; i++ {
		ret[i] = n[i] ^ other[i]
	}
	return
}

type contact struct {
	id   nodeID
	ip   net.IP
	port uint16
}

type kBucket *lru.Cache

func newKBucket() kBucket {
	return lru.New(K)
}

type node struct {
	id      nodeID
	kBucket kBucket
}

func newNode() *node {
	return &node{
		id:      newID(),
		kBucket: newKBucket(),
	}
}
