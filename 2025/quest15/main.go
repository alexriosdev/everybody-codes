package main

import (
	"everybody-codes/utils"
	"fmt"
	"math"
	"strings"
)

func main() {
	line1, _ := utils.ReadLine("2025/quest15/input1.txt")
	line2, _ := utils.ReadLine("2025/quest15/input2.txt")
	lines3, _ := utils.ReadLines("2025/quest15/input3.txt")
	fmt.Println("2025 Quest 15 Solution")
	fmt.Printf("Part 1: %v\n", part1(line1))
	fmt.Printf("Part 2: %v\n", part2(line2))
	fmt.Printf("Part 3: %v\n", part3(lines3))
}

func part1(line string) int {
	return NavigateMaze(line)
}

func part2(line string) int {
	return NavigateMaze(line)
}

func part3(lines []string) int {
	return len(lines)
}

func NavigateMaze(line string) int {
	const (
		Up = iota
		Right
		Down
		Left
	)
	var Origin = Coordinate{0, 0}
	curr := State{Origin, Up}
	positions := []Coordinate{curr.Pos}
	for _, s := range strings.Split(line, ",") {
		dir, steps := s[0], utils.StrToInt(s[1:])
		if dir == 'R' {
			curr.Dir = (curr.Dir + 1) % 4
		} else {
			curr.Dir = (curr.Dir + 3) % 4
		}
		switch curr.Dir {
		case Up:
			curr.Pos.Y -= steps
		case Right:
			curr.Pos.X += steps
		case Down:
			curr.Pos.Y += steps
		case Left:
			curr.Pos.X -= steps
		}
		positions = append(positions, curr.Pos)
	}
	minY, maxY := math.MaxInt, math.MinInt
	minX, maxX := math.MaxInt, math.MinInt
	for _, pos := range positions {
		minY, maxY = min(minY, pos.Y), max(maxY, pos.Y)
		minX, maxX = min(minX, pos.X), max(maxX, pos.X)
	}
	rows, cols := Abs(maxY-minY)+1, Abs(maxX-minX)+1
	for i, pos := range positions {
		positions[i].Y = (rows - 1) - (maxY - pos.Y)
		positions[i].X = (cols - 1) - (maxX - pos.X)
	}
	grid := BuildGrid(rows, cols)
	grid.BuildPath(positions)
	return grid.FindShortestPathLength(positions[0], positions[len(positions)-1])
}

type State struct {
	Pos Coordinate
	Dir int
}

type Coordinate struct {
	Y, X int
}

type Grid [][]rune

func BuildGrid(rows, cols int) *Grid {
	grid := &Grid{}
	for r := 0; r < rows; r++ {
		runes := []rune{}
		for c := 0; c < cols; c++ {
			runes = append(runes, ' ')
		}
		*grid = append(*grid, runes)
	}
	return grid
}

func (g *Grid) BuildPath(positions []Coordinate) {
	curr := positions[0]
	for _, pos := range positions {
		if Abs(pos.Y-curr.Y) != 0 {
			for startY, endY := min(pos.Y, curr.Y), max(pos.Y, curr.Y); startY <= endY; startY++ {
				next := Coordinate{startY, pos.X}
				g.Update(next, '#')
			}
		}
		if Abs(pos.X-curr.X) != 0 {
			for startX, endX := min(pos.X, curr.X), max(pos.X, curr.X); startX <= endX; startX++ {
				next := Coordinate{curr.Y, startX}
				g.Update(next, '#')
			}
		}
		curr = pos
	}
	g.Update(curr, 'E')
	g.Update(positions[0], 'S')
}

func (g *Grid) FindShortestPathLength(start, end Coordinate) int {
	dirs := []Coordinate{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	queue := []Coordinate{start}
	visited := map[Coordinate]bool{}
	count := 0
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]
			if curr == end {
				return count
			}
			if visited[curr] {
				continue
			}
			visited[curr] = true
			for _, dir := range dirs {
				next := Coordinate{curr.Y + dir.Y, curr.X + dir.X}
				if !g.IsInBounds(next) {
					continue
				}
				if g.Equals(next, '#') {
					continue
				}
				queue = append(queue, next)
			}
		}
		count++
	}
	return count
}

func (g *Grid) Update(c Coordinate, val rune) {
	(*g)[c.Y][c.X] = val
}

func (g *Grid) Equals(c Coordinate, val rune) bool {
	return (*g)[c.Y][c.X] == val
}

func (g *Grid) Rows() int {
	return len(*g)
}

func (g *Grid) Cols() int {
	return len((*g)[0])
}

func (g *Grid) IsInBounds(c Coordinate) bool {
	return 0 <= c.Y && c.Y < g.Rows() && 0 <= c.X && c.X < g.Cols()
}

func (g *Grid) Display() {
	for _, row := range *g {
		fmt.Println(string(row))
	}
	fmt.Println()
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
