package main

import (
	"fmt"
	"os"

	"github.com/wouterh/aoc2021/internal/input"
)

func main() {
	lines, err := input.ReadStrings(os.Args[1])
	if err != nil {
		panic(err)
	}
	m, err := NewMap(lines)
	if err != nil {
		panic(err)
	}

	fmt.Println("No double small")
	paths := m.FindPaths(false)
	paths.Print()
	fmt.Println(len(paths))

	fmt.Println()
	fmt.Println("One double small")
	paths = m.FindPaths(true)
	paths.Print()
	fmt.Println(len(paths))
}
