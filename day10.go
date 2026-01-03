package aoc

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type s struct {
	targetState int
	mutations   []int
	bitVoltages []int
}

func day10(inputFile string) {
	lines, err := fileReadLines(inputFile)
	pie(err)

	s := 0
	for _, line := range lines {
		lineState := parseLine(line)
		s += reachTargetState(lineState)
	}
	fmt.Println(s)
}

func reachTargetState(st s) int {
	iters := 0
	originStates := []int{0}
	for {
		iters++
		newOriginStates := []int{}
		for _, mu := range st.mutations {
			for _, ost := range originStates {
				res := ost ^ mu
				if res == st.targetState {
					return iters
				}
				newOriginStates = append(newOriginStates, res)
			}
		}
		originStates = newOriginStates
	}
}

func parseLine(input string) s {
	var result s

	re := regexp.MustCompile(`\[(.*?)\]\s*(.*)\s*\{(.*?)\}`)
	matches := re.FindStringSubmatch(input)
	if len(matches) < 4 {
		pie(fmt.Errorf("invalid format"))
	}

	rawBits := matches[1]
	mutationPart := matches[2]
	voltagePart := matches[3]

	width := len(rawBits)
	for i, char := range rawBits {
		if char == '#' {
			result.targetState |= (1 << (width - 1 - i))
		}
	}

	mutGroupRe := regexp.MustCompile(`\((.*?)\)`)
	mutGroups := mutGroupRe.FindAllStringSubmatch(mutationPart, -1)
	for _, group := range mutGroups {
		var m int
		indices := strings.Split(group[1], ",")
		for _, idxStr := range indices {
			idx, _ := strconv.Atoi(strings.TrimSpace(idxStr))
			if idx < width {
				m |= (1 << (width - 1 - idx))
			}
		}
		result.mutations = append(result.mutations, m)
	}

	for _, vStr := range strings.Split(voltagePart, ",") {
		v, _ := strconv.Atoi(strings.TrimSpace(vStr))
		result.bitVoltages = append(result.bitVoltages, v)
	}

	return result
}
