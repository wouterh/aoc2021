package main

import (
	"fmt"
	"os"

	"github.com/wouterh/aoc2021/internal/input"
)

func findBoundaries(positions []int) (min, max int) {
	min = positions[0]
	max = positions[0]
	for _, v := range positions {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}

func constantFuel(from, to int) int {
	if from > to {
		return from - to
	} else {
		return to - from
	}
}

var growingCosts []int

func initGrowingCosts(max int) {
	growingCosts = make([]int, max)
	growingCosts[0] = 1
	for i := 1; i < max; i++ {
		growingCosts[i] = growingCosts[i-1] + i + 1
	}
}

func growingFuel(from, to int) int {
	if from > to {
		return growingCosts[from-to-1]
	} else if to > from {
		return growingCosts[to-from-1]
	} else {
		return 0
	}
}

func findFuelForAlignment(positions []int, min, max int, cost func(int, int) int) int {
	var minFuel *int

	for i := min; i <= max; i++ {
		fuel := 0
		for _, v := range positions {
			fuel = fuel + cost(v, i)
		}
		if minFuel == nil {
			minFuel = &fuel
		} else if *minFuel > fuel {
			*minFuel = fuel
		}
	}

	if minFuel != nil {
		return *minFuel
	} else {
		return 0
	}
}

func main() {
	positions, err := input.ReadNumbersOnLine(os.Args[1])
	if err != nil {
		panic(err)
	}
	min, max := findBoundaries(positions)

	minFuel := findFuelForAlignment(positions, min, max, constantFuel)
	fmt.Println(minFuel)

	initGrowingCosts(max)
	minFuel = findFuelForAlignment(positions, min, max, growingFuel)
	fmt.Println(minFuel)
}
