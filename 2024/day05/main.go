package main

import (
	"fmt"
	"log"
	"os"
  "slices"
	"strconv"
	"strings"
)

func parseInputs(vals string) (map[int][]int, [][]int) {
	parts := strings.Split(vals, "\n\n")
	if len(parts) != 2 {
		log.Printf("Error parsing the file")
	}

	rules := strings.Split(string(parts[0]), "\n")
	ordering_rules := make(map[int][]int)
	for _, line := range rules {
		p := strings.Split(line, "|")
		x, _ := strconv.Atoi(p[0])
		y, _ := strconv.Atoi(p[1])
		ordering_rules[x] = append(ordering_rules[x], y)
	}

  update := strings.Split(string(parts[1]), "\n")
  lines := make([][]int, len(update) - 1)
  for i := 0; i < len(update) - 1; i++ {
    part := update[i]
    line := strings.Split(part, ",")
    var nums []int
    for j := range line {
      num, _ := strconv.Atoi(line[j])
      nums = append(nums, num)
    }
    lines[i] = nums
  }
  return ordering_rules, lines
}

func checkRow(row []int, rules map[int][]int) bool {
  for i, num := range row {
    if vals, ok := rules[num]; ok {
      if i > 0 {
        prev := row[0:i] 
        for _, p := range prev {
          if slices.Contains(vals, p) {
            return false
          }
        }
      }
    }
  }
  return true
}

func fixRow(row []int, rules map[int][]int) []int {
  for i, num := range row {
    if vals, ok := rules[num]; ok {
      if i > 0 {
        prev := row[0:i] 
        for _, p := range prev {
          if slices.Contains(vals, p) {
            j := slices.Index(row, p)
            row[j], row[i] = row[i], row[j] 
          }
        }
      }
    }
  }
  return row
}

func main() {
	val_arr, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Printf("File cannot be read: %v", err)
	}
	vals := string(val_arr)
  ordering_rules, lines := parseInputs(vals)

  result := 0
  result2 := 0
  for _, row := range lines {
    if checkRow(row, ordering_rules) {
      result += row[len(row) / 2]
    } else {
      row = fixRow(row, ordering_rules)
      result2 += row[len(row) / 2]
    }
  }

  fmt.Printf("Part I: %d\n", result)
  fmt.Printf("Part II: %d\n", result2)
}
