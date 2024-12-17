package main

import (
	"aoc.themacar.com/common"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Queue struct {
	vals [][]int
}

func (q *Queue) dequeue() []int {
	p := q.vals[0]
	q.vals = q.vals[1:]
	return p
}

func (q *Queue) enqueue(elem []int) {
	q.vals = append(q.vals, elem)
}

func (q *Queue) isEmpty() bool {
	return len(q.vals) == 0
}

func parseInput(bytes []uint8) (a, b, c int, program []int) {
	input := strings.Split(string(bytes), "\n")
	a, _ = strconv.Atoi(strings.Split(input[0], ": ")[1])
	b, _ = strconv.Atoi(strings.Split(input[1], ": ")[1])
	c, _ = strconv.Atoi(strings.Split(input[2], ": ")[1])
	str := strings.Split(strings.Split(input[4], ": ")[1], ",")
	for _, p := range str {
		n, _ := strconv.Atoi(p)
		program = append(program, n)
	}
	return a, b, c, program
}

func parseOut(numbers []int) (out string) {
	var strSlice []string
	for _, num := range numbers {
		strSlice = append(strSlice, strconv.Itoa(num))
	}
	out = strings.Join(strSlice, ",")
	return
}

func parseOperand(operand, a, b, c int) int {
	if operand <= 3 {
		return operand
	} else if operand == 4 {
		return a
	} else if operand == 5 {
		return b
	} else if operand == 6 {
		return c
	} else if operand == 7 {
		panic("operand is invalid")
	}
	return -1
}

func runProgram(A, B, C int, opcodes []int) (out []int) {
	ptr := 0
	for ptr < len(opcodes) {
		inst, operand := opcodes[ptr], opcodes[ptr+1]
		opt := parseOperand(operand, A, B, C)
		switch inst {
		case 0: // division
			A /= common.Pow(2, opt)
		case 1: // bitwise XOR
			B = B ^ operand
		case 2: // modulo 8
			B = opt % 8
		case 3: // jump if A != 0
			if A != 0 {
				ptr = operand
				continue
			}
		case 4: // legacy XOR
			B = B ^ C
		case 5: // out modulo 8
			out = append(out, opt%8)
		case 6:
			B = A / common.Pow(2, opt)
		case 7:
			C = A / common.Pow(2, opt)
		default:
			continue
		}
		ptr += 2
	}
	return
}

func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func reverseEngineer(A, B, C int, opcodes []int) int {
	candidates := Queue{vals: [][]int{{len(opcodes) - 1, 0}}}
	for !candidates.isEmpty() {
		curr := candidates.dequeue()
		for i := 0; i < 8; i++ {
			nA := (curr[1] * 8) + i
			out := runProgram(nA, B, C, opcodes)
			ok := true
			for j := curr[0]; j < len(opcodes); j++ {
				if opcodes[j] != out[j-curr[0]] {
					ok = false
					break
				}
			}
			if ok {
				fmt.Println(out, curr[0])
				if curr[0] == 0 {
					return nA
				}
				candidates.enqueue([]int{curr[0] - 1, nA})
			}
		}
	}
	return -1
}

func main() {
	bytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading the input file: %v\n", err)
		os.Exit(1)
	}
	A, B, C, opcodes := parseInput(bytes)

	fmt.Printf("Part I: %s\n", parseOut(runProgram(A, B, C, opcodes)))
	fmt.Printf("Part II: %d\n", reverseEngineer(A, B, C, opcodes))
}
