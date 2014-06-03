// Token
package GoRio

type Token struct {
	Text        []byte
	frequency   int
	PrefixToken []*Token
	SubToken    []*Token
}
