package main

import (
	"fmt"
	"strconv"
	"strings"

	"aoc.themacar.com/common"
)

func parseInputs(vals []string) [][]int {
  var num int
  var nums []int
  out := make([][]int, len(vals))
  for i, line := range vals {
    nums = nil
    line = strings.Replace(line, ":", "", -1)
    parts := strings.Fields(line) 
    for _, part := range parts {
      num, _ = strconv.Atoi(part)
      nums = append(nums, num)
    }
    out[i] = nums
  }
  return out
}

func getAllPossibleCombinations(choices []string, n int) [][]string {
  totalCombinations := common.Pow(len(choices), n)
  out := make([][]string, totalCombinations)
  for i := 0; i < totalCombinations; i++ {
		combination := make([]string, n)
		remainder := i
		for j := 0; j < n; j++ {
			combination[j] = choices[remainder%len(choices)]
			remainder /= len(choices)
		}
		out[i] = combination
	}
  return out
}

func tryCombinations(nums []int, target int, choices []string) bool {
  combinations := getAllPossibleCombinations(choices, len(nums) - 1)
  var i, result int 
  for _, combination := range combinations {
    result = nums[0]
    i = 1
    for _, operation := range combination {
      if operation == "add" {
        result += nums[i]
      } else if operation == "mul" {
        result *= nums[i]
      } else if operation == "concat" {
        left := strconv.Itoa(result)
        right := strconv.Itoa(nums[i])
        result, _ = strconv.Atoi(left + right)
      }
      i++
    }
    if result == target {
      return true
    }
  }
  return false
}

func part1(lines [][]int) int {
  total := 0
  choices := []string{"add", "mul"}
  for _, nums := range lines {
    if tryCombinations(nums[1:], nums[0], choices) {
      total += nums[0]
    }
  }
  return total
}

func part2(lines [][]int) int {
  total := 0
  choices := []string{"add", "mul", "concat"}
  for _, nums := range lines {
    if tryCombinations(nums[1:], nums[0], choices) {
      total += nums[0]
    }
  }
  return total
}

func main() {
  lines, _ := common.ReadLinesFromFile("./input.txt")
  nums := parseInputs(lines)
  ans := part1(nums)
  fmt.Printf("Part I: %d\n", ans)
  ans = part2(nums)
  fmt.Printf("Part I: %d\n", ans)
}
