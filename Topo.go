// Topo
package GoRio

type TopoNode struct {
	Key     int
	Childs  map[int]*TopoNode
	Weight  int
	Prefix  []*Token
	PLength []int
	EdgeNum int
	Single  int
}

func NewTopoNode(key int) *TopoNode {
	tn := &TopoNode{
		Key:     key,
		Childs:  make(map[int]*TopoNode),
		Weight:  0,
		Prefix:  []*Token{},
		PLength: []int{},
		EdgeNum: 0,
		Single:  0,
	}
	return tn
}

func (tn *TopoNode) ComputeWeight() {
	tn.Weight = (10000-tn.EdgeNum*10)*10 - tn.Single*200
	pLen := len(tn.PLength)
	for i := 0; i < pLen; i++ {
		tn.Weight += tn.PLength[i] * (3000 - i*30)
	}
}

type Topo struct {
	Root *TopoNode
}

func NewTopo() *Topo {
	t := &Topo{
		Root: NewTopoNode(0),
	}
	return t
}
