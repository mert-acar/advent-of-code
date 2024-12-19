package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// countArrangements counts the number of possible arrangements without storing them
func countArrangements(design string, patterns []string, memo map[string]int64) int64 {
	if design == "" {
		return 1
	}
	
	if result, exists := memo[design]; exists {
		return result
	}

	var count int64 = 0
	for _, pattern := range patterns {
		if len(pattern) <= len(design) && strings.HasPrefix(design, pattern) {
			remainingDesign := design[len(pattern):]
			count += countArrangements(remainingDesign, patterns, memo)
		}
	}

	memo[design] = count
	return count
}

// canMakeDesign checks if a design can be made using available patterns
func canMakeDesign(design string, patterns []string, memo map[string]bool) bool {
	if design == "" {
		return true
	}
	
	if result, exists := memo[design]; exists {
		return result
	}

	for _, pattern := range patterns {
		if len(pattern) <= len(design) {
			if strings.HasPrefix(design, pattern) {
				remainingDesign := design[len(pattern):]
				if canMakeDesign(remainingDesign, patterns, memo) {
					memo[design] = true
					return true
				}
			}
		}
	}

	memo[design] = false
	return false
}

func parseInput(bytes []uint8) ([]string, []string) {
	parts := strings.Split(strings.TrimSpace(string(bytes)), "\n\n")
	patternsStr := parts[0]
	designsStr := parts[1]
	patterns := strings.Split(patternsStr, ", ")
	designs := strings.Split(strings.TrimSpace(designsStr), "\n")
	return patterns, designs
}

func main() {
	bytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading the input file: %v\n", err)
		os.Exit(1)
	}
	patterns, designs := parseInput(bytes)

	// Count possible designs
	possibleCount := 0
	for _, design := range designs {
		if design == "" {
			continue
		}
		memo := make(map[string]bool)
		possible := canMakeDesign(design, patterns, memo)
		if possible {
			possibleCount++
		}
	}

	fmt.Printf("Part I: %d\n", possibleCount)

	// Part 2: Count all possible arrangements
	var totalArrangements int64 = 0
	for _, design := range designs {
		if design == "" {
			continue
		}
		memo := make(map[string]int64)
		arrangements := countArrangements(design, patterns, memo)
		totalArrangements += arrangements
	}
	fmt.Printf("Part II: %d\n", totalArrangements)
}
