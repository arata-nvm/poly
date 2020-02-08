package poly

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
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
