package main

//Time：O(mlog(m * n)), Space：O(1)
func findKthNumber(m int, n int, k int) int {
	left, right := 1, m * n + 1

	for left < right {
		mid := left + ((right - left) >> 1)
		if nLessMid(m, n, mid, k) >= k {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func nLessMid(m, n, mid, k int) int {
	count := 0
	for i := 1; i <= minFunc(m, mid); i++ {
		count += minFunc(n, mid / i)
		if count >= k {
			return count
		}
	}
	return count
}

func minFunc(a, b int) int {
	if a < b {
		return a
	}

	return b
}
