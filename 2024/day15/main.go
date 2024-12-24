package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type Point struct {
	x, y int
}

type Warehouse struct {
	grid     [][]rune
	robotPos Point
}

func NewWarehouse(input string) *Warehouse {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]rune, len(lines))
	var robotPos Point

	for y, line := range lines {
		grid[y] = []rune(line)
		if idx := strings.IndexRune(line, '@'); idx != -1 {
			robotPos = Point{idx, y}
		}
	}

	return &Warehouse{
		grid:     grid,
		robotPos: robotPos,
	}
}

func (w *Warehouse) canMove(x, y int) bool {
	return x >= 0 && x < len(w.grid[0]) && y >= 0 && y < len(w.grid) && w.grid[y][x] != '#'
}

func (w *Warehouse) tryMove(dir rune) bool {
	dx, dy := 0, 0
	switch dir {
	case '^':
		dy = -1
	case 'v':
		dy = 1
	case '<':
		dx = -1
	case '>':
		dx = 1
	default:
		return false
	}

	newX, newY := w.robotPos.x+dx, w.robotPos.y+dy

	if !w.canMove(newX, newY) {
		return false
	}

	if w.grid[newY][newX] == '.' {
		w.grid[w.robotPos.y][w.robotPos.x] = '.'
		w.grid[newY][newX] = '@'
		w.robotPos = Point{newX, newY}
		return true
	}

	if w.grid[newY][newX] == 'O' {

		boxes := []Point{{newX, newY}}
		currX, currY := newX+dx, newY+dy

		for w.canMove(currX, currY) {
			if w.grid[currY][currX] == 'O' {
				boxes = append(boxes, Point{currX, currY})
				currX += dx
				currY += dy
			} else {
				break
			}
		}

		if w.canMove(currX, currY) && w.grid[currY][currX] == '.' {

			w.grid[currY][currX] = 'O'
			for i := len(boxes) - 1; i >= 0; i-- {
				box := boxes[i]
				if i == 0 {

					w.grid[box.y][box.x] = '@'
				} else {

					w.grid[box.y][box.x] = 'O'
				}
			}

			w.grid[w.robotPos.y][w.robotPos.x] = '.'
			w.robotPos = Point{newX, newY}
			return true
		}
	}

	return false
}

func (w *Warehouse) processMovements(moves string) {
	moves = strings.ReplaceAll(moves, "\n", "")
	moves = strings.ReplaceAll(moves, " ", "")
	for _, move := range moves {
		w.tryMove(move)
	}
}

func (w *Warehouse) calculateGPS() int {
	sum := 0
	for y := 0; y < len(w.grid); y++ {
		for x := 0; x < len(w.grid[y]); x++ {
			if w.grid[y][x] == 'O' {
				gps := y*100 + x
				sum += gps
			}
		}
	}
	return sum
}

func (w *Warehouse) String() string {
	var sb strings.Builder
	for _, row := range w.grid {
		sb.WriteString(string(row))
		sb.WriteString("\n")
	}
	return sb.String()
}

func main() {
	bytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read the file: %v\n", err)
	}

	parts := strings.Split(string(bytes), "\n\n")

	warehouse := NewWarehouse(parts[0])
	warehouse.processMovements(parts[1])
	fmt.Printf("Part I: %d\n", warehouse.calculateGPS())
}
