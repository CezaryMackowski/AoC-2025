package day_2

import "strconv"

type identifierRange struct {
	start int
	end   int
}

func (ir identifierRange) startAsString() string {
	return strconv.Itoa(ir.start)
}

func (ir identifierRange) endAsString() string {
	return strconv.Itoa(ir.end)
}

func hasTwoEqualHalves(number int) bool {
	digits := strconv.Itoa(number)
	if len(digits)%2 != 0 {
		return false
	}

	middle := len(digits) / 2
	return digits[:middle] == digits[middle:]
}

func hasRepeatedSequence(number int) bool {
	digits := strconv.Itoa(number)
	length := len(digits)

	for size := 1; size <= length/2; size++ {
		if length%size != 0 {
			continue
		}

		pattern := digits[:size]
		valid := true
		for offset := size; offset < length; offset += size {
			test := digits[offset : offset+size]
			if test != pattern {
				valid = false
				break
			}
		}

		if valid {
			return true
		}
	}

	return false
}

func day2(idRanges []identifierRange) int {
	sumOfInvalidIds := 0

	for _, idRange := range idRanges {
		start, end := idRange.start, idRange.end

		for i := start; i <= end; i++ {
			if hasTwoEqualHalves(i) {
				sumOfInvalidIds += i
			}
		}
	}

	return sumOfInvalidIds
}

func day2Extra(idRanges []identifierRange) int {
	sumOfInvalidIds := 0

	for _, idRange := range idRanges {
		start, end := idRange.start, idRange.end

		for i := start; i <= end; i++ {
			if hasRepeatedSequence(i) {
				sumOfInvalidIds += i
			}
		}
	}

	return sumOfInvalidIds
}
