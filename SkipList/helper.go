package main

import (
	"math/rand"
)

type SkipList struct {
	maxHeight int
	height    int
	size      int
	head      *SkipListNode
}

type SkipListNode struct {
	key   string
	value []byte
	next  []*SkipListNode
}

func (s *SkipList) roll() int {
	level := 0
	// possible ret values from rand are 0 and 1
	// we stop shen we get a 0
	for ; rand.Int31n(2) == 1; level++ {
		if level >= s.maxHeight {
			if level > s.height {
				s.height = level
			}
			return level
		}
	}
	if level > s.height {
		s.height = level
	}
	return level
}

func newNode(key_ string, node1 *SkipListNode, node2 *SkipListNode) *SkipListNode {
	nodes := []*SkipListNode{node1, node2}
	data := []byte{0}
	sln := SkipListNode{key: key_, value: data, next: nodes} // dodato je ovo data iznad
	return &sln
}
