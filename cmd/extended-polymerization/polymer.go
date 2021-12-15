package main

import (
	"fmt"
	"math"
	"strings"
)

type Polymerization struct {
	elementCounts  map[rune]int
	pairCounts     map[string]int
	insertionRules map[string]rune
}

func NewPolymerization(lines []string) *Polymerization {
	p := &Polymerization{
		elementCounts:  make(map[rune]int),
		insertionRules: make(map[string]rune, len(lines)-2),
		pairCounts:     make(map[string]int),
	}
	template := lines[0]
	for i, r := range template {
		p.elementCounts[r]++
		if i < len(template)-1 {
			p.pairCounts[template[i:i+2]]++
		}
	}
	for _, line := range lines[2:] {
		parts := strings.Split(line, " -> ")
		p.insertionRules[parts[0]] = rune(parts[1][0])
	}
	return p
}

func (p *Polymerization) Print() {
	for k, v := range p.elementCounts {
		fmt.Printf("%s: %d ", string(k), v)
	}
	fmt.Println("")
	for k, v := range p.pairCounts {
		fmt.Printf("%s: %d ", k, v)
	}
	fmt.Println()
}

func (p *Polymerization) Step() {
	oldPairCounts := make(map[string]int, len(p.pairCounts))
	for k, v := range p.pairCounts {
		oldPairCounts[k] = v
	}
	// For each pair, we insert an element and create two new pairs
	for pair, count := range oldPairCounts {
		if insert, ok := p.insertionRules[pair]; ok {
			// There is a rule
			p.elementCounts[insert] += count
			p.pairCounts[pair] -= count
			np1, np2 := string([]rune{rune(pair[0]), insert}), string([]rune{insert, rune(pair[1])})
			p.pairCounts[np1] += count
			p.pairCounts[np2] += count
		}
	}
}

func (p *Polymerization) MinimaxCounts() (min, max int) {
	min = math.MaxInt64
	for _, v := range p.elementCounts {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return min, max
}
