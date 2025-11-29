package main

import (
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	lines1, _ := os.ReadFile("input_test1.txt")
	lines2, _ := os.ReadFile("input1.txt")
	tests := []struct {
		input    []byte
		expected string
	}{
		{lines1, "[357,862]"},
		{lines2, "[321579,827750]"},
	}
	for _, test := range tests {
		result := part1(test.input)
		if result != test.expected {
			t.Errorf("Result %v not equal to expected %v", result, test.expected)
		}
	}
}

func TestPart2(t *testing.T) {
	lines1, _ := os.ReadFile("input_test2.txt")
	lines2, _ := os.ReadFile("input2.txt")
	tests := []struct {
		input    []byte
		expected int
	}{
		{lines1, 4076},
		{lines2, 1121},
	}
	for _, test := range tests {
		result := part2(test.input)
		if result != test.expected {
			t.Errorf("Result %v not equal to expected %v", result, test.expected)
		}
	}
}

func BenchmarkPart1(b *testing.B) {
	lines, _ := os.ReadFile("input1.txt")
	for n := 0; n < b.N; n++ {
		part1(lines)
	}
}

func BenchmarkPart2(b *testing.B) {
	lines, _ := os.ReadFile("input2.txt")
	for n := 0; n < b.N; n++ {
		part2(lines)
	}
}
