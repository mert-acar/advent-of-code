package main

import (
	"fmt"
	"gonum.org/v1/gonum/stat"
	"io"
	"os"
	"strings"
)

type Robot struct {
	p, v [2]int
}

func parseInputs(bytes []uint8) []Robot {
	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")
	robots := []Robot{}
	var p, v [2]int
	for _, line := range lines {
		_, err := fmt.Sscanf(line, "p=%d,%d v=%d,%d", &p[0], &p[1], &v[0], &v[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not parse string: %s", line)
			continue
		}
		robots = append(robots, Robot{p, v})
	}
	return robots
}

func wrapCoordinate(coord, size int) int {
	wrapped := coord % size
	if wrapped < 0 {
		wrapped += size
	}
	return wrapped
}

func simulate(robots []Robot, s, m, n int) []Robot {
  var simRobots []Robot
	for j := range robots {
		p := robots[j].p
		v := robots[j].v
		newP := [2]int{
			wrapCoordinate(p[0]+s*v[0], n),
			wrapCoordinate(p[1]+s*v[1], m),
		}
    simRobots = append(simRobots, Robot{p: newP, v: v})
	}
	return simRobots
}

func main() {
	bytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not read file: %v\n", err)
		os.Exit(1)
	}
	robots := parseInputs(bytes)

	s, m, n := 100, 103, 101
	grid := make([][]int, m)
	for i := range grid {
		grid[i] = make([]int, n)
	}

	simRobots := simulate(robots, s, m, n)

	for j := range simRobots {
		grid[simRobots[j].p[1]][simRobots[j].p[0]]++
	}

	sumQuadrant := func(rowStart, rowEnd, colStart, colEnd int) int {
		sum := 0
		for i := rowStart; i < rowEnd; i++ {
			for j := colStart; j < colEnd; j++ {
				sum += grid[i][j]
			}
		}
		return sum
	}

	total := 1

	total *= sumQuadrant(0, m/2, 0, n/2)     // Top-left
	total *= sumQuadrant(m/2+1, m, 0, n/2)   // Bottom-left
	total *= sumQuadrant(0, m/2, n/2+1, n)   // Top-right
	total *= sumQuadrant(m/2+1, m, n/2+1, n) // Bottom-right
	fmt.Printf("Part I: %d\n", total)

	var xs, ys []float64
	bt, bvar := 0, 1000000000000.0
	for t := 0; t < 10000; t++ {
    xs, ys = nil, nil
    simRobots := simulate(robots, t, m, n)
		for _, robot := range simRobots {
			xs = append(xs, float64(robot.p[0]))
			ys = append(ys, float64(robot.p[1]))
		}
		variance := stat.Variance(xs, nil) + stat.Variance(ys, nil)
		if variance < bvar {
			bt, bvar = t, variance
		}
	}
	fmt.Printf("Part II: %d\n", bt)
}
