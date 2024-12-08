package main

import (
	"aoc.themacar.com/common"
	"fmt"
)

type Point struct {
	x, y int
}

func parseInput(lines []string) [][]rune {
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}
	return grid
}

func isValidPoint(p Point, rows, cols int) bool {
	return p.x >= 0 && p.x < rows && p.y >= 0 && p.y < cols
}

func findAntennas(grid [][]rune) map[rune][]Point {
	n, m := len(grid), len(grid[0])
	antennas := make(map[rune][]Point)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] != '.' {
				antennas[grid[i][j]] = append(antennas[grid[i][j]], Point{i, j})
			}
		}
	}
	return antennas
}

func checkDirection(start Point, dx, dy int, n, m int, antinodes map[Point]bool, harmonics bool) {
	antinode := Point{start.x - dx, start.y - dy}
	if harmonics {
		antinodes[start] = true
	}
	for isValidPoint(antinode, n, m) {
		antinodes[antinode] = true
		if !harmonics {
			break
		}
		antinode = Point{antinode.x - dx, antinode.y - dy}
	}
}

func findAntinodes(grid [][]rune, harmonics bool) int {
	n, m := len(grid), len(grid[0])
	antennas := findAntennas(grid)
	antinodes := make(map[Point]bool)
	for _, loc_arr := range antennas {
		for i := 0; i < len(loc_arr); i++ {
			for j := i + 1; j < len(loc_arr); j++ {
				p1, p2 := loc_arr[i], loc_arr[j]
				dx := p2.x - p1.x
				dy := p2.y - p1.y
				checkDirection(p1, dx, dy, n, m, antinodes, harmonics)
				checkDirection(p2, -dx, -dy, n, m, antinodes, harmonics)
			}
		}
	}
	return len(antinodes)
}

func main() {
	lines, _ := common.ReadLinesFromFile("input.txt")
	grid := parseInput(lines)

	var ans int

	ans = findAntinodes(grid, false)
	fmt.Printf("Part I: %d\n", ans)
	ans = findAntinodes(grid, true)
	fmt.Printf("Part II: %d\n", ans)
}
