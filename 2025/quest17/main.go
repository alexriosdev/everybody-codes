package main

import (
	"everybody-codes/utils"
	"fmt"
	"math"
)

func main() {
	lines1, _ := utils.ReadLines("2025/quest17/input1.txt")
	lines2, _ := utils.ReadLines("2025/quest17/input2.txt")
	lines3, _ := utils.ReadLines("2025/quest17/input3.txt")
	fmt.Println("2025 Quest 17 Solution")
	fmt.Printf("Part 1: %v\n", part1(lines1))
	fmt.Printf("Part 2: %v\n", part2(lines2))
	fmt.Printf("Part 3: %v\n", part3(lines3))
}

func part1(lines []string) int {
	dirs := []Coordinate{{-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}}
	radius := 10
	grid := NewGrid(lines)
	volcano := grid.Get('@')
	queue := []Coordinate{volcano}
	visited := map[Coordinate]bool{}
	rad := 0
	sum := 0
	for len(queue) > 0 && rad <= radius {
		size := len(queue)
		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]
			if visited[curr] {
				continue
			}
			visited[curr] = true
			if curr != volcano {
				sum += int((*grid)[curr.Y][curr.X] - '0')
			}
			for _, dir := range dirs {
				next := Coordinate{curr.Y + dir.Y, curr.X + dir.X}
				if !grid.IsInBounds(next) {
					continue
				}
				if IsInRadius(volcano, next, rad+1) {
					queue = append(queue, next)
				}
			}
		}
		rad++
	}
	return sum
}

func part2(lines []string) int {
	dirs := []Coordinate{{-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}}
	grid := NewGrid(lines)
	volcano := grid.Get('@')
	queue := []Coordinate{volcano}
	visited := map[Coordinate]bool{}
	maxSum := math.MinInt
	result := 0
	rad := 0
	for len(queue) > 0 {
		sum := 0
		size := len(queue)
		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]
			if visited[curr] {
				continue
			}
			visited[curr] = true
			if curr != volcano {
				sum += int((*grid)[curr.Y][curr.X] - '0')
			}
			for _, dir := range dirs {
				next := Coordinate{curr.Y + dir.Y, curr.X + dir.X}
				if !grid.IsInBounds(next) {
					continue
				}
				if IsInRadius(volcano, next, rad+1) {
					queue = append(queue, next)
				}
			}
		}
		if sum > maxSum {
			maxSum = sum
			result = rad * sum
		}
		rad++
	}
	return result
}

func part3(lines []string) int {
	return len(lines)
}

type Coordinate struct {
	Y, X int
}

func IsInRadius(a, b Coordinate, radius int) bool {
	return (a.X-b.X)*(a.X-b.X)+(a.Y-b.Y)*(a.Y-b.Y) <= radius*radius
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

func (g *Grid) GetAll(target rune) []Coordinate {
	result := []Coordinate{}
	for r, row := range *g {
		for c, val := range row {
			if val == target {
				result = append(result, Coordinate{r, c})
			}
		}
	}
	return result
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
