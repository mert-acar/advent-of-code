package main

import (
	"fmt"
	"strconv"
)

func isOk2(candidate string) bool {
	sameAdjacent := false
	i := 0
	for i < len(candidate)-1 {
		if candidate[i] > candidate[i+1] {
			return false
		}

		// Find groups of matching digits
		runLength := 1
		for i+runLength < len(candidate) && candidate[i] == candidate[i+runLength] {
			runLength++
		}

		// Check if this run contains exactly two digits
		if runLength == 2 {
			sameAdjacent = true
		}

    // HORRIBLE
    if runLength > 1 {
      i += runLength - 1
    } else {
      i += runLength
    }

	}
	return sameAdjacent
}

func isOk(candidate string) bool {
	sameAdjacent := false
	for i := 0; i < len(candidate)-1; i++ {
		if candidate[i] == candidate[i+1] {
			sameAdjacent = true
		}
		if candidate[i] > candidate[i+1] {
			return false
		}
	}
	return sameAdjacent
}

func main() {
	low, high := 153517, 630395
	// low := 599566
	// high := low + 1
	count1, count2 := 0, 0
	for i := low; i < high; i++ {
		candidate := strconv.Itoa(i)
		if isOk(candidate) {
			count1++
		}
		if isOk2(candidate) {
			count2++
		}
	}
	fmt.Println("part I:", count1)
	fmt.Println("part II:", count2)
}
