package aoc

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func day5(inputFile string) {
	lines, err := fileReadLines(inputFile)
	pie(err)

	ranges := make([][2]int64, 0)

	numFresh := 0
	readingRanges := true
	for _, l := range lines {
		if l == "" && readingRanges {
			readingRanges = false
			continue
		}

		if readingRanges {
			nums := strings.Split(l, "-")
			n1, err := strconv.ParseInt(nums[0], 10, 64)
			pie(err)
			n2, err := strconv.ParseInt(nums[1], 10, 64)
			pie(err)
			ranges = append(ranges, [2]int64{n1, n2})
		} else {
			num, err := strconv.ParseInt(l, 10, 64)
			pie(err)
			if someRangeContains(ranges, num) {
				numFresh++
			}
		}
	}
	fmt.Println(numFresh)

	cntFresh := int64(0)
	for _, rng := range mergeRanges(ranges) {
		cntFresh += rng[1] - rng[0] + 1
	}
	fmt.Println(cntFresh)
}

func someRangeContains(ranges [][2]int64, num int64) bool {
	for _, r := range ranges {
		if num >= r[0] && num <= r[1] {
			return true
		}
	}
	return false
}

func mergeRanges(ranges [][2]int64) [][2]int64 {
	if len(ranges) == 0 {
		return ranges
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	merged := make([][2]int64, 0)
	current := ranges[0]

	for _, r := range ranges[1:] {
		if r[0] <= current[1] {
			if r[1] > current[1] {
				current[1] = r[1]
			}
		} else {
			merged = append(merged, current)
			current = r
		}
	}

	merged = append(merged, current)
	return merged
}
