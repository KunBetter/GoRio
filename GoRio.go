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
	tk := &Token{Text: word, frequency: 0}
	rio.ForwardTrie.AddToken(tk)
	rio.ForwardTrie.AddSubToken(tk)
}

func (rio *GoRio) Tokens2String(tokens []*Token) []string {
	ws := []string{}
	for _, tk := range tokens {
		ws = append(ws, string(tk.Text))
	}
	return ws
}

func (rio *GoRio) CutWord(word []byte, smart bool) []*Token {
	tk := []*Token{}
	tokens := rio.ForwardTrie.SmartSegments(word)
	for _, v := range tokens {
		if smart {
			tk = append(tk, v)
		} else {
			if v.PrefixToken != nil {
				tk = append(tk, v.PrefixToken...)
				tk = append(tk, v.SubToken...)
			} else {
				tk = append(tk, v)
			}
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
				&Token{Text: []byte(line), frequency: 0})
		}
	}
}
