package main

import (
	"everybody-codes/utils"
	"fmt"
	"strings"
)

func main() {
	line1, _ := utils.ReadLine("2025/quest08/input1.txt")
	lines2, _ := utils.ReadLines("2025/quest08/input2.txt")
	lines3, _ := utils.ReadLines("2025/quest08/input3.txt")
	fmt.Println("2025 Quest 08 Solution")
	fmt.Printf("Part 1: %v\n", part1(line1, 32))
	fmt.Printf("Part 2: %v\n", part2(lines2))
	fmt.Printf("Part 3: %v\n", part3(lines3))
}

func part1(lines string, nails int) int {
	delta := nails / 2
	count := 0
	split := strings.Split(lines, ",")
	for i := 0; i < len(split)-1; i++ {
		a, b := utils.StrToInt(split[i]), utils.StrToInt(split[i+1])
		switch {
		case ((a+delta)%nails) == b && ((b+delta)%nails) == a:
			count++
		case a == delta && ((b+delta)%nails) == a:
			count++
		case b == delta && ((a+delta)%nails) == b:
			count++
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
