package main

import (
	"aoc.themacar.com/common"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isValid(p Point, n int) bool {
	return 0 <= p.x && p.x < n && 0 <= p.y && p.y < n
}

type Point struct {
	x, y int
}

type Queue struct {
	vals [][]Point
}

func (q *Queue) dequeue() []Point {
	p := q.vals[0]
	q.vals = q.vals[1:]
	return p
}

func (q *Queue) enqueue(elem []Point) {
	q.vals = append(q.vals, elem)
}

func (q *Queue) isEmpty() bool {
	return len(q.vals) == 0
}

var directions = []Point{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func shortestPathBFS(grid [][]rune, start, end Point) []Point {
	n := len(grid)
	if !isValid(start, n) || !isValid(end, n) {
		return []Point{}
	}

	to_visit := Queue{[][]Point{{start}}}
	visited := make(map[Point]bool)
	for !to_visit.isEmpty() {
		path := to_visit.dequeue()
		node := path[len(path)-1]
		if node == end {
			return path
		}
		if !visited[node] {
			visited[node] = true
			for _, d := range directions {
				next := Point{node.x + d.x, node.y + d.y}
				if isValid(next, n) && grid[next.y][next.x] != '#' && !visited[next] {
					new_path := make([]Point, len(path))
					copy(new_path, path)
					new_path = append(new_path, next)
					to_visit.enqueue(new_path)
				}
			}
		}
	}
	fmt.Println("No path found.")
	return []Point{}
}

func main() {
	// n, u := 7, 12
	n, u := 71, 1024
	lines, err := common.ReadLinesFromFile("./input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading the input file: %v\n", err)
		os.Exit(1)
	}

	coords := []Point{}
	for _, line := range lines {
		coord := strings.Split(line, ",")
		x, _ := strconv.Atoi(coord[0])
		y, _ := strconv.Atoi(coord[1])
		coords = append(coords, Point{x, y})
	}

	grid := make([][]rune, n)
	for i := range grid {
		grid[i] = make([]rune, n)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}

	for _, point := range coords[:u] {
		grid[point.y][point.x] = '#'
	}

	path := shortestPathBFS(grid, Point{0, 0}, Point{n - 1, n - 1})
	for _, p := range path {
		grid[p.y][p.x] = 'O'
	}
	for _, row := range grid {
		for _, r := range row {
			fmt.Print(string(r))
		}
		fmt.Println()
	}
	fmt.Printf("Part I: %d\n", len(path)-1)

	new_grid := make([][]rune, n)
	copy(new_grid, grid)
	for _, point := range coords[u:] {
		new_grid[point.y][point.x] = '#'
		path = shortestPathBFS(new_grid, Point{0, 0}, Point{n - 1, n - 1})
		if len(path) == 0 {
			fmt.Printf("Part II: %d, %d\n", point.x, point.y)
			break
		}
	}

}
