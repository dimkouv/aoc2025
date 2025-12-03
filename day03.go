package aoc

import (
	"fmt"
	"strconv"
)

func day3(inputFile string) {
	banks, err := fileReadLines(inputFile)
	pie(err)

	sum := 0
	sumP2 := int64(0)
	for _, bank := range banks {
		maxJolt := findMaxJolt(bank)
		sum += maxJolt

		maxJoltP2 := findMaxJoltP2(bank)
		sumP2 += maxJoltP2
	}

	fmt.Println(sum)
	fmt.Println(sumP2)
}

func findMaxJolt(bank string) int {
	a, rem := findNextBigJolt(bank, 2)
	b, _ := findNextBigJolt(rem, 1)
	num := a*10 + b
	return num
}

func findMaxJoltP2(bank string) int64 {
	rem := bank
	numStr := ""

	for i := 12; i >= 1; i-- {
		var d int
		d, rem = findNextBigJolt(rem, i)
		numStr += strconv.Itoa(d)
	}
	num, err := strconv.ParseInt(numStr, 10, 64)
	pie(err)
	return num
}

func findNextBigJolt(bank string, size int) (int, string) {
	a := rune(bank[0])
	posA := 0

	for i, v := range bank[:len(bank)-size+1] {
		if v > a {
			a = rune(v)
			posA = i
		}
	}

	return int(a - '0'), bank[posA+1:]
}
