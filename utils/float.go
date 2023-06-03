package utils

import "math"

const Epsilon float64 = 0.00001

func Equalf(a float64, b float64) bool {
	if math.Abs(a-b) <= Epsilon {
		return true
	}
	return false
}
