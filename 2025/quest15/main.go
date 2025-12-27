package main

import (
	"container/heap"
	"everybody-codes/utils"
	"fmt"
	"math"
	"sort"
	"strings"
)

func main() {
	line1, _ := utils.ReadLine("2025/quest15/input1.txt")
	line2, _ := utils.ReadLine("2025/quest15/input2.txt")
	line3, _ := utils.ReadLine("2025/quest15/input3.txt")
	fmt.Println("2025 Quest 15 Solution")
	fmt.Printf("Part 1: %v\n", part1(line1))
	fmt.Printf("Part 2: %v\n", part2(line2))
	fmt.Printf("Part 3: %v\n", part3(line3))
}

func part1(line string) int {
	return NavigateMaze(line, false)
}

func part2(line string) int {
	return NavigateMaze(line, false)
}

func part3(line string) int {
	return NavigateMaze(line, true)
}

func NavigateMaze(line string, useCoordinateCompression bool) int {
	const (
		Up = iota
		Right
		Down
		Left
	)
	var Origin = Coordinate{0, 0}
	curr := State{Origin, Up}
	positions := Coordinates{curr.Pos}
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
	if useCoordinateCompression {
		return positions.FindShortestDistance()
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
	grid.BuildWall(positions)
	return grid.FindShortestDistance(positions[0], positions[len(positions)-1])
}

type State struct {
	Pos Coordinate
	Dir int
}

type Coordinate struct {
	Y, X int
}

type Coordinates []Coordinate

type CompressedCoordinates struct {
	OriginalY, OriginalX     []int
	CompressedY, CompressedX map[int]int
}

func (c *Coordinates) Compress() CompressedCoordinates {
	dirs := []Coordinate{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	setY, setX := map[int]bool{}, map[int]bool{}
	for _, coord := range *c {
		for _, dir := range dirs {
			next := Coordinate{coord.Y + dir.Y, coord.X + dir.X}
			setY[next.Y] = true
			setX[next.X] = true
		}
	}
	originalY, originalX := []int{}, []int{}
	for y := range setY {
		originalY = append(originalY, y)
	}
	for x := range setX {
		originalX = append(originalX, x)
	}
	sort.Ints(originalY)
	sort.Ints(originalX)
	compressedY, compressedX := make(map[int]int), make(map[int]int)
	for i, y := range originalY {
		compressedY[y] = i
	}
	for i, x := range originalX {
		compressedX[x] = i
	}
	return CompressedCoordinates{
		OriginalY:   originalY,
		OriginalX:   originalX,
		CompressedY: compressedY,
		CompressedX: compressedX,
	}
}

func (c *Coordinates) BuildWall() map[Coordinate]bool {
	wall := map[Coordinate]bool{}
	curr := (*c)[0]
	for _, pos := range *c {
		if Abs(pos.Y-curr.Y) != 0 {
			for startY, endY := min(pos.Y, curr.Y), max(pos.Y, curr.Y); startY <= endY; startY++ {
				next := Coordinate{startY, pos.X}
				wall[next] = true
			}
		}
		if Abs(pos.X-curr.X) != 0 {
			for startX, endX := min(pos.X, curr.X), max(pos.X, curr.X); startX <= endX; startX++ {
				next := Coordinate{curr.Y, startX}
				wall[next] = true
			}
		}
		curr = pos
	}
	wall[curr] = false
	wall[(*c)[0]] = false
	return wall
}

func (coordinates *Coordinates) FindShortestDistance() int {
	dirs := []Coordinate{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	cc := (*coordinates).Compress()
	for i, pos := range *coordinates {
		(*coordinates)[i].Y = cc.CompressedY[pos.Y]
		(*coordinates)[i].X = cc.CompressedX[pos.X]
	}
	wall := (*coordinates).BuildWall()
	var Start = (*coordinates)[0]
	var End = (*coordinates)[len((*coordinates))-1]
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{Pos: Start, Dist: 0})
	dist := make(map[Coordinate]int)
	dist[Start] = 0
	for pq.Len() > 0 {
		curr := heap.Pop(pq).(*Item)
		if curr.Pos == End {
			return curr.Dist
		}
		if curr.Dist < dist[curr.Pos] {
			continue
		}
		for _, dir := range dirs {
			next := Coordinate{curr.Pos.Y + dir.Y, curr.Pos.X + dir.X}
			if !(0 <= next.Y && next.Y < len(cc.OriginalY) && 0 <= next.X && next.X < len(cc.OriginalX)) {
				continue
			}
			if wall[next] {
				continue
			}
			nextDist := curr.Dist + Abs(cc.OriginalY[curr.Pos.Y]-cc.OriginalY[next.Y]) + Abs(cc.OriginalX[curr.Pos.X]-cc.OriginalX[next.X])
			if prevDist, ok := dist[next]; !ok || nextDist < prevDist {
				dist[next] = nextDist
				heap.Push(pq, &Item{Pos: next, Dist: nextDist})
			}
		}
	}
	return -1
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

func (g *Grid) BuildWall(positions []Coordinate) {
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

func (g *Grid) FindShortestDistance(start, end Coordinate) int {
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
	return -1
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

type Item struct {
	Pos  Coordinate
	Dist int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].Dist < pq[j].Dist }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
