package utils

import (
	"bufio"
	"os"
	"strconv"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func ReadLine(path string) (string, error) {
	lines, err := ReadLines(path)
	if err != nil {
		return "", err
	}
	return lines[0], nil
}

func StrToInt(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

func StrToInt64(s string) int64 {
	num, _ := strconv.ParseInt(s, 10, 64)
	return num
}

func StrToFloat(s string) float64 {
	num, _ := strconv.ParseFloat(s, 64)
	return num
}
