package main

import (
	"everybody-codes/utils"
	"testing"
)

func TestPart1(t *testing.T) {
	lines1, _ := utils.ReadLines("input_test1.txt")
	lines2, _ := utils.ReadLines("input1.txt")
	tests := []struct {
		input    []string
		moves    int
		expected int
	}{
		{lines1, 3, 27},
		{lines2, 4, 147},
	}
	for _, test := range tests {
		result := part1(test.input, test.moves)
		if result != test.expected {
			t.Errorf("Result %v not equal to expected %v", result, test.expected)
		}
	}
}

func TestPart2(t *testing.T) {
	lines1, _ := utils.ReadLines("input_test2.txt")
	lines2, _ := utils.ReadLines("input2.txt")
	tests := []struct {
		input    []string
		moves    int
		expected int
	}{
		{lines1, 3, 27},
		{lines2, 20, 1669},
	}
	for _, test := range tests {
		result := part2(test.input, test.moves)
		if result != test.expected {
			t.Errorf("Result %v not equal to expected %v", result, test.expected)
		}
	}
}

func TestPart3(t *testing.T) {
	lines1, _ := utils.ReadLines("input_test3.txt")
	lines2, _ := utils.ReadLines("input_test4.txt")
	lines3, _ := utils.ReadLines("input_test5.txt")
	lines4, _ := utils.ReadLines("input_test6.txt")
	lines5, _ := utils.ReadLines("input_test7.txt")
	lines6, _ := utils.ReadLines("input3.txt")
	tests := []struct {
		input    []string
		expected int
	}{
		{lines1, 15},
		{lines2, 8},
		{lines3, 44},
		{lines4, 4406},
		{lines5, 13033988838},
		{lines6, 21622707167488},
	}
	for _, test := range tests {
		result := part3(test.input)
		if result != test.expected {
			t.Errorf("Result %v not equal to expected %v", result, test.expected)
		}
	}
}

func BenchmarkPart1(b *testing.B) {
	lines, _ := utils.ReadLines("input1.txt")
	for n := 0; n < b.N; n++ {
		part1(lines, 4)
	}
}

func BenchmarkPart2(b *testing.B) {
	lines, _ := utils.ReadLines("input2.txt")
	for n := 0; n < b.N; n++ {
		part2(lines, 20)
	}
}

func BenchmarkPart3(b *testing.B) {
	lines, _ := utils.ReadLines("input3.txt")
	for n := 0; n < b.N; n++ {
		part3(lines)
	}
}
