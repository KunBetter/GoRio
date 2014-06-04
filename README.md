GoRio
==========
go word segmentation plug-in

Installation
-----
go get github.com/KunBetter/GoRio  
$GOPATH/bin/GoRio

Smart Word Segmentation
-----
1.Segmentation result set is as small as possible.  
2.Result set is as little as single words.

Usage
-----
```go
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
	ws := []byte("广东联通将联合百度推出跨界互联网金融理财产品")
	//max segments
	tokens := rio.CutWord(ws, false)
	fmt.Printf("max tokens:%v.\n", tokens)
	//smart segments
	tokens = rio.CutWord(ws, true)
	fmt.Printf("smart tokens:%v.\n", tokens)
}
```