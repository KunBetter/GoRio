// GoRio
package GoRio

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type GoRio struct {
	DicFilesName []string
	ForwardTrie  *SliceTrie
	Tokens       []*Token
}

func GoRior() *GoRio {
	rio := &GoRio{
		DicFilesName: []string{},
		ForwardTrie:  NewSliceTrie(),
		Tokens:       []*Token{},
	}
	return rio
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
	for _, token := range rio.Tokens {
		rio.ForwardTrie.AddToken(token)
	}
	//PreProcess
	for _, token := range rio.Tokens {
		rio.ForwardTrie.AddSubToken(token)
	}
}

func (rio *GoRio) AddWord(word []byte) {
	tk := &Token{word, 0, nil, nil}
	rio.ForwardTrie.AddToken(tk)
	rio.ForwardTrie.AddSubToken(tk)
}

func (rio *GoRio) CutWord(word []byte, smart bool) []Segs {
	tk := []Segs{}
	offset := 0
	if smart {
		tokens := rio.ForwardTrie.SmartSegments(word)
		for _, v := range tokens {
			tk = append(tk, Segs{string(v.Text), offset})
			offset += len(v.Text)
		}
	} else {
		tokens := rio.ForwardTrie.SmartSegments(word)
		for _, v := range tokens {
			if v.SubToken != nil {
				for i, st := range v.SubToken {
					tk = append(tk, Segs{string(st.Text), v.SubPos[i] + offset})
				}
			} else {
				tk = append(tk, Segs{string(v.Text), offset})
			}
			offset += len(v.Text)
		}
	}
	return tk
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
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
			rio.Tokens = append(rio.Tokens,
				&Token{[]byte(line), 0, nil, nil})
		}
	}
}
