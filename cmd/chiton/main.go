package main

import (
	"fmt"
	"os"

	"github.com/wouterh/aoc2021/internal/input"
)

func riskFor(path []*Cell) int {
	// The path is reversed, so the first cell is the destination -> its dist is
	// the risk
	return path[0].dist
}

func main() {
	lines, err := input.ReadStrings(os.Args[1])
	if err != nil {
		panic(err)
	}
	cave, err := NewChitonCave(lines)
	if err != nil {
		panic(err)
	}
	reversedShortest := cave.ShortestPath(Coord{X: 0, Y: 0}, Coord{X: cave.XSize - 1, Y: cave.YSize - 1})
	fmt.Println("Risk for part 1:", riskFor(reversedShortest))
}
