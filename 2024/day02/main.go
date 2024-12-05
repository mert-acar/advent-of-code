package main

import (
	"fmt"
	"strconv"
	"strings"

	"aoc.themacar.com/common"
)

func can_be_safe_with_dampener(nums []int) bool {
	// Check each possible removal of one element
	for i := range nums {
		modified := append([]int{}, nums[:i]...)
		modified = append(modified, nums[i+1:]...)
		difference := common.Diff(modified, 1)
		if is_safe(difference) {
			return true
		}
	}
	return false
}

func is_safe(difference []int) bool {
	increasing := difference[0] > 0
	for i := range difference {
		x := common.Abs(difference[i])
		if (increasing && difference[i] < 0) || (!increasing && difference[i] > 0) {
			return false
		} else if (x < 1) || (x > 3) {
			return false
		}
	}
	return true
}

func main() {
	vals, err := common.ReadLinesFromFile("./input.txt")
	if err != nil {
		panic(err)
	}
	count1 := 0
	count2 := 0
	var nums []int
	for _, report := range vals {
		nums = nil // clear the slice
		levels := strings.Fields(report)
		for _, level := range levels {
			num, _ := strconv.Atoi(level)
			nums = append(nums, num)
		}
    difference := common.Diff(nums, 1)

		if is_safe(difference) {
			count1++
      count2++
    } else if can_be_safe_with_dampener(nums) {
			count2++ // Safe with dampener
		}
    
	}
	fmt.Printf("Part I: %d\n", count1)
	fmt.Printf("Part II: %d\n", count2)
}
