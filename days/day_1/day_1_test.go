package day_1

import (
	"bufio"
	"os"
	"strconv"
	"testing"
)

func getInput(filename string) []rotation {
	f, err := os.Open(filename)
	if err != nil {
		return nil
	}
	defer f.Close()

	var rotations []rotation

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		lineDirection := direction(line[0])
		lineNumber, _ := strconv.Atoi(line[1:])
		rotations = append(rotations, rotation{lineDirection, lineNumber})
	}

	return rotations
}

func TestDay1(t *testing.T) {
	inputRotations := getInput("../inputs/day_1.txt")

	var tests = []struct {
		rotations      []rotation
		expectedResult int
	}{
		{[]rotation{{left, 68}, {left, 30}, {right, 48}, {left, 5}, {right, 60}, {left, 55}, {left, 1}, {left, 99}, {right, 14}, {left, 82}}, 3},
		{inputRotations, 1132},
	}

	for _, tt := range tests {
		t.Run("Day 1", func(t *testing.T) {
			result := day1(tt.rotations)

			if result != tt.expectedResult {
				t.Errorf("Wrong - got %d, expected %d", result, tt.expectedResult)
			}
		})
	}
}

func TestDay1Extra(t *testing.T) {
	inputRotations := getInput("../inputs/day_1.txt")

	var tests = []struct {
		rotations      []rotation
		expectedResult int
	}{
		{[]rotation{{left, 68}, {left, 30}, {right, 48}, {left, 5}, {right, 60}, {left, 55}, {left, 1}, {left, 99}, {right, 14}, {left, 82}}, 6},
		{inputRotations, 6623},
	}

	for _, tt := range tests {
		t.Run("Day 1 (Extra)", func(t *testing.T) {
			result := day1Extra(tt.rotations)

			if result != tt.expectedResult {
				t.Errorf("Wrong - got %d, expected %d", result, tt.expectedResult)
			}
		})
	}
}
