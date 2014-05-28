// GoRio
package main

import (
	"fmt"
	"github.com/KunBetter/GoRio"
	"time"
)

func main() {
	rio := GoRio.GoRior()
	rio.SetDicFilesName([]string{"../dic/words.txt"})
	rio.Start()
	startT := time.Now()
	//max words
	segs := rio.CutWord("中华人民共和国", false)
	fmt.Printf("time:%d nano.\n", time.Now().UnixNano()-startT.UnixNano())
	fmt.Printf("segs:%v.\n", segs)
}
