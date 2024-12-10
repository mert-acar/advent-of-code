package common

import "math"

type number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64 | uint | uint8 | uint16 | uint32 | uint64
}

func Abs[num number](x num) num {
	if x < 0 {
		return -x
	}
	return x
}

func Diff[num number](nums []num, n int) []num {
	var difference []num
	for i := 0; i < len(nums)-1; i++ {
		difference = append(difference, nums[i+1]-nums[i])
	}
	if n == 1 {
		return difference
	} else {
		return Diff(difference, n-1)
	}
}

func Sum1D[num number](slice []num) num {
	total := num(0)
	for _, value := range slice {
    total += value
	}
	return total
}

func Sum2D[num number](slice [][]num) num {
	total := num(0)
	for _, row := range slice {
		for _, value := range row {
			total += value
		}
	}
	return total
}

func Pow[num number](n, m num) num {
	return num(math.Pow(float64(n), float64(m)))
}
