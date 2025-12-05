package day_5

import (
	"math"
	"sort"
)

type idRange struct {
	start int
	end   int
}

func day5(idRanges []idRange, ingredientIds []int) int {
	freshIngredients := make(map[int]bool)

	for _, ingredientId := range ingredientIds {
		for _, id := range idRanges {
			if ingredientId >= id.start && ingredientId <= id.end {
				freshIngredients[ingredientId] = true
			}
		}
	}

	return len(freshIngredients)
}

func day5Extra(idRanges []idRange) int {
	sort.Slice(idRanges, func(i, j int) bool {
		if idRanges[i].start == idRanges[j].start {
			return idRanges[i].end < idRanges[j].end
		}
		return idRanges[i].start < idRanges[j].start
	})

	total := 0
	curStart, curEnd := idRanges[0].start, idRanges[0].end

	for _, r := range idRanges {
		if r.start > curEnd+1 {
			total += curEnd - curStart + 1
			curStart, curEnd = r.start, r.end
		} else if r.end > curEnd {
			curEnd = r.end
		}
	}

	if curStart != math.MinInt {
		total += curEnd - curStart + 1
	}

	return total
}
