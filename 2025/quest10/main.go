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
	fmt.Printf("Part 1: %v\n", part1(lines1, 4))
	fmt.Printf("Part 2: %v\n", part2(lines2, 20))
	fmt.Printf("Part 3: %v\n", part3(lines3))
}

func part1(lines []string, moves int) int {
	dirs := []Coordinate{{-2, 1}, {-2, -1}, {1, 2}, {-1, 2}, {2, 1}, {2, -1}, {1, -2}, {-1, -2}}
	grid := NewGrid(lines)
	start := grid.Get('D')
	queue := []Coordinate{start}
	count := 0
	for ; moves > 0; moves-- {
		visited := map[Coordinate]bool{}
		size := len(queue)
		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]
			if visited[curr] {
				continue
			}
			visited[curr] = true
			for _, dir := range dirs {
				next := Coordinate{curr.Y + dir.Y, curr.X + dir.X}
				if !grid.IsInBounds(next) {
					continue
				}
				if grid.Equals(next, 'S') {
					count++
				}
				grid.Update(next, 'X')
				queue = append(queue, next)
			}
		}
	}
	return count
}

func part2(lines []string, moves int) int {
	dragonDirs := []Coordinate{{-2, 1}, {-2, -1}, {1, 2}, {-1, 2}, {2, 1}, {2, -1}, {1, -2}, {-1, -2}}
	sheepDir := Coordinate{1, 0}
	grid := NewGrid(lines)
	dragonStart := grid.Get('D')
	dragonQueue := []Coordinate{dragonStart}
	sheepQueue := grid.GetAll('S')
	count := 0
	for ; moves > 0; moves-- {
		dragonVisited := map[Coordinate]bool{}
		huntingSpot := map[Coordinate]bool{}
		dragonSize := len(dragonQueue)
		for i := 0; i < dragonSize; i++ {
			curr := dragonQueue[0]
			dragonQueue = dragonQueue[1:]
			if dragonVisited[curr] {
				continue
			}
			dragonVisited[curr] = true
			for _, dir := range dragonDirs {
				next := Coordinate{curr.Y + dir.Y, curr.X + dir.X}
				if !grid.IsInBounds(next) {
					continue
				}
				if !grid.Equals(next, '#') {
					huntingSpot[next] = true
				}
				dragonQueue = append(dragonQueue, next)
			}
		}
		sheepSize := len(sheepQueue)
		for i := 0; i < sheepSize; i++ {
			curr := sheepQueue[0]
			sheepQueue = sheepQueue[1:]
			if huntingSpot[curr] {
				count++
				continue
			}
			if !grid.Equals(curr, '#') {
				grid.Update(curr, '.')
			}
			next := Coordinate{curr.Y + sheepDir.Y, curr.X + sheepDir.X}
			if !grid.IsInBounds(next) {
				continue
			}
			if huntingSpot[next] {
				count++
				continue
			}
			if !grid.Equals(next, '#') {
				grid.Update(next, 'S')
			}
			sheepQueue = append(sheepQueue, next)
		}
	}
	return count
}

func part3(lines []string) int {
	const (
		DragonTurn = iota
		SheepTurn
	)
	dragonDirs := []Coordinate{{-2, 1}, {-2, -1}, {1, 2}, {-1, 2}, {2, 1}, {2, -1}, {1, -2}, {-1, -2}}
	grid := NewGrid(lines)
	sheep := grid.GetAll('S')
	dragon := grid.Get('D')
	visited := map[string]int{}
	var dfs func(sheep []Coordinate, dragon Coordinate, turn int) int
	dfs = func(sheep []Coordinate, dragon Coordinate, turn int) int {
		if len(sheep) == 0 {
			return 1
		}
		key := fmt.Sprint(sheep, dragon, turn)
		if val, ok := visited[key]; ok {
			return val
		}
		count := 0
		switch turn {
		case SheepTurn:
			canMove := false
			for i, s := range sheep {
				next := Coordinate{s.Y + 1, s.X}
				if next.Y == len(*grid) {
					canMove = true
					continue
				}
				if grid.Equals(next, '#') || next != dragon {
					canMove = true
					newSheep := make([]Coordinate, len(sheep))
					copy(newSheep, sheep)
					newSheep[i] = next
					count += dfs(newSheep, dragon, DragonTurn)
				}
			}
			if !canMove {
				count += dfs(sheep, dragon, DragonTurn)
			}
		case DragonTurn:
			for _, dir := range dragonDirs {
				next := Coordinate{dragon.Y + dir.Y, dragon.X + dir.X}
				if !grid.IsInBounds(next) {
					continue
				}
				newSheep := []Coordinate{}
				for _, s := range sheep {
					if grid.Equals(s, '#') || s != next {
						newSheep = append(newSheep, s)
					}
				}
				count += dfs(newSheep, next, SheepTurn)
			}
		}
		visited[key] = count
		return count
	}
	return dfs(sheep, dragon, SheepTurn)
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
