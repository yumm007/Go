package simplemath

import "math"

func Sqrt(a int) (ret float64) {
	v := math.Sqrt(float64(a))
	return v
}

