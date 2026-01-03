package aoc

import (
	"fmt"
	"strings"
)

type N struct {
	id    string
	neibs []*N
}

type G struct {
	nodes map[string]*N
}

func day11(inputFile string) {
	g := parseG(inputFile)
	day11part1(g)
	day11part2(g)
}

func day11part1(g *G) {
	fmt.Println(g.dfs("you", "out", map[string]struct{}{}))
}

func parseG(inputFile string) *G {
	lines, err := fileReadLines(inputFile)
	pie(err)

	nodes := make(map[string]*N)
	for _, ln := range lines {
		parts := strings.Split(ln, ": ")
		fromID := parts[0]
		toIDs := strings.Split(parts[1], " ")

		from, ok := nodes[fromID]
		if !ok {
			from = &N{neibs: make([]*N, 0), id: fromID}
			nodes[fromID] = from
		}

		for _, toID := range toIDs {
			to, ok := nodes[toID]
			if !ok {
				to = &N{neibs: make([]*N, 0), id: toID}
				nodes[toID] = to
			}
			from.neibs = append(from.neibs, to)
		}
	}

	return &G{nodes: nodes}
}

func day11part2(g *G) {
	nodesSeenFromFFT := g.exploreNodes("fft")
	nodesSeenFromDAC := g.exploreNodes("dac")

	// fmt.Println("dac", len(nodesSeenFromDAC)) // 131
	// fmt.Println("fft", len(nodesSeenFromFFT)) // 370 (fft is first)
	// svr -> fft -> dac -> out

	c1 := g.dfs("svr", "fft", nodesSeenFromFFT)
	c2 := g.dfs("fft", "dac", nodesSeenFromDAC)
	c3 := g.dfs("dac", "out", map[string]struct{}{})

	fmt.Println(c1 * c2 * c3)
}

func (g *G) dfs(a, b string, filter map[string]struct{}) int64 {
	cnt := int64(0)
	queue := []string{a}
	for {
		next := queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		neibs := g.nodes[next].neibs
		for _, n := range neibs {
			if n.id == b {
				cnt++
			} else {
				if _, filtered := filter[n.id]; !filtered {
					queue = append(queue, n.id)
				}
			}
		}
		if len(queue) == 0 {
			break
		}
	}

	return cnt
}

func (g *G) exploreNodes(startNodeID string) map[string]struct{} {
	queue := []string{startNodeID}
	seen := make(map[string]struct{}, 0)

	for {
		next := queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		neibs := g.nodes[next].neibs
		for _, n := range neibs {
			if _, ok := seen[n.id]; !ok {
				queue = append(queue, n.id)
				seen[n.id] = struct{}{}
			}
		}

		if len(queue) == 0 {
			break
		}
	}

	return seen
}
