GoRio
==========
go concurrent word segmentation plug-in

Installation
-----
go get github.com/KunBetter/GoRio  
$GOPATH/bin/GoRio

Usage
-----
```go
import (
	"github.com/KunBetter/GoRio"
)

func main() {
	rio := GoRio.GoRior()
	rio.SetDicFilesName([]string{"../dic/words.txt"})
	rio.Start()
	rio.FindWord("中华人民共和国")
}
```