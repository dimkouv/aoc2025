package aoc

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type loc3d struct {
	x, y, z int64
}

func (l loc3d) distance(o loc3d) float64 {
	dx := float64(l.x - o.x)
	dy := float64(l.y - o.y)
	dz := float64(l.z - o.z)
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

type circuit struct {
	items map[loc3d]struct{}
}

func newCircuit() *circuit {
	return &circuit{items: make(map[loc3d]struct{})}
}

func (c *circuit) contains(i loc3d) bool {
	_, ok := c.items[i]
	return ok
}

func (c *circuit) add(i loc3d) {
	c.items[i] = struct{}{}
}

func (c *circuit) clear() {
	c.items = nil
}

func day8(inputFile string) {
	lines, err := fileReadLines(inputFile)
	pie(err)

	circuits := make([]circuit, 0)
	points := make([]loc3d, 0)

	for _, l := range lines {
		parts := strings.Split(l, ",")
		x, err := strconv.ParseInt(parts[0], 10, 64)
		pie(err)
		y, err := strconv.ParseInt(parts[1], 10, 64)
		pie(err)
		z, err := strconv.ParseInt(parts[2], 10, 64)
		pie(err)
		l := loc3d{x, y, z}
		points = append(points, l)
	}

	distances := make(map[float64][2]loc3d)
	for i, p1 := range points {
		for _, p2 := range points[i+1:] {
			d := p1.distance(p2)
			distances[d] = [2]loc3d{p1, p2}
		}
	}

	keys := make([]float64, 0, len(distances))
	for k := range distances {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	idx := -1
	lastVal := int64(0)
	for _, k := range keys {
		idx++
		if idx == 1000 {
			cntCircuitSizes := make([]int, 0)
			for _, circ := range circuits {
				cntCircuitSizes = append(cntCircuitSizes, len(circ.items))
			}
			sort.Slice(cntCircuitSizes, func(i, j int) bool { return cntCircuitSizes[i] > cntCircuitSizes[j] })
			fmt.Println(cntCircuitSizes[0] * cntCircuitSizes[1] * cntCircuitSizes[2])
		}

		p1 := distances[k][0]
		p2 := distances[k][1]

		c1, circP1 := getCircuit(circuits, p1)
		c2, circP2 := getCircuit(circuits, p2)
		if c1 == c2 && c1 != -1 && c2 != -1 {
			continue
		}

		if circP1 != nil && circP2 != nil {
			for k := range circP2.items {
				circP1.add(k)
			}
			circuits = append(circuits[:c2], circuits[c2+1:]...)
		} else if circP1 != nil {
			circP1.add(p2)
		} else if circP2 != nil {
			circP2.add(p1)
		} else {
			newCirc := newCircuit()
			newCirc.add(p1)
			newCirc.add(p2)
			circuits = append(circuits, *newCirc)
		}

		lastVal = p1.x * p2.x
	}
	fmt.Print(lastVal)
}

func getCircuit(circuits []circuit, p loc3d) (int, *circuit) {
	for i, circ := range circuits {
		if circ.contains(p) {
			return i, &circ
		}
	}
	return -1, nil
}
