package main

import (
	"everybody-codes/utils"
	"fmt"
)

func main() {
	lines1, _ := utils.ReadLines("2025/quest10/input1.txt")
	lines2, _ := utils.ReadLines("2025/quest10/input2.txt")
	lines3, _ := utils.ReadLines("2025/quest10/input3.txt")
	fmt.Println("2025 Quest 10 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines1, 3))
	fmt.Printf("Part 2: %v\n", part2(lines2))
	fmt.Printf("Part 3: %v\n", part3(lines3))
}

func part1(lines []string, moves int) int {
	dirs := [][]int{{-2, 1}, {-2, -1}, {1, 2}, {-1, 2}, {2, 1}, {2, -1}, {1, -2}, {-1, -2}}
	grid := NewGrid(lines)
	start := grid.Get('D')
	queue := []Coordinate{start}
	visited := map[Coordinate]bool{}
	count := 0
	for ; moves > 0; moves-- {
		size := len(queue)
		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]
			if visited[curr] {
				continue
			}
			visited[curr] = true
			for _, dir := range dirs {
				r := curr.Y + dir[0]
				c := curr.X + dir[1]
				next := Coordinate{r, c}
				if !grid.IsInBounds(next) {
					continue
				}
				if grid.Contains(next, 'S') {
					count++
				}
				grid.Update(next, 'X')
				queue = append(queue, next)
			}
		}
	}
	return count
}

func part2(lines []string) int {
	return len(lines)
}

func part3(lines []string) int {
	return len(lines)
}

type Coordinate struct {
	Y, X int
}

type Grid [][]rune

func NewGrid(lines []string) *Grid {
	grid := &Grid{}
	for _, line := range lines {
		*grid = append(*grid, []rune(line))
	}
	return grid
}

func (g *Grid) Get(target rune) Coordinate {
	for r, row := range *g {
		for c, val := range row {
			if val == target {
				return Coordinate{r, c}
			}
		}
	}
	return Coordinate{-1, -1}
}

func (g *Grid) Update(c Coordinate, val rune) {
	(*g)[c.Y][c.X] = val
}

func (g *Grid) Contains(c Coordinate, val rune) bool {
	return (*g)[c.Y][c.X] == val
}

func (g *Grid) IsInBounds(c Coordinate) bool {
	return 0 <= c.Y && c.Y < len(*g) && 0 <= c.X && c.X < len((*g)[0])
}
