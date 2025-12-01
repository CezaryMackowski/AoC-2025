package day_1

const (
	MinDialPoints     = 0
	StartingDialPoint = 50
	MaxDialPoints     = 99
)

type direction string

const (
	left  direction = "L"
	right direction = "R"
)

type rotation struct {
	direction direction
	steps     int
}

type dial struct {
	numbers       []int
	startingPoint int
	currentPoint  int
}

func newDial(start int, max int) *dial {
	var a []int

	for i := MinDialPoints; i <= max; i++ {
		a = append(a, i)
	}

	return &dial{
		numbers:       a,
		startingPoint: start,
		currentPoint:  start,
	}
}

func (d *dial) stepsUntilZero(r rotation) int {
	dialSize := len(d.numbers)

	switch r.direction {
	case left:
		if d.currentPoint == MinDialPoints {
			return dialSize
		}
		return d.currentPoint
	case right:
		distance := (dialSize - d.currentPoint) % dialSize
		if distance == 0 {
			return dialSize
		}
		return distance
	}
	return dialSize
}

func (d *dial) move(r rotation) (int, int) {
	dialSize := len(d.numbers)
	totalSteps := r.steps

	stepsUntilZero := d.stepsUntilZero(r)

	zeroHits := 0
	if totalSteps >= stepsUntilZero {
		zeroHits = 1 + (totalSteps-stepsUntilZero)/dialSize
	}

	switch r.direction {
	case left:
		moveSteps := totalSteps % dialSize
		d.currentPoint = (d.currentPoint - moveSteps + dialSize) % dialSize
	case right:
		d.currentPoint = (d.currentPoint + totalSteps) % dialSize
	}

	if d.currentPoint == MinDialPoints && zeroHits > 0 {
		zeroHits-- // final click is handled by day1Extra
	}

	return d.currentPoint, zeroHits
}

func day1(rotations []rotation) int {
	d := newDial(StartingDialPoint, MaxDialPoints)
	numberOfZeros := 0

	for _, r := range rotations {
		point, _ := d.move(r)
		if point == MinDialPoints {
			numberOfZeros++
		}
	}

	return numberOfZeros
}

func day1Extra(rotations []rotation) int {
	d := newDial(StartingDialPoint, MaxDialPoints)
	numberOfZeros := 0

	for _, r := range rotations {
		point, wraps := d.move(r)
		if point == MinDialPoints {
			numberOfZeros++
		}
		numberOfZeros += wraps
	}

	return numberOfZeros
}
