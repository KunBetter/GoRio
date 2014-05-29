// SliceTrie
package GoRio

import (
	"unicode/utf8"
)

type SliceTrieNode struct {
	Key    byte
	ID     int
	Childs []*SliceTrieNode
	Flag   bool
}

func NewSliceTrieNode(key byte, id int) *SliceTrieNode {
	node := &SliceTrieNode{
		Key:    key,
		ID:     id,
		Childs: []*SliceTrieNode{},
		Flag:   false,
	}
	return node
}

//***************LinkedTrie****************

type SliceTrie struct {
	Root *SliceTrieNode
}

//***************PUBLIC****************

func NewSliceTrie() *SliceTrie {
	trie := &SliceTrie{
		Root: NewSliceTrieNode(0, -1),
	}
	return trie
}

func (trie *SliceTrie) AddWord(word string, id int) {
	key := []byte(word)
	trie.AddKey(key, id)
}

func (trie *SliceTrie) FindWord(word string) bool {
	key := []byte(word)
	return trie.FindKey(key)
}

func (trie *SliceTrie) MaxSegments(word string) []Token {
	key := []rune(word)
	return trie.MaxSegs(key)
}

func (trie *SliceTrie) MaxSegs(key []rune) []Token {
	if len(key) <= 1 {
		return []Token{}
	}
	token := []Token{}
	node := trie.Root
	if node == nil {
		return nil
	}
	i := 0
	path := []byte{}
	p := []byte{0, 0, 0, 0}
L:
	for i = 0; i < len(key); i++ {
		pLen := utf8.EncodeRune(p, key[i])
		for j := 0; j < pLen; j++ {
			childs := node.Childs
			pos, found := trie.FindInChilds(childs, p[j])
			if !found {
				break L
			}
			node = childs[pos]
		}
		path = append(path, p[0:pLen]...)
		if node.Flag {
			token = append(token, Token{string(path), node.ID})
		}
	}
	token = append(token, trie.MaxSegs(key[1:])...)
	return token
}

func (trie *SliceTrie) FindKey(key []byte) bool {
	node := trie.Root
	if node == nil {
		return false
	}
	i := 0
	for i = 0; i < len(key); i++ {
		childs := node.Childs
		pos, found := trie.FindInChilds(childs, key[i])
		if !found {
			break
		}
		node = childs[pos]
	}
	if i != len(key) || !node.Flag {
		return false
	}
	return true
}

//***************PRIVATE****************

func (trie *SliceTrie) FindInChilds(childs []*SliceTrieNode, key byte) (int, bool) {
	clen := len(childs)
	if clen <= 0 {
		return 0, false
	}
	l := 0
	h := clen - 1
	for l <= h {
		m := (l + h) / 2
		if childs[m].Key < key {
			l = m + 1
		} else if childs[m].Key > key {
			h = m - 1
		} else {
			return m, true
		}
	}
	return l, false
}

func (trie *SliceTrie) FindLastMatchNode(key []byte) (*SliceTrieNode, int) {
	node := trie.Root
	var last *SliceTrieNode = nil
	if node == nil {
		return nil, -1
	}
	lastI := 0
	i := 0
	for i = 0; i < len(key); i++ {
		childs := node.Childs
		pos, found := trie.FindInChilds(childs, key[i])
		if !found {
			break
		}
		node = childs[pos]
		last = node
	}
	lastI = i
	return last, lastI
}

func (trie *SliceTrie) AddKey(key []byte, id int) {
	node, lastI := trie.FindLastMatchNode(key)
	if node == nil {
		node = trie.Root
		lastI = 0
	}
	for i := lastI; i < len(key); i++ {
		newNode := NewSliceTrieNode(key[i], -1)
		clen := len(node.Childs)
		if clen <= 0 {
			node.Childs = append(node.Childs, newNode)
		} else if len(node.Childs) == 1 {
			node.Childs = append(node.Childs, newNode)
			if node.Childs[0].Key > key[i] {
				node.Childs[0], node.Childs[1] = node.Childs[1], node.Childs[0]
			}
		} else {
			pos, _ := trie.FindInChilds(node.Childs, key[i])
			node.Childs = append(node.Childs, newNode)
			copy(node.Childs[pos+1:], node.Childs[pos:clen])
			node.Childs[pos] = newNode
		}
		node = newNode
	}
	node.Flag = true
	node.ID = id
}
