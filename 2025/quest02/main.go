package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("2025/quest02/input1.txt")
	fmt.Println("2025 Quest 02 Solution")
	fmt.Printf("Part 1: %v\n", part1(input))
}

func part1(input []byte) string {
	replacer := strings.NewReplacer("A", "", "=", "", "[", "", ",", " ", "]", "")
	split := strings.Fields(replacer.Replace(string(input)))
	a := Complex{StrToInt(split[0]), StrToInt(split[1])}
	b := Complex{10, 10}
	result := Complex{0, 0}
	for i := 0; i < 3; i++ {
		result.Multiply(&result)
		result.Divide(&b)
		result.Add(&a)
	}
	return result.ToString()
}

type Complex struct {
	X, Y int32
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

func (a *Complex) ToString() string {
	return fmt.Sprintf("[%v,%v]", a.X, a.Y)
}

func StrToInt(s string) int32 {
	num, _ := strconv.Atoi(s)
	return int32(num)
}
