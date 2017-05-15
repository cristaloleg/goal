package graph

import (
	"sort"

	"github.com/cristaloleg/golds/misc"
)

// Edge represents a graph edge
type Edge struct {
	src    int
	dst    int
	weight float64
}

// Kruskal return MST(minimum spanning tree)
func Kruskal(edges []Edge) []Edge {
	sort.SliceStable(edges, func(i, j int) bool {
		return edges[i].weight < edges[j].weight
	})

	size := len(edges)
	uf := misc.NewUnionFind(size)
	var res []Edge

	for _, p := range edges {
		if !uf.IsUnited(p.src, p.dst) {
			uf.Union(p.src, p.dst)
			res = append(res, p)
		}
	}
	return res
}

// KruskalSorted return MST(minimum spanning tree)
// Assumes that edges are already sorted by weight
func KruskalSorted(edges []Edge) []Edge {
	size := len(edges)
	uf := misc.NewUnionFind(size)
	var res []Edge

	for _, p := range edges {
		if !uf.IsUnited(p.src, p.dst) {
			uf.Union(p.src, p.dst)
			res = append(res, p)
		}
	}
	return res
}
