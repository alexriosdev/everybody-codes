package main

import (
	"container/heap"
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
	grid.Update(volcano, '0')
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
			sum += int((*grid)[curr.Y][curr.X] - '0')
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
	grid.Update(volcano, '0')
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
			sum += int((*grid)[curr.Y][curr.X] - '0')
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
	grid := NewGrid(lines)
	volcano := grid.Get('@')
	start := grid.Get('S')
	grid.Update(volcano, '0')
	grid.Update(start, '0')
	for rad := 0; rad <= grid.Rows()/2; rad++ {
		end := Coordinate{volcano.Y + rad + 1, volcano.X}
		leftEnd := Coordinate{end.Y, end.X - 1}
		rightEnd := Coordinate{end.Y, end.X + 1}
		minTime := (rad + 1) * 30
		leftDist := grid.FindShortestDistance(start, leftEnd, volcano, rad)
		rightDist := grid.FindShortestDistance(start, rightEnd, volcano, rad)
		sum := leftDist + rightDist + int((*grid)[end.Y][end.X]-'0')
		if sum < minTime {
			return rad * sum
		}
	}
	return 0
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

func (g *Grid) FindShortestDistance(start, end, volcano Coordinate, rad int) int {
	dirs := []Coordinate{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{Pos: start, Dist: 0})
	dist := map[Coordinate]int{start: 0}
	for pq.Len() > 0 {
		curr := heap.Pop(pq).(*Item)
		if curr.Pos == end {
			return curr.Dist
		}
		for _, dir := range dirs {
			next := Coordinate{curr.Pos.Y + dir.Y, curr.Pos.X + dir.X}
			if !g.IsInBounds(next) {
				continue
			}
			if IsInRadius(volcano, next, rad) {
				continue
			}
			if next.X == volcano.X && next.Y > volcano.Y {
				continue
			}
			nextDist := curr.Dist + int((*g)[next.Y][next.X]-'0')
			if prevDist, ok := dist[next]; !ok && (nextDist < prevDist || prevDist == 0) {
				dist[next] = nextDist
				heap.Push(pq, &Item{Pos: next, Dist: nextDist})
			}
		}
	}
	return -1
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
