// GoRio
package GoRio

import (
	"bufio"
	"fmt"
	"github.com/KunBetter/GoRio/LinkedTrie"
	"io"
	"os"
	"strings"
)

type WordSeg struct {
	word  string
	resp  chan map[int]string
	smart bool
}

type GoRio struct {
	DicFilesName []string
	ForwardTrie  *LinkedTrie.LinkedTrie
	ReverseTrie  *LinkedTrie.LinkedTrie
	req          chan *WordSeg
}

func GoRior() *GoRio {
	rio := &GoRio{
		DicFilesName: []string{},
		ForwardTrie:  LinkedTrie.NewLinkedTrie(),
		ReverseTrie:  LinkedTrie.NewLinkedTrie(),
		req:          make(chan *WordSeg, 10),
	}
	go rio.process()
	return rio
}

func (rio *GoRio) process() {
	for {
		select {
		case ws := <-rio.req:
			go func(ws *WordSeg) {
				if !ws.smart {
					segs := rio.MaxSegments(ws.word)
					ws.resp <- segs
				}
			}(ws)
		}
	}
}

func (rio *GoRio) SetDicFilesName(fn []string) {
	rio.DicFilesName = fn
}

func (rio *GoRio) Start() {
	if len(rio.DicFilesName) == 0 {
		fmt.Println("no dic files.")
		return
	}
	id := 1
	for _, fn := range rio.DicFilesName {
		id = rio.LoadDicFiles(fn, id)
	}
}

func (rio *GoRio) CutWord(word string, smart bool) map[int]string {
	ws := &WordSeg{
		word:  word,
		resp:  make(chan map[int]string),
		smart: smart,
	}
	rio.req <- ws
	return <-ws.resp
}

func (rio *GoRio) MaxSegments(word string) map[int]string {
	segs := make(map[int]string)
	fWords, fIDs := rio.ForwardTrie.MaxSegments(word)
	for i := 0; i < len(fIDs); i++ {
		segs[fIDs[i]] = fWords[i]
	}
	rWords, rIDs := rio.ReverseTrie.MaxSegments(Reverse(word))
	for i := 0; i < len(fIDs); i++ {
		_, ok := segs[rIDs[i]]
		if !ok {
			segs[rIDs[i]] = Reverse(rWords[i])
		}
	}
	return segs
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func (rio *GoRio) LoadDicFiles(fn string, id int) int {
	tid := id
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
			rio.ForwardTrie.AddWord(line, tid)
			rio.ReverseTrie.AddWord(Reverse(line), tid)
			tid++
		}
	}
	return tid
}
