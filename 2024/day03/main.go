package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"aoc.themacar.com/common"
)

func parse_mul_args(args string) ([2]int, error) {
	args = strings.TrimLeft(args, "mul(")
	args = strings.TrimRight(args, ")")
	parts := strings.Split(args, ",")
	if len(parts) != 2 {
		return [2]int{}, fmt.Errorf("Invalid argument format: %s", args)
	}
	var out [2]int
	for i, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			return [2]int{}, fmt.Errorf("Invalid number %q: %v", part, err)
		}
		out[i] = num
	}
	return out, nil
}

func calculate_sum(lines []string, re *regexp.Regexp) int {
	var result int
	active := true

	for _, line := range lines {
		matches := re.FindAllString(line, -1)
		for _, match := range matches {
			switch match {
			case "do()":
				active = true
			case "don't()":
				active = false
			default:
				if active {
					args, err := parse_mul_args(match)
					if err != nil {
						log.Printf("failed to parse arguments in %q: %v", match, err)
						continue
					}
					result += args[0] * args[1]
				}
			}
		}
	}
	return result
}

func main() {
	vals, err := common.ReadLinesFromFile("./input.txt")
	if err != nil {
		log.Fatalf("failed to read input file: %v", err)
	}

	mul_regex := regexp.MustCompile(`mul\(\d+,\d+\)`)
	part1 := calculate_sum(vals, mul_regex)
	fmt.Printf("Part I: %d\n", part1)

	actionRegex := regexp.MustCompile(`do\(\)|don't\(\)|mul\(\d+,\d+\)`)
	part2 := calculate_sum(vals, actionRegex)
	fmt.Printf("Part II: %d\n", part2)
}
