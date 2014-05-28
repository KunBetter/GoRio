// Trie
package LinkedTrie

import (
	"unicode/utf8"
)

type TrieNode struct {
	Key        byte
	ID         int
	Silbling   *TrieNode
	FirstChild *TrieNode
	flag       bool
}

func NewTrieNode(key byte, id int) *TrieNode {
	node := &TrieNode{
		Key:        key,
		ID:         id,
		Silbling:   nil,
		FirstChild: nil,
		flag:       false,
	}
	return node
}

//***************LinkedTrie****************

type LinkedTrie struct {
	Root *TrieNode
}

//***************PUBLIC****************

func NewLinkedTrie() *LinkedTrie {
	trie := &LinkedTrie{
		Root: NewTrieNode(0, -1),
	}
	return trie
}

func (trie *LinkedTrie) AddWord(word string, id int) {
	key := []byte(word)
	trie.AddKey(key, id)
}

func (trie *LinkedTrie) FindWord(word string) bool {
	key := []byte(word)
	return trie.FindKey(key)
}

func (trie *LinkedTrie) MaxSegments(word string) ([]string, []int) {
	key := []rune(word)
	return trie.MaxSegs(key)
}

func (trie *LinkedTrie) MaxSegs(key []rune) ([]string, []int) {
	if len(key) <= 1 {
		return []string{}, []int{}
	}
	words := []string{}
	ids := []int{}
	node := trie.Root
	if node == nil {
		return nil, nil
	}
	i := 0
	path := []byte{}
	p := []byte{0, 0, 0, 0}
L:
	for i = 0; i < len(key); i++ {
		pLen := utf8.EncodeRune(p, key[i])
		for j := 0; j < pLen; j++ {
			node = node.FirstChild
			node = trie.FindInSilbling(node, p[j])
			if node == nil {
				break L
			}
		}
		path = append(path, p[0:pLen]...)
		if node.flag {
			words = append(words, string(path))
			ids = append(ids, node.ID)
		}
	}
	nw, nid := trie.MaxSegs(key[1:])
	words = append(words, nw...)
	ids = append(ids, nid...)
	return words, ids
}

func (trie *LinkedTrie) FindKey(key []byte) bool {
	node := trie.Root
	if node == nil {
		return false
	}
	i := 0
	for i = 0; i < len(key); i++ {
		node = node.FirstChild
		node = trie.FindInSilbling(node, key[i])
		if node == nil {
			break
		}
	}
	if i != len(key) || !node.flag {
		return false
	}
	return true
}

//***************PRIVATE****************

func (trie *LinkedTrie) FindInSilbling(node *TrieNode, key byte) *TrieNode {
	if node == nil {
		return nil
	}
	tNode := node
	for tNode != nil {
		if key == tNode.Key {
			break
		}
		tNode = tNode.Silbling
	}
	return tNode
}

func (trie *LinkedTrie) FindLastMatchNode(key []byte) (*TrieNode, int) {
	node := trie.Root.FirstChild
	var last *TrieNode = nil
	if node == nil {
		return nil, -1
	}
	lastI := 0
	i := 0
	for i = 0; i < len(key); i++ {
		node = trie.FindInSilbling(node, key[i])
		if node == nil {
			break
		}
		last = node
		node = node.FirstChild
	}
	lastI = i
	return last, lastI
}

func (trie *LinkedTrie) AddKey(key []byte, id int) {
	node, lastI := trie.FindLastMatchNode(key)
	if node == nil {
		node = trie.Root
		lastI = 0
	}
	for i := lastI; i < len(key); i++ {
		newNode := NewTrieNode(key[i], -1)
		fNode := node.FirstChild
		var preNode *TrieNode = nil
		for fNode != nil {
			if fNode.Key > newNode.Key {
				break
			}
			preNode = fNode
			fNode = fNode.Silbling
		}
		if preNode == nil {
			newNode.Silbling = node.FirstChild
			node.FirstChild = newNode
		} else {
			newNode.Silbling = fNode
			preNode.Silbling = newNode
		}
		node = newNode
	}
	node.flag = true
	node.ID = id
}
