package main

import (
	"everybody-codes/utils"
	"fmt"
	"maps"
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
	start := Coordinate{0, 0}
	queue := []Coordinate{start}
	visited := map[Coordinate]bool{}
	return len(igniteBarrels(&queue, grid, &visited))
}

func part2(lines []string) int {
	grid := NewGrid(lines)
	start, end := Coordinate{0, 0}, Coordinate{grid.Rows() - 1, grid.Cols() - 1}
	queue := []Coordinate{start, end}
	visited := map[Coordinate]bool{}
	return len(igniteBarrels(&queue, grid, &visited))
}

func part3(lines []string) int {
	grid := NewGrid(lines)
	visited := map[Coordinate]bool{}
	for i := 0; i < 3; i++ {
		maxBarrels := map[Coordinate]bool{}
		for r, row := range *grid {
			for c := range row {
				queue := []Coordinate{{r, c}}
				newVisited := maps.Clone(visited)
				barrels := igniteBarrels(&queue, grid, &newVisited)
				if len(barrels) >= len(maxBarrels) {
					maxBarrels = barrels
				}
			}
		}
		maps.Copy(visited, maxBarrels)
	}
	return len(visited)
}

func igniteBarrels(queue *[]Coordinate, grid *Grid, visited *map[Coordinate]bool) map[Coordinate]bool {
	dirs := []Coordinate{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	for len(*queue) > 0 {
		curr := (*queue)[0]
		(*queue) = (*queue)[1:]
		if (*visited)[curr] {
			continue
		}
		(*visited)[curr] = true
		for _, dir := range dirs {
			next := Coordinate{curr.Y + dir.Y, curr.X + dir.X}
			if !grid.IsInBounds(next) {
				continue
			}
			if grid.IsGreaterOrEqual(curr, next) {
				*queue = append(*queue, next)
			}
		}
	}
	return *visited
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

func (g *Grid) Rows() int {
	return len(*g)
}

func (g *Grid) Cols() int {
	return len((*g)[0])
}

func (g *Grid) IsGreaterOrEqual(a, b Coordinate) bool {
	return (*g)[a.Y][a.X] >= (*g)[b.Y][b.X]
}

func (g *Grid) IsInBounds(c Coordinate) bool {
	return 0 <= c.Y && c.Y < g.Rows() && 0 <= c.X && c.X < g.Cols()
}
