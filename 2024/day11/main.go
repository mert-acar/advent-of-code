package main

import (
	"fmt"
	"strconv"
)

func blink(stone int) []int {
	if stone == 0 {
		return []int{1}
	}

	s := strconv.Itoa(stone)
	if len(s)%2 == 0 {
		mid := len(s) / 2
		part1, _ := strconv.Atoi(s[:mid])
		part2, _ := strconv.Atoi(s[mid:])
		return []int{part1, part2}
	}

	return []int{stone * 2024}
}

func run(input []int, iterations int) int {
	stones := make(map[int]int)
	for _, stone := range input {
		stones[stone]++
	}

	for i := 0; i < iterations; i++ {
		newStones := make(map[int]int)
		for rock, count := range stones {
			blinkResults := blink(rock)
			for _, blinkResult := range blinkResults {
				newStones[blinkResult] += count
			}
		}
		stones = newStones
	}

	total := 0
	for _, count := range stones {
		total += count
	}
	return total
}

func main() {
	input := []int{773, 79858, 0, 71, 213357, 2937, 1, 3998391}

	ans := run(input, 25)
	fmt.Printf("Part I: %d\n", ans)

	ans = run(input, 75)
	fmt.Printf("Part II: %d\n", ans)
}
