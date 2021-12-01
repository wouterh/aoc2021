package main

import (
	"fmt"
	"os"

	"github.com/wouterh/aoc2021/internal/input"
)

func countIncreases(depths []int64) int {
	counter := int(0)
	for i := 1; i < len(depths); i++ {
		if depths[i] > depths[i-1] {
			counter++
		}
	}
	return counter
}

func countIncreasesSliding(depths []int64) int {
	counter := int(0)
	for i := 3; i < len(depths); i++ {
		a := depths[i-3] + depths[i-2] + depths[i-1]
		b := depths[i-2] + depths[i-1] + depths[i]
		if b > a {
			counter++
		}
	}
	return counter
}

func main() {
	depths, err := input.ReadNumbers(os.Args[1])
	if err != nil {
		panic("Could not fetch input")
	}
	count := countIncreases(depths)
	fmt.Println(count)
	countSliding := countIncreasesSliding(depths)
	fmt.Println(countSliding)
}
