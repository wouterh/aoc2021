package main

import (
	"container/heap"
)

// An MinVertexHeap is a min-heap of Vertices.
type MinVertexHeap []*Cell

func (h MinVertexHeap) Len() int           { return len(h) }
func (h MinVertexHeap) Less(i, j int) bool { return h[i].dist < h[j].dist }
func (h MinVertexHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinVertexHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*Cell))
}

func (h *MinVertexHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

var _ heap.Interface = &MinVertexHeap{}
