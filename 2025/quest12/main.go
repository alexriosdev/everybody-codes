package main

import (
	"everybody-codes/utils"
	"fmt"
)

func main() {
	lines1, _ := utils.ReadLines("2025/quest12/input1.txt")
	lines2, _ := utils.ReadLines("2025/quest12/input2.txt")
	lines3, _ := utils.ReadLines("2025/quest12/input3.txt")
	fmt.Println("2025 Quest 12 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines1))
	fmt.Printf("Part 2: %v\n", part2(lines2))
	fmt.Printf("Part 3: %v\n", part3(lines3))
}

func part1(lines []string) int {
	grid := NewGrid(lines)
	dirs := []Coordinate{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	start := Coordinate{0, 0}
	queue := []Coordinate{start}
	visited := map[Coordinate]bool{}
	count := 0
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]
			if visited[curr] {
				continue
			}
			visited[curr] = true
			count++
			for _, dir := range dirs {
				next := Coordinate{curr.Y + dir.Y, curr.X + dir.X}
				if !grid.IsInBounds(next) {
					continue
				}
				if grid.IsGreaterOrEqual(curr, next) {
					queue = append(queue, next)
				}
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

func (g *Grid) IsGreaterOrEqual(a, b Coordinate) bool {
	return (*g)[a.Y][a.X] >= (*g)[b.Y][b.X]
}

func (g *Grid) IsInBounds(c Coordinate) bool {
	return 0 <= c.Y && c.Y < len(*g) && 0 <= c.X && c.X < len((*g)[0])
}
