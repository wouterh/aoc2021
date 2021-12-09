package main

import (
	"fmt"
	"os"
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
}
