package main

import (
	"fmt"
	"strings"
)

const (
	Start = "start"
	End   = "end"
)

type Cave struct {
	Name        string
	Connections []*Cave
}

func (c Cave) IsStart() bool {
	return c.Name == Start
}

func (c Cave) IsEnd() bool {
	return c.Name == End
}

func (c Cave) IsSmall() bool {
	return !c.IsLarge() && !c.IsEnd() && !c.IsStart()
}

func (c Cave) IsLarge() bool {
	f := []byte(c.Name)[0]
	return f < 96
}

func (m *Map) AddConnection(line string) {
	parts := strings.Split(line, "-")
	var from, to *Cave
	var ok bool
	if from, ok = m.Caves[parts[0]]; !ok {
		from = &Cave{
			Name: parts[0],
		}
		m.Caves[from.Name] = from
	}
	if to, ok = m.Caves[parts[1]]; !ok {
		to = &Cave{
			Name: parts[1],
		}
		m.Caves[to.Name] = to
	}
	from.Connections = append(from.Connections, to)
	to.Connections = append(to.Connections, from)
}

type Map struct {
	Caves map[string]*Cave
}

func NewMap(lines []string) (*Map, error) {
	m := &Map{
		Caves: make(map[string]*Cave),
	}
	for _, line := range lines {
		m.AddConnection(line)
	}

	return m, nil
}

type Path []*Cave
type PathList []Path

func (pl PathList) Print() {
	for _, p := range pl {
		fmt.Println(strings.Join(p.Names(), ","))
	}
}

func (p Path) Names() []string {
	var names []string = make([]string, len(p))
	for i, c := range p {
		names[i] = c.Name
	}
	return names
}

func (p Path) Contains(cave *Cave) bool {
	for _, c := range p {
		if cave == c {
			return true
		}
	}
	return false
}

func (p Path) Last() *Cave {
	return p[len(p)-1]
}

func (p Path) Expand(tail *Cave) Path {
	n := make(Path, len(p)+1)
	copy(n, p)
	n[len(p)] = tail
	return n
}

type empty struct{}

func (p Path) HasDoubleSmall() bool {
	var set map[*Cave]empty = map[*Cave]empty{}
	for _, c := range p {
		if c.IsSmall() {
			if _, ok := set[c]; ok {
				return true
			} else {
				set[c] = empty{}
			}
		}
	}
	return false
}

func (m *Map) ExpandPaths(unfinished PathList, allowOneDouble bool) (expanded, finished PathList) {
	for _, p := range unfinished {
		last := p.Last()
		for _, c := range last.Connections {
			if c.IsEnd() {
				finished = append(finished, p.Expand(c))
			} else if c.IsLarge() || (c.IsSmall() && (!p.Contains(c) || (allowOneDouble && !p.HasDoubleSmall()))) {
				expanded = append(expanded, p.Expand(c))
			}
		}
	}
	return expanded, finished
}

func (m *Map) FindPaths(allowOneDouble bool) PathList {
	finished := PathList{}
	unfinished := PathList{{(*m).Caves[Start]}}
	for len(unfinished) > 0 {
		var newFinished []Path
		unfinished, newFinished = m.ExpandPaths(unfinished, allowOneDouble)
		finished = append(finished, newFinished...)
	}
	return finished
}
