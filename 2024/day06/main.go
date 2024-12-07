package main

import (
	"aoc.themacar.com/common"
	"fmt"
	"strconv"
	"strings"
)

func getStartPos(vals []string) [2]int {
	guard := "^"
	i, j := 0, 0
	for i = 0; i < len(vals); i++ {
		if strings.Contains(vals[i], guard) {
			j = strings.Index(vals[i], guard)
			runes := []rune(vals[i])
			runes[j] = '.'
			vals[i] = string(runes)
			break
		}
	}
	return [2]int{i, j}
}

func getGrid(n, m int) [][]int {
	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, m)
	}
	return grid
}

func replaceCharecter(vals []string, ch rune, i, j int) []string {
	out := make([]string, len(vals))
	for k := range vals {
		if k == i {
			runes := []rune(vals[k])
			runes[j] = ch
			out[k] = string(runes)
		} else {
			out[k] = vals[k]
		}
	}
	return out
}

func moveGuard(vals []string, i, j int) int {
	n, m := len(vals), len(vals[0])
	grid := getGrid(n, m)
	dx, dy := -1, 0
	var next string
	for i > 0 && i < n-1 && j > 0 && j < m-1 {
		grid[i][j] = 1
		next = string(vals[i+dx][j+dy])
		if next == "#" {
			dx, dy = dy, -dx
		}
		i += dx
		j += dy
	}
	grid[i][j] = 1
	return common.Sum2D(grid)
}

func trapGuard(vals []string, sx, sy int) int {
	n, m := len(vals), len(vals[0])
	sdx, sdy := -1, 0
  total := 0
  var pvals []string
  var x, y int
  for i := 0; i < n; i++{
    for j := 0; j < m; j++{
      if vals[i][j] != '.' {
        continue
      }
      pvals = replaceCharecter(vals, '#', i, j)
      dx, dy := sdx, sdy
      x, y = sx, sy
      h := map[string]bool{}
      for x+dx >= 0 && x+dx < n && y+dy >= 0 && y+dy < m {
        if pvals[x+dx][y+dy] == '#' {
          key := strconv.Itoa(x,) + ":" + strconv.Itoa(y,) + ":" + strconv.Itoa(dx,) + strconv.Itoa(dy,)
					_, ok := h[key]
					if ok {
						total++
						break
					}
					h[key] = true  
          dx, dy = dy, -dx
          continue
        }
        x, y = x+dx, y+dy
      }
    }
  }
  return total
}

func main() {
	vals, err := common.ReadLinesFromFile("./input.txt")
	if err != nil {
		panic("failed to parse input file")
	}

	startPos := getStartPos(vals)
	i, j := startPos[0], startPos[1]

	// Remove guard symbol
	vals = replaceCharecter(vals, '.', i, j)

	totalVisited  := moveGuard(vals, i, j)
	fmt.Printf("Part I: %d\n", totalVisited)

	obs := trapGuard(vals, i, j)
	fmt.Printf("Part II: %d\n", obs)
}
