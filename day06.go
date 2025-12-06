package aoc

import (
	"fmt"
	"strconv"
	"strings"
)

func day6(inputFile string) {
	lines, err := fileReadLines(inputFile)
	pie(err)
	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	multResults := make([]int64, 0)
	sumResults := make([]int64, 0)

	for i, line := range lines {
		if i < len(lines)-1 {
			nums := make([]int64, 0)
			for _, v := range strings.Split(strings.Trim(line, " "), " ") {
				if v != "" {
					num, err := strconv.ParseInt(v, 10, 64)
					pie(err)
					nums = append(nums, num)
				}
			}

			if i == 0 {
				for _, n := range nums {
					multResults = append(multResults, n)
					sumResults = append(sumResults, n)
				}
			} else {
				for j, num := range nums {
					multResults[j] *= num
					sumResults[j] += num
				}
			}
		} else {
			s := int64(0)
			j := 0
			for _, v := range strings.Split(strings.Trim(line, " "), " ") {
				if v == "" {
					continue
				}

				switch v {
				case "*":
					s += multResults[j]
				case "+":
					s += sumResults[j]
				default:
					panic("invalid symbol")
				}

				j++
			}
			fmt.Println(s)
		}
	}
}

func part2(lines []string) {
	symbolsTable := make([][]rune, 0)
	for _, l := range lines {
		row := make([]rune, 0)
		for _, c := range l {
			row = append(row, c)
		}
		symbolsTable = append(symbolsTable, row)
	}

	lastRow := symbolsTable[len(symbolsTable)-1]
	op := ' '
	interm := int64(-1)
	s := int64(0)

	for i := range len(lastRow) {
		numStr := ""
		for j := range symbolsTable {
			v := symbolsTable[j][i]

			if v == '*' || v == '+' {
				op = v
			}
			if v >= '0' && v <= '9' {
				numStr += string(v)
			}
		}

		if numStr == "" {
			s += interm
			interm = -1
		} else {
			num, err := strconv.ParseInt(numStr, 10, 64)
			pie(err)

			if interm == -1 {
				interm = num
			} else if op == '+' {
				interm += num
			} else {
				interm *= num
			}
		}
	}

	if interm != -1 {
		s += interm
	}

	fmt.Println(s)
}
