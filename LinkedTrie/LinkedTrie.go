// Trie
package LinkedTrie

import (
	"bufio"
	"io"
	"os"
	"strings"
)

type TrieNode struct {
	Key        byte
	Value      interface{}
	Silbling   *TrieNode
	FirstChild *TrieNode
	flag       bool
}

func NewTrieNode(key byte, value interface{}) *TrieNode {
	node := &TrieNode{
		Key:        key,
		Value:      value,
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
		Root: NewTrieNode(0, nil),
	}
	return trie
}

func (trie *LinkedTrie) AddWord(word string, value interface{}) {
	key := []byte(word)
	trie.AddKey(key, value)
}

func (trie *LinkedTrie) FindWord(word string) bool {
	key := []byte(word)
	return trie.FindKey(key)
}

func (trie *LinkedTrie) CutWord(word string) []string {
	key := []byte(word)
	return trie.CutKey(key, true)
}

func (trie *LinkedTrie) CutKey(key []byte, smart bool) []string {
	words := []string{}
	node := trie.Root
	if node == nil {
		return nil
	}
	i := 0
	path := []byte{}
	for i = 0; i < len(key); i++ {
		node = node.FirstChild
		node = trie.FindInSilbling(node, key[i])
		if node == nil {
			break
		}
		path = append(path, node.Key)
		if node.flag {
			words = append(words, string(path))
			if smart {
				break
			}
		}
	}
	if smart && i < len(key) {
		words = append(words, trie.CutKey(key[i+1:], true)...)
	}
	return words
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

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func (trie *LinkedTrie) LoadDicFiles(fn string, reverse bool) {
	f, err := os.Open(fn)
	defer f.Close()
	if nil == err {
		buff := bufio.NewReader(f)
		for {
			line, err := buff.ReadString('\n')
			if err != nil || io.EOF == err {
				break
			}
			line = strings.TrimRight(line, "\r\n")
			if reverse {
				trie.AddWord(Reverse(line), nil)
			} else {
				trie.AddWord(line, nil)
			}
		}
	}
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

func (trie *LinkedTrie) AddKey(key []byte, value interface{}) {
	node, lastI := trie.FindLastMatchNode(key)
	if node == nil {
		node = trie.Root
		lastI = 0
	}
	for i := lastI; i < len(key); i++ {
		newNode := NewTrieNode(key[i], nil)
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
	node.Value = value
}
