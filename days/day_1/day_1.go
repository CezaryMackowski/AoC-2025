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

func (d *dial) move(r rotation) int {
	dialSteps := len(d.numbers)
	step := r.steps % dialSteps

	switch r.direction {
	case left:
		// In go -150 mod 100 = -50, so need add dialSteps to change symbol
		d.currentPoint = (d.currentPoint - step + dialSteps) % dialSteps
	case right:
		d.currentPoint = (d.currentPoint + step) % dialSteps
	}

	return d.currentPoint
}

func day1(rotations []rotation) int {
	d := newDial(StartingDialPoint, MaxDialPoints)
	numberOfZeros := 0

	for _, r := range rotations {
		point := d.move(r)
		if point == MinDialPoints {
			numberOfZeros++
		}
	}

	return numberOfZeros
}
