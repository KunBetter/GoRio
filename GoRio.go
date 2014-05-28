// GoRio
package GoRio

import (
	"fmt"
	"github.com/KunBetter/GoRio/LinkedTrie"
	"time"
)

type GoRio struct {
	DicFilesName []string
	ForwardTrie  *LinkedTrie.LinkedTrie
	ReverseTrie  *LinkedTrie.LinkedTrie
}

func GoRior() *GoRio {
	rio := &GoRio{
		DicFilesName: []string{},
		ForwardTrie:  LinkedTrie.NewLinkedTrie(),
		ReverseTrie:  LinkedTrie.NewLinkedTrie(),
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
		rio.ForwardTrie.LoadDicFiles(fn, false)
		rio.ReverseTrie.LoadDicFiles(fn, true)
	}
}

func (rio *GoRio) FindWord(word string) {
	start := time.Now()
	fmt.Println(rio.ForwardTrie.CutWord(word))
	fmt.Println(rio.ReverseTrie.CutWord(LinkedTrie.Reverse(word)))
	fmt.Println(time.Now().UnixNano() - start.UnixNano())
}
