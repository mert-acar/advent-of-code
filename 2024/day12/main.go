package main

import (
	"aoc.themacar.com/common"
	"fmt"
	"slices"
)

var directions = [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

type Point struct {
	x, y int
}

type Region struct {
	points []Point
	fence  int
}

func (r *Region) area() int {
	return len(r.points)
}

type Queue struct {
	vals []Point
}

func (q *Queue) dequeue() Point {
	p := q.vals[0]
	q.vals = q.vals[1:]
	return p
}

func (q *Queue) enqueue(elem Point) {
	q.vals = append(q.vals, elem)
}

func (q *Queue) isEmpty() bool {
	return len(q.vals) == 0
}

func (q *Queue) contains(p Point) bool {
	for _, c := range q.vals {
		if c == p {
			return true
		}
	}
	return false
}

func isValid(grid [][]rune, plant rune, s Point) bool {
	m, n := len(grid), len(grid[0])
	return 0 <= s.x && s.x < m && 0 <= s.y && s.y < n && grid[s.x][s.y] == plant
}

func exploreRegion(grid [][]rune, start Point) (region Region) {
	plant := grid[start.x][start.y]
	seen := make(map[Point]bool)
	queue := Queue{vals: []Point{start}}

	for !queue.isEmpty() {
		p := queue.dequeue()
		if seen[p] {
			continue
		}
		seen[p] = true
		region.points = append(region.points, p)

		for _, d := range directions {
			next := Point{p.x + d[0], p.y + d[1]}
			if !isValid(grid, plant, next) {
				region.fence++
			} else if !seen[next] {
				queue.enqueue(next)
			}
		}
	}
	return
}

func countCorners(region Region) int {
	cornerChecks := [4][3][2]int{
		// N, E, NW
		{{-1, 0}, {0, -1}, {-1, -1}},
		// N, W, NE
		{{-1, 0}, {0, 1}, {-1, 1}},
		// S, W, NW
		{{1, 0}, {0, -1}, {1, -1}},
		// S, E, NE
		{{1, 0}, {0, 1}, {1, 1}},
	}
	corners := 0
	for _, p := range region.points {
		for _, checks := range cornerChecks {
			// Convex Corner Check
			if (slices.Contains(region.points, Point{p.x + checks[0][0], p.y + checks[0][1]}) &&
				slices.Contains(region.points, Point{p.x + checks[1][0], p.y + checks[1][1]}) &&
				!slices.Contains(region.points, Point{p.x + checks[2][0], p.y + checks[2][1]})) {
				corners++
			}
			// Concave Corner Check
			if !(slices.Contains(region.points, Point{p.x + checks[0][0], p.y + checks[0][1]}) ||
				slices.Contains(region.points, Point{p.x + checks[1][0], p.y + checks[1][1]})) {
				corners++
			}
		}
	}
	return corners
}

func getRegions(grid [][]rune) []Region {
	visited := make(map[Point]bool)
	regions := []Region{}
	for i, row := range grid {
		for j := range row {
			start := Point{i, j}
			if visited[start] {
				continue
			}

			region := exploreRegion(grid, start)
			for _, p := range region.points {
				visited[p] = true
			}
			regions = append(regions, region)
		}
	}
	return regions
}

func calculateCost(regions []Region, cornerMultiplier bool) int {
	totalCost := 0
	for _, region := range regions {
		regionSize := region.area()
		if cornerMultiplier {
			totalCost += regionSize * countCorners(region)
		} else {
			totalCost += regionSize * region.fence
		}
	}
	return totalCost
}

func main() {
	grid, err := common.ReadRuneGrid("./input.txt")
	if err != nil {
		panic(err)
	}

	regions := getRegions(grid)

	fmt.Printf("Part I: %d\n", calculateCost(regions, false))
	fmt.Printf("Part II: %d\n", calculateCost(regions, true))
}
