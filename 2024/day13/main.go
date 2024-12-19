package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

type Point struct {
	x, y int
}

func parseInput(filename string) (data [][][]int) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading the input file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()
	var curr [][]int
	buttonRegex := regexp.MustCompile(`-?\d+`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if len(curr) == 3 {
				data = append(data, curr)
				curr = nil
			}
			continue
		}
		matches := buttonRegex.FindAllString(line, -1)
		var numbers []int
		for _, match := range matches {
			num, _ := strconv.Atoi(match)
			numbers = append(numbers, num)
		}
		curr = append(curr, numbers)
	}
	if len(curr) == 3 {
		data = append(data, curr)
	}
	return
}

func solveMachine(m [][]int, maxPresses int, offset int64) (bool, float64, float64) {
	det := float64((m[0][0] * m[1][1]) - (m[1][0] * m[0][1]))
	if math.Abs(det) < 1e-10 {
		return false, 0, 0
	}

	targetX := float64(m[2][0] + int(offset))
	targetY := float64(m[2][1] + int(offset))

	// x0yb - y0xb
	x_coef := (float64(m[1][1])*targetX - float64(m[1][0])*targetY) / det
	// y0xa - x0ya
	y_coef := (float64(m[0][0])*targetY - float64(m[0][1])*targetX) / det

	// Check if coefficients are "effectively" integers
	if math.Abs(x_coef-math.Round(x_coef)) < 1e-10 &&
		math.Abs(y_coef-math.Round(y_coef)) < 1e-10 {
		x_coef = math.Round(x_coef)
		y_coef = math.Round(y_coef)

		if x_coef >= 0 && y_coef >= 0 {
			if maxPresses == 0 || (x_coef <= float64(maxPresses) && y_coef <= float64(maxPresses)) {
				return true, x_coef, y_coef
			}
		}
	}
	return false, 0, 0
}

func main() {
	data := parseInput("./input.txt")

	// Part 1
	total1 := 0.0
	for _, m := range data {
		if possible, a, b := solveMachine(m, 100, 0); possible {
			total1 += (a * 3) + b
		}
	}
	fmt.Printf("Part 1: %d\n", int(math.Round(total1)))

	// Part 2
	total2 := 0.0
	offset := int64(10000000000000)
	for _, m := range data {
		if possible, a, b := solveMachine(m, 0, offset); possible {
			total2 += (a * 3) + b
		}
	}
	fmt.Printf("Part 2: %d\n", int(math.Round(total2)))
}
