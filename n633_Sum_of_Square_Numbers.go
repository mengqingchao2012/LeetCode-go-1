package main

import "math"

//方法一：a^2 + b^2 = c, 则令i=0开始，如果c - i^2 的结果是完全平方数，则满足要求，只需遍历到根号c即可
//Time：O(n^1/2)，Space：O(1)
func judgeSquareSum(c int) bool {
	x := int(math.Sqrt(float64(c)))
	for i := 0; i <= x; i++ {
		if isSquared(c - i * i) {
			return true
		}
	}
	return false
}

func isSquared(i int) bool {
	x := int(math.Sqrt(float64(i)))
	return x * x == i
}

//方法二：双指针
//Time：O(n^1/2)，Space：O(1)
func judgeSquareSum1(c int) bool {
	i, j := 0, int(math.Sqrt(float64(c)))
	for i <= j {
		sum := i * i + j * j
		switch {
		case sum == c:
			return true
		case sum < c:
			i++
		case sum > c:
			j--
		}
	}
	return false
}
