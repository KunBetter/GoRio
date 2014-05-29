// GoRio
package GoRio

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type WordReq struct {
	Word  string
	Resp  chan []Token
	Smart bool
}

type GoRio struct {
	DicFilesName []string
	ForwardTrie  *SliceTrie
	Req          chan *WordReq
	WordCount    int
}

func GoRior() *GoRio {
	rio := &GoRio{
		DicFilesName: []string{},
		ForwardTrie:  NewSliceTrie(),
		Req:          make(chan *WordReq, 100),
		WordCount:    1,
	}
	go rio.Process()
	return rio
}

func (rio *GoRio) Process() {
	for {
		select {
		case ws := <-rio.Req:
			go func(ws *WordReq) {
				if !ws.Smart {
					tokens := rio.MaxSegments(ws.Word)
					ws.Resp <- tokens
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
	for _, fn := range rio.DicFilesName {
		rio.LoadDicFiles(fn)
	}
}

func (rio *GoRio) CutWord(word string, smart bool) []Token {
	ws := &WordReq{
		Word:  word,
		Resp:  make(chan []Token),
		Smart: smart,
	}
	rio.Req <- ws
	return <-ws.Resp
}

func (rio *GoRio) MaxSegments(word string) []Token {
	return rio.ForwardTrie.MaxSegments(word)
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func (rio *GoRio) AddWork(word string) {
	rio.ForwardTrie.AddWord(word, rio.WordCount)
	rio.WordCount++
}

func (rio *GoRio) LoadDicFiles(fn string) {
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
			rio.AddWork(line)
		}
	}
}
