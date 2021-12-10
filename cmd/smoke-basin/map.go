package main

import (
	"container/list"
	"strconv"

	"github.com/wouterh/aoc2021/internal/input"
)

type Map [][]int
type Coord struct {
	x int
	y int
}

func (c Coord) Neighbours(width, height int) (neighbours []Coord) {
	if c.x > 0 {
		neighbours = append(neighbours, Coord{x: c.x - 1, y: c.y})
	}
	if c.x < width-1 {
		neighbours = append(neighbours, Coord{x: c.x + 1, y: c.y})
	}
	if c.y > 0 {
		neighbours = append(neighbours, Coord{x: c.x, y: c.y - 1})
	}
	if c.y < height-1 {
		neighbours = append(neighbours, Coord{x: c.x, y: c.y + 1})
	}
	return neighbours
}

func NewMap(filename string) (*Map, error) {
	lines, err := input.ReadStrings(filename)
	if err != nil {
		return nil, err
	}
	var m Map = make(Map, len(lines))
	for i, l := range lines {
		var row []int = make([]int, len(l))
		for j, r := range l {
			n, err := strconv.Atoi(string(r))
			if err != nil {
				return nil, err
			}
			row[j] = n
		}
		m[i] = row

	}
	return &m, nil
}

func (m *Map) Dimensions() (int, int) {
	return len(*m), len((*m)[0])
}

func (m *Map) FindLowPoints() []int {
	coords := m.FindLowPointCoordinates()
	var lowPoints []int = make([]int, len(coords))
	for i, c := range coords {
		lowPoints[i] = (*m)[c.x][c.y]
	}
	return lowPoints
}

func (m *Map) FindLowPointCoordinates() (coords []Coord) {
	rows, cols := m.Dimensions()
	for i, row := range *m {
		for j, v := range row {
			if j < cols-1 && v >= row[j+1] {
				continue
			} else if j > 0 && v >= row[j-1] {
				continue
			} else if i < rows-1 && v >= (*m)[i+1][j] {
				continue
			} else if i > 0 && v >= (*m)[i-1][j] {
				continue
			} else {
				coords = append(coords, Coord{
					x: i,
					y: j,
				})
			}
		}
	}
	return coords
}

func (m *Map) FillTo(startingPoint Coord, level int) int {
	width, height := m.Dimensions()
	queue := list.New()
	queue.PushBack(startingPoint)
	done := 0
	for next := queue.Front(); next != nil; next = queue.Front() {
		queue.Remove(next)

		c := next.Value.(Coord)
		if (*m)[c.x][c.y] < level {
			done++
			(*m)[c.x][c.y] = level
			for _, c := range c.Neighbours(width, height) {
				queue.PushBack(c)
			}
		}
	}
	return done
}

func (m *Map) Fill(startingPoint Coord) int {
	var done int = 0
	for i := (*m)[startingPoint.x][startingPoint.y] + 1; i <= 9; i++ {
		done = m.FillTo(startingPoint, i)
	}
	return done
}
