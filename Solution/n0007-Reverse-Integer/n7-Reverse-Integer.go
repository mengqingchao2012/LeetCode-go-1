package main

import "math"

func reverse(x int) int {
	var flag bool
	if x < 0 {
		flag = true
		x = -x
	}
	sum := 0
	for x > 0 {
		sum = sum * 10 + x % 10
		x /= 10
	}
	if flag {
		sum = -sum
	}

	if sum > math.MaxInt32 || sum < math.MinInt32 {
		return 0
	}

	return sum
}
