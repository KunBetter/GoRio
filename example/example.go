// GoRio
package main

import (
	"fmt"
	"github.com/KunBetter/GoRio"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	rio := GoRio.GoRior()
	rio.SetDicFilesName([]string{"../dic/words.txt"})
	rio.Start()
	ws := []byte("中华人民共和国")
	startT := time.Now()
	for i := 0; i < 10000; i++ {
		rio.CutWord(ws, false)
	}
	tokens := rio.CutWord(ws, false)
	TimeConsumed := float64(time.Now().UnixNano()-startT.UnixNano()) / 1e9
	fmt.Printf("TimeConsumed: %f s.\n", TimeConsumed)
	fmt.Printf("tokens:%v.\n", rio.Tokens2String(tokens))
	startT = time.Now()
	for i := 0; i < 10000; i++ {
		rio.CutWord(ws, true)
	}
	TimeConsumed = float64(time.Now().UnixNano()-startT.UnixNano()) / 1e9
	fmt.Printf("TimeConsumed: %f s.\n", TimeConsumed)
	tokens = rio.CutWord(ws, true)
	fmt.Printf("tokens:%v.\n", rio.Tokens2String(tokens))
}
