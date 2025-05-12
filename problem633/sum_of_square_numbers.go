package problem633

import "math"

func judgeSquareSum(c int) bool {

	for a := 0; a*a <= c/2; a++ {
		if isSquare(c - a*a) {
			return true
		}
	}

	return false
}

func isSquare(c int) bool {
	sqRoot := math.Sqrt(float64(c))
	return math.Floor(sqRoot) == sqRoot
}
