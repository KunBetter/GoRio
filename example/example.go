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
	startT := time.Now()
	//max words
	for i := 0; i < 100; i++ {
		rio.CutWord("中国人民共和国", false)
	}
	tokens := rio.CutWord("中国人民共和国", false)
	TimeConsumed := float64(time.Now().UnixNano()-startT.UnixNano()) / 1e9
	fmt.Printf("TimeConsumed: %f s.\n", TimeConsumed)
	fmt.Printf("tokens:%v.\n", tokens)
}
