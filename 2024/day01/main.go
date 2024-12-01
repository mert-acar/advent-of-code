package main

import (
	"fmt"
	"strconv"
	"slices"
	"strings"

	"aoc.themacar.com/common"
)

func abs(x int) int{
  if x < 0 {
    return -x
  }
  return x
}

func main() {
  vals, err := common.ReadLinesFromFile("./input.txt")
	if err != nil {
		panic(err)
	}

  var left, right []int
	for _, v := range vals {
    nums := strings.Fields(v)
    for i, num := range nums {
      n, _ := strconv.Atoi(num)
      if i == 0 {
        left = append(left, n)
      } else {
        right = append(right, n)
      }
    }
	}  

  slices.Sort(left)
  slices.Sort(right)

  total := 0
  counts := make(map[int]int)
  for i := range left {
    total += abs(left[i] - right[i])
    counts[right[i]]++
  }
  fmt.Printf("Part I: %d\n", total)

  similarity_score := 0
  for i := range left {
    similarity_score += left[i] * counts[left[i]]
  }
  fmt.Printf("Part II: %d\n", similarity_score)
}
