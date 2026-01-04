package aoc

import (
	"fmt"
	"strconv"
	"strings"
)

type Shape struct {
	points map[Point2D]struct{}
}

func NewShape(points map[Point2D]struct{}) *Shape {
	return &Shape{points}
}

func (s *Shape) addPoint(p Point2D) {
	s.points[p] = struct{}{}
}

func (s Shape) area() int {
	return len(s.points)
}

type Board struct {
	SizeX  int64
	SizeY  int64
	Shapes []int
}

func parseShapes(lines []string) []Shape {
	idx := -1

	shapes := make([]Shape, 0)
	currentShapePoints := make(map[Point2D]struct{}, 0)

	i := 0
	for _, ln := range lines {
		if strings.HasSuffix(ln, ":") {
			i = -1
			idx += 1
		}

		if ln == "" {
			shape := NewShape(currentShapePoints)
			shapes = append(shapes, *shape)
			i = 0
			currentShapePoints = make(map[Point2D]struct{}, 0)
		}

		if strings.Contains(ln, "x") {
			break
		}

		for j, char := range ln {
			if char == '#' {
				currentShapePoints[Point2D{X: int64(i), Y: int64(j)}] = struct{}{}
			}
		}

		i++
	}

	return shapes
}

func parseBoards(lines []string) []Board {
	startParsing := false
	boards := make([]Board, 0)
	for _, ln := range lines {
		if !startParsing && strings.Contains(ln, "x") {
			startParsing = true
		}
		if !startParsing {
			continue
		}

		parts := strings.Split(ln, ": ")
		size := parts[0]
		sizeParts := strings.Split(size, "x")
		sizeX, err := strconv.ParseInt(sizeParts[0], 10, 64)
		pie(err)
		sizeY, err := strconv.ParseInt(sizeParts[1], 10, 64)
		pie(err)

		shapes := make([]int, 0)
		shapeCounts := strings.Split(parts[1], " ")
		for _, cntStr := range shapeCounts {
			cnt, err := strconv.ParseInt(cntStr, 10, 64)
			pie(err)
			shapes = append(shapes, int(cnt))
		}

		board := Board{sizeX, sizeY, shapes}
		boards = append(boards, board)
	}

	return boards
}

func day12(inputFile string) {
	lines, err := fileReadLines(inputFile)
	pie(err)

	shapes := parseShapes(lines)
	boards := parseBoards(lines)

	cnt := 0
	for _, board := range boards {
		boardArea := board.SizeX * board.SizeY

		shapesArea := 0
		for shapeIdx, shapesCnt := range board.Shapes {
			shape := shapes[shapeIdx]
			shapesArea += shape.area() * shapesCnt
		}

		if boardArea >= int64(shapesArea) {
			cnt++
		}
	}

	fmt.Println(cnt)
}
