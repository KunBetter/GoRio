// SliceTrie
package GoRio

import ()

type SliceTrieNode struct {
	Key    byte
	Childs []*SliceTrieNode
	Flag   bool
	Token  *Token
}

func NewSliceTrieNode(key byte) *SliceTrieNode {
	node := &SliceTrieNode{
		Key:    key,
		Childs: []*SliceTrieNode{},
		Flag:   false,
		Token:  nil,
	}
	return node
}

//***************Trie****************

type SliceTrie struct {
	Root      *SliceTrieNode
	MaxWeight *TopoNode
}

//***************PUBLIC****************

func NewSliceTrie() *SliceTrie {
	trie := &SliceTrie{
		Root:      NewSliceTrieNode(0),
		MaxWeight: NewTopoNode(0),
	}
	return trie
}

func (trie *SliceTrie) SmartTopo(tpn *TopoNode, key []byte) {
	if len(key) <= 0 {
		tpn.ComputeWeight()
		if tpn.Weight > trie.MaxWeight.Weight {
			trie.MaxWeight = tpn
		}
		return
	}
	curPrefixToken := trie.PrefixToken(key)
	tLen := len(curPrefixToken)
	if tLen <= 0 {
		aLen := 1
		if key[0] >= 0x80 {
			aLen = 3
		}
		ntpn := NewTopoNode(aLen)
		token := Token{key[0:aLen], 0, nil, nil}
		ntpn.Prefix = append(tpn.Prefix, &token)
		ntpn.EdgeNum = tpn.EdgeNum + 1
		ntpn.Single = tpn.Single + 1
		tpn.Childs[aLen] = ntpn
		trie.SmartTopo(ntpn, key[aLen:])
	}
	for i := 0; i < tLen; i++ {
		cLen := len(curPrefixToken[i].Text)
		ntpn := NewTopoNode(cLen)
		ntpn.Prefix = append(tpn.Prefix, curPrefixToken[i])
		ntpn.EdgeNum = tpn.EdgeNum + 1
		tpn.Childs[cLen] = ntpn
		trie.SmartTopo(ntpn, key[cLen:])
	}
}

func (trie *SliceTrie) SmartSegments(key []byte) []*Token {
	topo := NewTopo()
	trie.SmartTopo(topo.Root, key)
	return trie.MaxWeight.Prefix
}

func (trie *SliceTrie) PrefixToken(key []byte) []*Token {
	if len(key) <= 1 {
		return []*Token{}
	}
	token := []*Token{}
	node := trie.Root
	if node == nil {
		return []*Token{}
	}
	for i := 0; i < len(key); i++ {
		childs := node.Childs
		pos, found := trie.FindInChilds(childs, key[i])
		if !found {
			break
		}
		node = childs[pos]
		if node.Flag {
			token = append(token, node.Token)
		}
	}
	return token
}

func (trie *SliceTrie) AddSubToken(token *Token) bool {
	key := []byte(token.Text)
	node := trie.Root
	if node == nil {
		return false
	}
	for i := 0; i < len(key); i++ {
		childs := node.Childs
		pos, found := trie.FindInChilds(childs, key[i])
		if !found {
			break
		}
		node = childs[pos]
	}
	if node.Flag {
		node.Token.SubToken, node.Token.SubPos = trie.SubToken(key)
		return true
	}
	return false
}

func (trie *SliceTrie) SubToken(key []byte) ([]*Token, []int) {
	token := []*Token{}
	pos := []int{}
	for i := 0; i < len(key); i++ {
		ptks := trie.PrefixToken(key[i:])
		if len(ptks) > 0 {
			for _, v := range ptks {
				token = append(token, v)
				pos = append(pos, i)
			}
		}
	}
	return token, pos
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

func (trie *SliceTrie) AddToken(token *Token) {
	key := []byte(token.Text)
	node, lastI := trie.FindLastMatchNode(key)
	if node == nil {
		node = trie.Root
		lastI = 0
	}
	for i := lastI; i < len(key); i++ {
		newNode := NewSliceTrieNode(key[i])
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
	node.Token = token
}
