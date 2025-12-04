package aoc

import "fmt"

type loc [2]int

var adjacentLocs = [8]loc{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

func day4(inputFile string) {
	lines, err := fileReadLines(inputFile)
	pie(err)

	rolls := make(map[loc]bool)
	for i, line := range lines {
		for j, v := range line {
			if v == '@' {
				rolls[loc{i, j}] = true
			}
		}
	}

	totalRemoved := 0
	for {
		newRollsState := make(map[loc]bool)
		cntRemoved := 0
		for roll := range rolls {
			numAdj := 0
			for _, adj := range adjacentLocs {
				adjRoll := loc{roll[0] + adj[0], roll[1] + adj[1]}
				adjRollExists := rolls[adjRoll]
				if adjRollExists {
					numAdj++
				}
			}

			if numAdj < 4 {
				cntRemoved++
			} else {
				newRollsState[roll] = true // keep the roll in the next state
			}
		}

		if totalRemoved == 0 {
			fmt.Println(cntRemoved) // print part 1 result
		}

		rolls = newRollsState
		totalRemoved += cntRemoved
		if cntRemoved == 0 {
			break
		}
	}

	fmt.Println(totalRemoved)
}
