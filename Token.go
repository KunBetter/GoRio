// Token
package GoRio

type Token struct {
	Text      []byte
	frequency int
	SubToken  []*Token
	SubPos    []int
}

type Segs struct {
	Text   string
	Offset int
}
