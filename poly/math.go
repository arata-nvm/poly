package poly

import "math"

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func sign(n int) int {
	if n == 0 {
		return 0
	}
	if n > 0 {
		return 1
	} else {
		return -1
	}
}

func clamp(value, min, max float64) float64 {
	return math.Max(min, math.Min(max, value))
}

func interpolate(min, max, t float64) float64 {
	return min + (max-min)*clamp(t, 0, 1)
}
