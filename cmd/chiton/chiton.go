package main

import (
	"container/heap"
	"fmt"
	"math"
	"strconv"
)

type Cell struct {
	coords Coord
	value  int
	dist   int
	prev   *Cell
	done   bool
}

type ChitonCave struct {
	cells []*Cell
	XSize int
	YSize int
}

type Coord struct {
	X, Y int
}

func NewChitonCave(lines []string) (*ChitonCave, error) {
	xSize := len(lines[0])
	ySize := len(lines)
	cave := &ChitonCave{
		cells: make([]*Cell, xSize*ySize),
		XSize: xSize,
		YSize: ySize,
	}
	for i, line := range lines {
		for j, v := range line {
			iv, err := strconv.Atoi(string(v))
			if err != nil {
				return nil, err
			}
			cave.cells[j+xSize*i] = &Cell{
				coords: Coord{X: j, Y: i},
				value:  iv,
			}
		}
	}
	return cave, nil
}

func (c *ChitonCave) Get(x, y int) *Cell {
	return c.cells[x+c.XSize*y]
}

func (cc *ChitonCave) validCoords(c Coord) bool {
	x, y := c.X, c.Y
	return x >= 0 && x < cc.XSize && y >= 0 && y < cc.YSize
}

func (cc *ChitonCave) neighbourCoords(x, y int) (coords []Coord) {
	candidates := []Coord{
		{x, y - 1},
		{x - 1, y},
		{x + 1, y},
		{x, y + 1},
	}
	filtered := candidates[:0]
	for _, c := range candidates {
		if cc.validCoords(c) {
			filtered = append(filtered, c)
		}
	}
	return filtered
}

// We need the shortest path, let's use Dijkstra:
// https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm
func (cc *ChitonCave) ShortestPath(from, to Coord) []*Cell {
	// Initialize
	Q := make(MinVertexHeap, 0, cc.XSize*cc.YSize)
	heap.Init(&Q)
	for i := 0; i < cc.XSize; i++ {
		for j := 0; j < cc.YSize; j++ {
			cell := cc.Get(i, j)
			if cell.coords == from {
				cell.dist = 0
			} else {
				cell.dist = math.MaxInt32
			}
			heap.Push(&Q, cell)
		}
	}

	var u *Cell

	// Loop over Q
	for len(Q) > 0 {
		u = heap.Pop(&Q).(*Cell)
		u.done = true
		fmt.Print("Now at (", u.coords.X, ",", u.coords.Y, ") with distance ", u.dist)
		if u.prev != nil {
			fmt.Println(" | prev: (", u.prev.coords.X, ",", u.prev.coords.Y, ")")
		} else {
			fmt.Println("")
		}
		if u.coords == to {
			break
		}
		for _, n := range cc.neighbourCoords(u.coords.X, u.coords.Y) {
			v := cc.Get(n.X, n.Y)
			if !v.done {
				// still in Q
				alt := u.dist + v.value
				if alt < v.dist {
					v.dist = alt
					v.prev = u
					// restore the heap invariants
					heap.Init(&Q)
				}
			}
		}
	}

	// Construct the (reversed) path, u is now the destination
	shortest := make([]*Cell, 0)
	for u != nil {
		shortest = append(shortest, u)
		u = u.prev
	}
	return shortest
}
