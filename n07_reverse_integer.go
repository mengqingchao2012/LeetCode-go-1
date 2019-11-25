package main

import "math"

func reverse(x int) int {
	sign := 0
	if x < 0 {
		sign = 1
		x = -x
	}

	res := 0
	for x != 0 {
		num := x % 10
		x /= 10
		res = res*10 + num
	}

	if sign == 1 {
		res = -res
	}

	if res > math.MaxInt32 || res < math.MinInt32 {
		return 0
	}

	return res
}
