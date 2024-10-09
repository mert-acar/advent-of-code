package main

import (
	"fmt"
	"strconv"
	"strings"

	"aoc.themacar.com/common"
)

func parseInput(filename string) []int {
  vals, err := common.ReadLinesFromFile(filename)
  if err != nil {
    panic(err)
  }
  strVals := strings.Split(strings.TrimSpace(string(vals[0])), ",") 
  intVals := make([]int, len(strVals))
  for i, v := range strVals {
    val, err := strconv.Atoi(v)
    if err != nil {
      panic(err)
    }
    intVals[i] = val
  }
  return intVals
}

func runIntcode (program []int) int {
  for ptr := 0; program[ptr] != 99; ptr += 4 {
    optcode, a, b, out := program[ptr], program[ptr + 1], program[ptr +  2], program[ptr + 3]
    switch optcode{
    case 1:
      program[out] = program[a] + program[b]
    case 2:
      program[out] = program[a] * program[b]
    }
  }
  return program[0]
}

func findNounAndVerb(program []int, target int) (int, error) {
  original := make([]int, len(program))
  copy(original, program)

  for noun := 0; noun < 100; noun++ {
    for verb := 0; verb < 100; verb++{
      copy(program, original)
      program[1] = noun
      program[2] = verb
      if runIntcode(program) == target {
        return 100 * noun + verb, nil
      }
    }
  }
  return 0, fmt.Errorf("no combination satisfies")
}


func main () {
  program := parseInput("./input.txt")   
  part1Vals := make([]int, len(program))
  copy(part1Vals, program)
  part1Vals[1] = 12
  part1Vals[2] = 2
  result := runIntcode(part1Vals)
  fmt.Printf("part I: %d\n", result)

  const target = 19690720
  ans, err := findNounAndVerb(program, target)
  if err != nil {
    panic(err)
  }
  fmt.Printf("part II: %d\n", ans)
}
