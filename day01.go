package aoc

import (
	"fmt"
	"strconv"
)

type dial struct {
	pos int64
}

func (d *dial) left(v int64) {
	d.pos = ((d.pos-v)%100 + 100) % 100
}

func (d *dial) right(v int64) {
	d.pos = (d.pos + v) % 100
}

func day1(inputFile string) {
	d := &dial{pos: 50}

	lines, err := fileReadLines(inputFile)
	if err != nil {
		panic(err)
	}

	cntZeroes := 0
	cntZeroesPart2 := int64(0)

	for _, line := range lines {
		op := line[0]
		v, err := strconv.ParseInt(line[1:], 10, 64)
		if err != nil {
			panic(err)
		}

		switch op {
		case 'L':
			cntZeroesPart2 += wraps(d.pos, -v)
			d.left(v)
			fmt.Printf("%s: Left %d -> %d (p2: %d)\n", line, v, d.pos, cntZeroesPart2)
		case 'R':
			cntZeroesPart2 += wraps(d.pos, v)
			d.right(v)
			fmt.Printf("%s: Right %d -> %d (p2: %d)\n", line, v, d.pos, cntZeroesPart2)
		default:
			panic("invalid op")
		}

		if d.pos == 0 {
			cntZeroes++
		}
	}

	fmt.Println("total zeroes: ", cntZeroes)
	fmt.Println("total zeroes part2: ", cntZeroesPart2)
}

func wraps(a, b int64) int64 {
	x := a + b

	if x > 0 {
		return x / 100
	} else if x == 0 {
		return 1
	} else {
		if a == 0 {
			return (-x / 100)
		}
		return 1 + (-x / 100)
	}
}
