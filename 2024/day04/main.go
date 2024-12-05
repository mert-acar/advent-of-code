package main

import (
	"fmt"
	"regexp"

	"aoc.themacar.com/common"
)

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func transpose(matrix []string) []string {
	n := len(matrix)
	transposed := make([]string, n)
	for i := 0; i < n; i++ {
		var row []rune
		for j := 0; j < n; j++ {
			row = append(row, rune(matrix[j][i]))
		}
		transposed[i] = string(row)
	}
	return transposed
}

func extractDiagonals(matrix []string) []string {
	n := len(matrix)
	var diagonals []string

	// Top-left to bottom-right
	for d := 1 - n; d < n; d++ {
		var diag []rune
		for i := 0; i < n; i++ {
			j := i + d
			if j >= 0 && j < n {
				diag = append(diag, rune(matrix[i][j]))
			}
		}
		if len(diag) > 0 {
			diagonals = append(diagonals, string(diag))
		}
	}

	// Top-right to bottom-left
	for d := 0; d < 2*n-1; d++ {
		var diag []rune
		for i := 0; i < n; i++ {
			j := d - i
			if j >= 0 && j < n {
				diag = append(diag, rune(matrix[i][j]))
			}
		}
		if len(diag) > 0 {
			diagonals = append(diagonals, string(diag))
		}
	}

	return diagonals
}

func countMatches(lines []string, forward, backward *regexp.Regexp) int {
	count := 0
	for _, line := range lines {
		count += len(forward.FindAllString(line, -1))
		count += len(backward.FindAllString(line, -1))
	}
	return count
}

func search(lines []string, forward, backward *regexp.Regexp) int {
	return countMatches(lines, forward, backward) +
		countMatches(transpose(lines), forward, backward) +
		countMatches(extractDiagonals(lines), forward, backward)
}

func searchX(lines []string, searchWord string) int {
	reverseWord := reverse(searchWord)
	n := len(searchWord)
	count := 0

	for i := 0; i <= len(lines)-n; i++ {
		for j := 0; j <= len(lines[i])-n; j++ {
			x1 := string([]byte{lines[i][j], lines[i+1][j+1], lines[i+2][j+2]})
			x2 := string([]byte{lines[i][j+2], lines[i+1][j+1], lines[i+2][j]})

			if (x1 == searchWord || x1 == reverseWord) &&
				(x2 == searchWord || x2 == reverseWord) {
				count++
			}
		}
	}

	return count
}

func main() {
	lines, _ := common.ReadLinesFromFile("./input.txt")
	searchWord := "XMAS"

	forward := regexp.MustCompile(searchWord)
	backward := regexp.MustCompile(reverse(searchWord))

	part1 := search(lines, forward, backward)
	fmt.Printf("Part I: %d\n", part1)

	part2 := searchX(lines, "MAS")
	fmt.Printf("Part II: %d\n", part2)
}
