package ds

import (
	"testing"
)

func TestTopoOrder(t *testing.T) {
	fm := []int{0, 0, 0, 6, 6, 5, 3, 2, 2, 8, 7, 9, 9, 9, 11}
	to := []int{1, 6, 5, 4, 9, 4, 5, 0, 3, 7, 6, 10, 11, 12, 12}
	edgeSize := len(fm)

	g := NewDigraph()
	for i := 0; i < edgeSize; i++ {
		g.AddEdge(fm[i], to[i])
	}

	topo0 := NewTopo(g, 5)

	t.Log("graph")
	t.Log("topo: ", topo0.GetTopoOrder())

	t.Log("reverse graph")
	rg := g.Reverse()
	dfo1 := NewTopo(rg, 5)
	t.Log("topo: ", dfo1.GetTopoOrder())
}
