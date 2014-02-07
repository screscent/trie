//author gonghh
//Copyright 2014 gonghh. All rights reserved.
package trie

type TrieNode struct {
	Next   []*TrieNode
	IsLeaf bool
	Data   interface{}
}

//byte: 0 - 0xFF
const maxTrieNodeSize = 0x100

var root *TrieNode = nil

func init() {
	root = NewTriNode()
}

func NewTriNode() *TrieNode {
	n := &TrieNode{make([]*TrieNode, maxTrieNodeSize, maxTrieNodeSize), false, nil}
	for i := 0; i < maxTrieNodeSize; i++ {
		n.Next[i] = nil
	}
	return n
}

func (tn *TrieNode) Add(key string, data interface{}) {
	key_len := len(key)

	idx := tn
	for i := 0; i < key_len; i++ {
		next := idx.Next[key[i]]
		if next == nil {
			next = NewTriNode()
			idx.Next[key[i]] = next
		}

		idx = next
	}

	idx.Data = data
	idx.IsLeaf = true
}

func Add(key string, data interface{}) {
	root.Add(key, data)
}

func (tn *TrieNode) Get(key string) (interface{}, bool) {
	key_len := len(key)

	idx := tn
	for i := 0; i < key_len; i++ {
		idx = idx.Next[key[i]]
		if idx == nil {
			return nil, false
		}
	}
	return idx.Data, idx.IsLeaf
}

func Get(key string) (interface{}, bool) {
	key_len := len(key)

	idx := root
	for i := 0; i < key_len; i++ {
		idx = idx.Next[key[i]]
		if idx == nil {
			return nil, false
		}
	}
	return idx.Data, idx.IsLeaf
}
