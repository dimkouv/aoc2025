package aoc

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Point2D struct {
	X, Y int64
}

type Rectangle struct {
	p1, p2 Point2D
}

func (r Rectangle) area() int64 {
	width := r.p1.X - r.p2.X
	if width < 0 {
		width = -width
	}

	height := r.p1.Y - r.p2.Y
	if height < 0 {
		height = -height
	}

	width += 1
	height += 1

	return width * height
}

func day9(inputFile string) {
	points := parsePoints(inputFile)

	maxArea := int64(-1)
	for i, p1 := range points[:len(points)-1] {
		for _, p2 := range points[i+1:] {
			r := Rectangle{p1, p2}
			area := r.area()
			if area > int64(maxArea) {
				maxArea = area
			}
		}
	}
	fmt.Println(maxArea) // p1

	var xs, ys []int64
	xMap := make(map[int64]struct{})
	yMap := make(map[int64]struct{})
	for _, p := range points {
		xMap[p.X] = struct{}{}
		yMap[p.Y] = struct{}{}
	}
	for x := range xMap {
		xs = append(xs, x)
	}
	for y := range yMap {
		ys = append(ys, y)
	}
	slices.Sort(xs)
	slices.Sort(ys)

	pointInShapeMem := make([][]bool, len(xs)-1)
	for i := range xs[:len(xs)-1] {
		pointInShapeMem[i] = make([]bool, len(ys))
		for j := range ys[:len(ys)-1] {
			midX := float64(xs[i]+xs[i+1]) / 2.0
			midY := float64(ys[j]+ys[j+1]) / 2.0
			if isPointInShape(midX, midY, points) {
				pointInShapeMem[i][j] = true
			}
		}
	}

	var maxAreaPart2 int64 = 0
	for i, p1 := range points {
		for _, p2 := range points[i+1:] {
			r := Rectangle{p1, p2}
			minX, maxX := min64(p1.X, p2.X), max64(p1.X, p2.X)
			minY, maxY := min64(p1.Y, p2.Y), max64(p1.Y, p2.Y)

			isValid := true
			for ix := range xs[:len(xs)-1] {
				if xs[ix] < minX || xs[ix+1] > maxX {
					continue
				}

				for iy := range ys[:len(ys)-1] {
					if ys[iy] < minY || ys[iy+1] > maxY {
						continue
					}

					if !pointInShapeMem[ix][iy] {
						isValid = false
						break
					}
				}

				if !isValid {
					break
				}
			}

			if area := r.area(); isValid && area > maxAreaPart2 {
				maxAreaPart2 = area
			}
		}
	}
	fmt.Println(maxAreaPart2) // p2
}

func isPointInShape(x, y float64, points []Point2D) bool {
	inside := false
	for i := range points {
		p1 := points[i]
		p2 := points[(i+1)%len(points)]

		if p1.X == p2.X {
			minY, maxY := minMax(p1.Y, p2.Y)
			if float64(p1.X) >= x && y >= float64(minY) && y <= float64(maxY) {
				inside = !inside
			}
		}
	}
	return inside
}

func min64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
func max64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func minMax(a, b int64) (int64, int64) {
	if a < b {
		return a, b
	}
	return b, a
}

func parsePoints(inputFile string) []Point2D {
	lines, err := fileReadLines(inputFile)
	pie(err)

	points := make([]Point2D, 0)
	for _, ln := range lines {
		parts := strings.Split(ln, ",")
		x, _ := strconv.ParseInt(parts[0], 10, 64)
		y, _ := strconv.ParseInt(parts[1], 10, 64)
		points = append(points, Point2D{x, y})
	}
	return points
}
