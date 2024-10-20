package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"aoc.themacar.com/common"
)

type Point struct {
	x, y int
}

func traceWire(path []string) map[Point]int {
	x, y := 0, 0
	steps := 0
	grid := make(map[Point]int)
	for _, instruction := range path {
		direction := instruction[0]
		amount, _ := strconv.Atoi(instruction[1:])
		for range amount {
			steps++
			switch direction {
			case 'R':
				x++
			case 'L':
				x--
			case 'D':
				y--
			case 'U':
				y++
			}
			grid[Point{x, y}] = steps
		}
	}
	return grid
}

func minSignalLoss(path1, path2 map[Point]int) int {
	minS := math.MaxInt32
	for point, steps := range path1 {
		if _, ok := path2[point]; ok {
			dist := steps + path2[point]
			if dist < minS {
				minS = dist
			}
		}
	}
	return minS
}

func minDistance(path1, path2 map[Point]int) int {
	minD := math.MaxInt32
	for point := range path1 {
		if _, ok := path2[point]; ok {
			dist := int(math.Abs(float64(point.x))) + int(math.Abs(float64(point.y)))
			if dist < minD {
				minD = dist
			}
		}
	}
	return minD
}

func main() {
	vals, err := common.ReadLinesFromFile("./input.txt")
	if err != nil {
		panic(err)
	}

	wire1 := traceWire(strings.Split(vals[0], ","))
	wire2 := traceWire(strings.Split(vals[1], ","))

	fmt.Println("part I:", minDistance(wire1, wire2))
	fmt.Println("part II:", minSignalLoss(wire1, wire2))
}
