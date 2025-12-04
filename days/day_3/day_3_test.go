package day_3

import (
	"bufio"
	"os"
	"testing"
)

func getInput(filename string) []string {
	f, err := os.Open(filename)
	if err != nil {
		return nil
	}
	defer f.Close()

	var lines []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines
}

func TestDay3(t *testing.T) {
	inputRotations := getInput("../inputs/day_3.txt")

	var tests = []struct {
		lines          []string
		expectedResult int
	}{
		{[]string{"987654321111111", "811111111111119", "234234234234278", "818181911112111"}, 357},
		{inputRotations, 17244},
	}

	for _, tt := range tests {
		t.Run("Day 3", func(t *testing.T) {
			result := day3(tt.lines, 2)

			if result != tt.expectedResult {
				t.Errorf("Wrong - got %d, expected %d", result, tt.expectedResult)
			}
		})
	}
}

func TestDay3Extra(t *testing.T) {
	inputRotations := getInput("../inputs/day_3.txt")

	var tests = []struct {
		lines          []string
		expectedResult int
	}{
		{[]string{"987654321111111", "811111111111119", "234234234234278", "818181911112111"}, 3121910778619},
		{inputRotations, 171435596092638},
	}

	for _, tt := range tests {
		t.Run("Day 3 (Extra)", func(t *testing.T) {
			result := day3Extra(tt.lines, 12)

			if result != tt.expectedResult {
				t.Errorf("Wrong - got %d, expected %d", result, tt.expectedResult)
			}
		})
	}
}
