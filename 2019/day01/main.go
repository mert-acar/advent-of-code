package main

import (
	"fmt"
	"strconv"

	"aoc.themacar.com/common"
)

func calculateFuel(mass int) int {
	return (mass / 3) - 2
}

func calculateTotalFuel(mass int) int {
  t := 0
  for fuel := calculateFuel(mass); fuel >= 0; fuel = calculateFuel(fuel) {
    t += fuel
  }
  return t
}

func main() {
	vals, err := common.ReadLinesFromFile("./input.txt")
	if err != nil {
		panic(err)
	}

	var total1 int
	var total2 int
	for _, v := range vals {
		x, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		total1 += calculateFuel(x)
		total2 += calculateTotalFuel(x)
	}

	fmt.Printf("part I: %d\n", total1)
	fmt.Printf("part II: %d\n", total2)
}
