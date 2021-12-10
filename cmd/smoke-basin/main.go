package main

import (
	"container/heap"
	"fmt"
	"os"

	"github.com/wouterh/aoc2021/internal/util"
)

func riskSum(lows []int) (sum int) {
	for _, v := range lows {
		sum = sum + v + 1
	}
	return sum
}

func main() {
	m, err := NewMap(os.Args[1])
	if err != nil {
		panic(err)
	}
	lows := m.FindLowPoints()
	rSum := riskSum(lows)
	fmt.Println(rSum)

	lowCoords := m.FindLowPointCoordinates()
	basinSizes := &util.MaxIntHeap{}
	heap.Init(basinSizes)
	for _, l := range lowCoords {
		heap.Push(basinSizes, m.Fill(l))
	}
	product := heap.Pop(basinSizes).(int) * heap.Pop(basinSizes).(int) * heap.Pop(basinSizes).(int)
	fmt.Println(product)
}
