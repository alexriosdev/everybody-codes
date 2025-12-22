package main

import (
	"everybody-codes/utils"
	"fmt"
)

func main() {
	lines1, _ := utils.ReadLines("2025/quest14/input1.txt")
	lines2, _ := utils.ReadLines("2025/quest14/input2.txt")
	lines3, _ := utils.ReadLines("2025/quest14/input3.txt")
	fmt.Println("2025 Quest 14 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines1, 10))
	fmt.Printf("Part 2: %v\n", part2(lines2))
	fmt.Printf("Part 3: %v\n", part3(lines3))
}

func part1(lines []string, rounds int) int {
	grid := NewGrid(lines)
	dirs := []Coordinate{{-1, -1}, {-1, 1}, {1, 1}, {1, -1}}
	sum := 0
	for i := 0; i < rounds; i++ {
		nextActive := []Coordinate{}
		nextInactive := []Coordinate{}
		for r, row := range *grid {
			for c, val := range row {
				count := 0
				curr := Coordinate{r, c}
				for _, dir := range dirs {
					next := Coordinate{curr.Y + dir.Y, curr.X + dir.X}
					if !grid.IsInBounds(next) {
						continue
					}
					if grid.Equals(next, '#') {
						count++
					}
				}
				if (val == '#' && count%2 != 0) || (val == '.' && count%2 == 0) {
					nextActive = append(nextActive, curr)
					continue
				}
				nextInactive = append(nextInactive, curr)
			}
		}
		for _, next := range nextActive {
			grid.Update(next, '#')
		}
		for _, next := range nextInactive {
			grid.Update(next, '.')
		}
		sum += len(nextActive)
	}
	return sum
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

func (g *Grid) Update(c Coordinate, val rune) {
	(*g)[c.Y][c.X] = val
}

func (g *Grid) Equals(c Coordinate, val rune) bool {
	return (*g)[c.Y][c.X] == val
}

func (g *Grid) IsInBounds(c Coordinate) bool {
	return 0 <= c.Y && c.Y < len(*g) && 0 <= c.X && c.X < len((*g)[0])
}

func (g *Grid) Display() {
	for _, row := range *g {
		fmt.Println(string(row))
	}
	fmt.Println()
}
