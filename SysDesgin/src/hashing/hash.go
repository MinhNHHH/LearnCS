package hashing

import (
	"crypto/sha256"
	"encoding/binary"
	"sort"
	"strconv"
)

// Node represents a node in the hash ring.
type Node struct {
	ID   string
	Hash int
}

type ConsistentHashing struct {
	rings    map[uint64]string
	nodes    []uint64
	replicas int
}

func NewHashing(replicas int) *ConsistentHashing {
	return &ConsistentHashing{
		replicas: replicas,
	}
}

func (h *ConsistentHashing) hash(key string) uint64 {
	hash := sha256.Sum256([]byte(key))
	return binary.BigEndian.Uint64(hash[:8])
}

func (h *ConsistentHashing) sortNodes() {
	sort.Slice(h.nodes, func(i, j int) bool {
		return h.nodes[i] < h.nodes[j]
	})
}

func (h *ConsistentHashing) AddNode(node string) {
	for i := 0; i < h.replicas; i++ {
		virtualNode := node + "#" + strconv.Itoa(i)
		hash := h.hash(virtualNode)
		h.rings[hash] = node
		h.nodes = append(h.nodes, hash)
	}
	h.sortNodes()
}

func (h *ConsistentHashing) RemoveNode(node string) {
	for i := 0; i < h.replicas; i++ {
		virtualNode := node + "#" + strconv.Itoa(i)
		hash := h.hash(virtualNode)
		delete(h.rings, hash)
		for index, node := range h.nodes {
			if node == hash {
				h.nodes = append(h.nodes[:index], h.nodes[index+1:]...)
				break
			}
		}
	}
}

func (h *ConsistentHashing) GetNode(key string) (string, bool) {
	if len(h.rings) == 0 {
		return "", false
	}
	hash := h.hash(key)

	// Binary search to find the closest hash in sortedKeys
	idx := sort.Search(len(h.nodes), func(i int) bool {
		return h.nodes[i] >= hash
	})

	// Wrap around the ring if needed
	if idx == len(h.nodes) {
		idx = 0
	}

	nodeHash := h.nodes[idx]
	return h.rings[nodeHash], true
}
