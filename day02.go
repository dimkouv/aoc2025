package aoc

import (
	"fmt"
	"strconv"
	"strings"
)

func day2(inputFile string) {
	l, err := fileReadLines(inputFile)
	pie(err)

	ranges := strings.Split(l[0], ",")

	sumInvalidIds := int64(0)
	sumInvalidIdsPart2 := int64(0)

	for _, rng := range ranges {
		nums := strings.Split(rng, "-")
		n1, err := strconv.ParseInt(nums[0], 10, 64)
		pie(err)
		n2, err := strconv.ParseInt(nums[1], 10, 64)
		pie(err)

		for i := n1; i <= n2; i++ {
			iStr := strconv.FormatInt(i, 10)
			if isInvalid(iStr) {
				sumInvalidIds += i
			}

			if isInvalidPart2(iStr) {
				sumInvalidIdsPart2 += i
			}
		}
	}

	fmt.Println(sumInvalidIds)
	fmt.Println(sumInvalidIdsPart2)
}

func isInvalid(id string) bool {
	firstHalf := id[:len(id)/2]
	secondHalf := id[len(id)/2:]
	return strings.EqualFold(firstHalf, secondHalf)
}

func isInvalidPart2(id string) bool {
	for i := 1; i <= len(id)/2; i++ {
		isInvalid := true
		originPiece := id[:i]

		for j := i; j < len(id); j += i {
			if j+i > len(id) {
				isInvalid = false
				break
			}

			currPiece := id[j : j+i]
			if currPiece != originPiece {
				isInvalid = false
				break
			}
		}

		if isInvalid {
			return true
		}
	}

	return false
}
