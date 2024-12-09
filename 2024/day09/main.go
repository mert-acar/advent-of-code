package main

import (
	"fmt"
	"os"
	"strconv"

	"aoc.themacar.com/common"
)

// fileSpan represents a continuous span of a file
type fileSpan struct {
	start int
	end   int
	size  int
}

func calculateChecksum(blocks []string) int {
	checksum := 0
	for i, idx := range blocks {
		if idx != "." {
			val, err := strconv.Atoi(idx)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error converting block to int: %v\n", err)
				continue
			}
			checksum += i * val
		}
	}
	return checksum
}

func canFitFile(blocks []string, start, size int) bool {
	if start+size > len(blocks) {
		return false
	}

	for i := start; i < start+size; i++ {
		if blocks[i] != "." {
			return false
		}
	}
	return true
}

func findBestTargetSpace(blocks []string, fileSize, startPos int) int {
	for i := 0; i < startPos; i++ {
		if canFitFile(blocks, i, fileSize) {
			return i
		}
	}
	return -1
}

func moveFileSpan(blocks []string, span fileSpan, target int) {
	fileBlocks := make([]string, span.size)
	copy(fileBlocks, blocks[span.start:span.end+1])

	for i := span.start; i <= span.end; i++ {
		blocks[i] = "."
	}
	copy(blocks[target:], fileBlocks)
}

func findFileSpans(blocks []string, fileID int) []fileSpan {
	var spans []fileSpan
	var currentSpan *fileSpan
	for i, block := range blocks {
		if block == strconv.Itoa(fileID) {
			if currentSpan == nil {
				currentSpan = &fileSpan{start: i, end: i, size: 1}
			} else {
				currentSpan.end = i
				currentSpan.size++
			}
		} else if currentSpan != nil {
			spans = append(spans, *currentSpan)
			currentSpan = nil
		}
	}

	if currentSpan != nil {
		spans = append(spans, *currentSpan)
	}

	return spans
}

func rearrangeFiles(blocks []string, id int) []string {
	modifiedBlocks := append([]string{}, blocks...)
	for i := id; i >= 0; i-- {
		fileSpans := findFileSpans(modifiedBlocks, i)
		for _, span := range fileSpans {
			bestTarget := findBestTargetSpace(modifiedBlocks, span.size, span.start)
			if bestTarget != -1 {
				moveFileSpan(modifiedBlocks, span, bestTarget)
			}
		}
	}
	return modifiedBlocks
}

func rearrangeFragments(blocks []string, fptr int) []string {
	modifiedBlocks := append([]string{}, blocks...)

	eptr := len(modifiedBlocks) - 1
	for fptr < eptr {
		modifiedBlocks[fptr], modifiedBlocks[eptr] = modifiedBlocks[eptr], modifiedBlocks[fptr]
		for fptr < len(modifiedBlocks) && modifiedBlocks[fptr] != "." {
			fptr++
		}
		eptr--
	}

	return modifiedBlocks
}

func processBlocks(diskMap string) ([]string, int, int) {
	var blocks []string
	fptr, id := 0, 0

	for i := 0; i < len(diskMap); i++ {
		n, err := strconv.Atoi(string(diskMap[i]))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error converting character to int: %v\n", err)
			continue
		}

		if i%2 == 0 {
			for j := 0; j < n; j++ {
				blocks = append(blocks, strconv.Itoa(id))
			}
			id++
		} else {
			if fptr == 0 {
				fptr = len(blocks)
			}
			for j := 0; j < n; j++ {
				blocks = append(blocks, ".")
			}
		}
	}
	return blocks, fptr, id - 1
}

func main() {
	lines, err := common.ReadLinesFromFile("./input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading the input file: %v\n", err)
		os.Exit(1)
	}

	diskMap := lines[0]
	blocks, fptr, id := processBlocks(diskMap)

	// Part I
	modified := rearrangeFragments(blocks, fptr)
	checksum := calculateChecksum(modified)
	fmt.Printf("Part I: %d\n", checksum)

	// Part II
	modified = rearrangeFiles(blocks, id)
	checksum = calculateChecksum(modified)
	fmt.Printf("Part II: %d\n", checksum)
}
