package main

import (
	"everybody-codes/utils"
	"testing"
)

func TestPart1(t *testing.T) {
	line1, _ := utils.ReadLine("input_test1.txt")
	line2, _ := utils.ReadLine("input1.txt")
	tests := []struct {
		input    string
		nails    int
		expected int
	}{
		{line1, 8, 4},
		{line2, 32, 57},
	}
	for _, test := range tests {
		result := part1(test.input, test.nails)
		if result != test.expected {
			t.Errorf("Result %v not equal to expected %v", result, test.expected)
		}
	}
}

func TestPart2(t *testing.T) {
	line1, _ := utils.ReadLine("input_test2.txt")
	line2, _ := utils.ReadLine("input2.txt")
	tests := []struct {
		input    string
		nails    int
		expected int
	}{
		{line1, 8, 21},
		{line2, 256, 2926304},
	}
	for _, test := range tests {
		result := part2(test.input, test.nails)
		if result != test.expected {
			t.Errorf("Result %v not equal to expected %v", result, test.expected)
		}
	}
}

func TestPart3(t *testing.T) {
	line1, _ := utils.ReadLine("input_test3.txt")
	line2, _ := utils.ReadLine("input3.txt")
	tests := []struct {
		input    string
		nails    int
		expected int
	}{
		{line1, 8, 7},
		{line2, 256, 2797},
	}
	for _, test := range tests {
		result := part3(test.input, test.nails)
		if result != test.expected {
			t.Errorf("Result %v not equal to expected %v", result, test.expected)
		}
	}
}

func BenchmarkPart1(b *testing.B) {
	lines, _ := utils.ReadLine("input1.txt")
	for n := 0; n < b.N; n++ {
		part1(lines, 32)
	}
}

func BenchmarkPart2(b *testing.B) {
	line, _ := utils.ReadLine("input2.txt")
	for n := 0; n < b.N; n++ {
		part2(line, 256)
	}
}

func BenchmarkPart3(b *testing.B) {
	lines, _ := utils.ReadLine("input3.txt")
	for n := 0; n < b.N; n++ {
		part3(lines, 256)
	}
}
