package vecmath

import "math"

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Sign(n int) int {
	if n == 0 {
		return 0
	}
	if n > 0 {
		return 1
	} else {
		return -1
	}
}

func Clamp(value, min, max float64) float64 {
	return math.Max(min, math.Min(max, value))
}

func Interpolate(min, max, t float64) float64 {
	return min + (max-min)*Clamp(t, 0, 1)
}
