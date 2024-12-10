package main

import (
	"aoc.themacar.com/common"
	"fmt"
	"os"
)

var directions = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

type Point struct {
	x, y int
}

func parseInput(lines []string) (grid [][]int) {
	grid = make([][]int, len(lines))
	for i, line := range lines {
		grid[i] = make([]int, len(line))
		for j, ch := range line {
			grid[i][j] = int(ch - '0')
		}
	}
	return
}

func findTrailHeads(grid [][]int) (heads []Point) {
	for i, row := range grid {
		for j, elem := range row {
			if elem == 0 {
				heads = append(heads, Point{i, j})
			}
		}
	}
	return
}

func isValid(point Point, n, m int) bool {
	return 0 <= point.x && point.x < n && 0 <= point.y && point.y < m
}

func nextSteps(grid [][]int, pos Point) (next []Point) {
	n, m := len(grid), len(grid[0])
	for _, dir := range directions {
		candidate := Point{pos.x + dir[0], pos.y + dir[1]}
		if isValid(candidate, n, m) {
			if grid[candidate.x][candidate.y] == grid[pos.x][pos.y]+1 {
				next = append(next, candidate)
			}
		}
	}
	return
}

func findTrailScore(grid [][]int, head Point) int {
	height := 0
	positions := map[Point]bool{head: true}
	for height != 9 {
		newPositions := make(map[Point]bool)
		for position := range positions {
			for _, next := range nextSteps(grid, position) {
				newPositions[next] = true
			}
		}
		positions = newPositions
		height++
	}
	return len(positions)
}

func findTrailRaiting(grid [][]int) []int {
	results := []int{}
	raitings := make(map[Point]int)
	for height := 9; height > -1; height-- {
		for x, row := range grid {
			for y, h := range row {
				if h == height {
					pos := Point{x, y}
					if height == 9 {
						raitings[pos] = 1
					} else {
						raitings[pos] = 0
						for _, next := range nextSteps(grid, pos) {
							raitings[pos] += raitings[next]
						}
						if height == 0 {
							results = append(results, raitings[pos])
						}
					}
				}
			}
		}
	}
	return results
}

func walkTrail(grid [][]int) (int, int) {
	trailHeads := findTrailHeads(grid)
	score := 0
	for _, head := range trailHeads {
		score += findTrailScore(grid, head)
	}
	raiting := common.Sum1D(findTrailRaiting(grid))
	return score, raiting
}

func main() {
	lines, err := common.ReadLinesFromFile("./input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading the input file: %v\n", err)
		os.Exit(1)
	}
	grid := parseInput(lines)
	score, raiting := walkTrail(grid)
	fmt.Printf("Part I: %d\n", score)
	fmt.Printf("Part II: %d\n", raiting)
}
