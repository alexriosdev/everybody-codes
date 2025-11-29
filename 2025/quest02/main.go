package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input1, _ := os.ReadFile("2025/quest02/input1.txt")
	input2, _ := os.ReadFile("2025/quest02/input2.txt")
	fmt.Println("2025 Quest 02 Solution")
	fmt.Printf("Part 1: %v\n", part1(input1))
	fmt.Printf("Part 2: %v\n", part2(input2))
}

func part1(input []byte) string {
	replacer := strings.NewReplacer("A", "", "=", "", "[", "", ",", " ", "]", "")
	split := strings.Fields(replacer.Replace(string(input)))
	addend := Complex{StrToInt64(split[0]), StrToInt64(split[1])}
	divisor := Complex{10, 10}
	result := Complex{0, 0}
	result.Cycle(addend, divisor, 3)
	return result.ToString()
}

func part2(input []byte) int {
	replacer := strings.NewReplacer("A", "", "=", "", "[", "", ",", " ", "]", "")
	split := strings.Fields(replacer.Replace(string(input)))
	start := Complex{StrToInt64(split[0]), StrToInt64(split[1])}
	end := Complex{start.X + 1_000, start.Y + 1_000}
	count := 0
	for x := start.X; x <= end.X; x += 10 {
		for y := start.Y; y <= end.Y; y += 10 {
			addend := Complex{x, y}
			divisor := Complex{100_000, 100_000}
			result := Complex{0, 0}
			if result.Cycle(addend, divisor, 100) {
				count++
			}
		}
	}
	return count
}

type Complex struct {
	X, Y int64
}

func (a *Complex) Cycle(addend, divisor Complex, n int) bool {
	for i := 0; i < n; i++ {
		a.Multiply(a)
		a.Divide(&divisor)
		a.Add(&addend)
		if !a.IsEngraved() {
			return false
		}
	}
	return true
}

func (a *Complex) Add(b *Complex) {
	a.X += b.X
	a.Y += b.Y
}

func (a *Complex) Multiply(b *Complex) {
	X := a.X*b.X - a.Y*b.Y
	Y := a.X*b.Y + a.Y*b.X
	a.X, a.Y = X, Y
}

func (a *Complex) Divide(b *Complex) {
	a.X /= b.X
	a.Y /= b.Y
}

func (a *Complex) IsEngraved() bool {
	return (-1_000_000 <= a.X && a.X <= 1_000_000) && (-1_000_000 <= a.Y && a.Y <= 1_000_000)
}

func (a *Complex) ToString() string {
	return fmt.Sprintf("[%v,%v]", a.X, a.Y)
}

func StrToInt64(s string) int64 {
	num, _ := strconv.Atoi(s)
	return int64(num)
}
