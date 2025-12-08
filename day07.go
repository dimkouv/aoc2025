package aoc

import (
	"fmt"
)

type grid struct {
	mem  [][]rune
	memo map[[2]int]int
}

func newGrid(inputFile string) *grid {
	lines, err := fileReadLines(inputFile)
	pie(err)

	mem := make([][]rune, 0, len(lines))
	for _, ln := range lines {
		row := make([]rune, 0, len(ln))
		for _, r := range ln {
			row = append(row, r)
		}
		mem = append(mem, row)
	}

	return &grid{mem: mem, memo: make(map[[2]int]int)}
}

func (g *grid) simulate() {
	numSplit := 0

	for i, r := range g.mem[:len(g.mem)-1] {
		for j, v := range r {
			if v == 'S' {
				g.mem[i+1][j] = '|'
			}

			if v == '|' && g.mem[i+1][j] == '^' {
				numSplit++
				if g.mem[i+1][j-1] == '.' {
					g.mem[i+1][j-1] = '|'
				}
				if g.mem[i+1][j+1] == '.' {
					g.mem[i+1][j+1] = '|'
				}
			}

			if v == '|' && g.mem[i+1][j] == '.' {
				g.mem[i+1][j] = '|'
			}
		}
	}
	fmt.Println(numSplit)
}

func (g *grid) simulatePart2() {
	for i, r := range g.mem {
		for j, v := range r {
			if v == 'S' {
				fmt.Println(g.countTimelines(i+1, j))
			}
		}
	}
}

func (g *grid) countTimelines(i, j int) int {
	h := len(g.mem)
	w := len(g.mem[0])

	if i >= h || j < 0 || j >= w {
		return 1
	}

	if v, ok := g.memo[[2]int{i, j}]; ok {
		return v
	}

	v := g.mem[i][j]

	res := 0
	switch v {
	case '.', '|':
		res = g.countTimelines(i+1, j)
	case '^':
		res = g.countTimelines(i, j-1) +
			g.countTimelines(i, j+1)
	default:
		panic("invalid grid char")
	}

	g.memo[[2]int{i, j}] = res
	return res
}

func day7(inputFile string) {
	gr := newGrid(inputFile)
	gr.simulate()
	gr.simulatePart2()
}
