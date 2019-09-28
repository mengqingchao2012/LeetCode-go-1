package main

func hammingDistance(x int, y int) int {
	return countOne(x ^ y)
}

func countOne(n int) int {
	count := 0
	for n != 0 {
		count++
		n &= n - 1
	}
	return count
}
