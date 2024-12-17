package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Coord struct {
	X, Y int64
}

const OFFSET = 10000000000000

func parseInputFile(filename string) ([][][]int64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data [][][]int64
	var currentGroup [][]int64
	scanner := bufio.NewScanner(file)

	buttonRegex := regexp.MustCompile(`\d+`)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		matches := buttonRegex.FindAllString(line, -1)
		var numbers []int64
		for _, match := range matches {
			num, err := strconv.ParseInt(match, 10, 64)
			if err != nil {
				return nil, err
			}
			numbers = append(numbers, num)
		}

		currentGroup = append(currentGroup, numbers)

		if len(currentGroup) == 3 {
			data = append(data, currentGroup)
			currentGroup = nil
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

func main() {
	data, err := parseInputFile("input.txt")
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	var total int64
	for _, m := range data {
		a := Coord{X: m[0][0], Y: m[0][1]}
		b := Coord{X: m[1][0], Y: m[1][1]}
		c := Coord{X: m[2][0] + OFFSET, Y: m[2][1] + OFFSET}

		a1, b1, c1 := a.X, b.X, -c.X
		a2, b2, c2 := a.Y, b.Y, -c.Y

		x := b1*c2 - c1*b2
		y := c1*a2 - a1*c2
		z := a1*b2 - b1*a2

		if x%z != 0 || y%z != 0 {
			continue
		}

		x /= z
		y /= z

		if x >= 0 && y >= 0 {
			total += x*3 + y
		}
	}

	fmt.Println(total)
}
