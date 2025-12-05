package day_5

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"
)

func getInput(filename string) ([]idRange, []int) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, nil
	}
	defer f.Close()

	var idRanges []idRange
	var ingredientsIds []int

	readingRanges := true
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			readingRanges = false
		}

		if readingRanges {
			g := strings.Split(line, "-")
			start, _ := strconv.Atoi(g[0])
			end, _ := strconv.Atoi(g[1])
			idRanges = append(idRanges, idRange{start, end})
		} else {
			id, _ := strconv.Atoi(line)
			ingredientsIds = append(ingredientsIds, id)
		}
	}

	return idRanges, ingredientsIds
}

func TestDay5(t *testing.T) {
	idRanges, ingredientsIds := getInput("../inputs/day_5.txt")

	var tests = []struct {
		idRanges       []idRange
		ingredientsIds []int
		expectedResult int
	}{
		{[]idRange{{3, 5}, {10, 14}, {16, 20}, {12, 18}}, []int{1, 5, 8, 11, 17, 32}, 3},
		{idRanges, ingredientsIds, 679},
	}

	for _, tt := range tests {
		t.Run("Day 5", func(t *testing.T) {
			result := day5(tt.idRanges, tt.ingredientsIds)

			if result != tt.expectedResult {
				t.Errorf("Wrong - got %d, expected %d", result, tt.expectedResult)
			}
		})
	}
}

func TestDay5Extra(t *testing.T) {
	idRanges, _ := getInput("../inputs/day_5.txt")

	var tests = []struct {
		idRanges       []idRange
		expectedResult int
	}{
		{[]idRange{{3, 5}, {10, 14}, {16, 20}, {12, 18}}, 14},
		{idRanges, 358155203664116},
	}

	for _, tt := range tests {
		t.Run("Day 5 (Extra)", func(t *testing.T) {
			result := day5Extra(tt.idRanges)

			if result != tt.expectedResult {
				t.Errorf("Wrong - got %d, expected %d", result, tt.expectedResult)
			}
		})
	}
}
