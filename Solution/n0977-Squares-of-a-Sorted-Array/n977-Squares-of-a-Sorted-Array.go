package main

func sortedSquares(A []int) []int {
	size := len(A)
	if size == 0 {
		return []int{}
	}

	start, end := 0, size - 1
	res := make([]int, size)
	size -= 1

	for start <= end {
		// 因为数组有序，故只要 A[start] + A[end] >= 0，就说明 A[end] >= A[start]
		if A[start] + A[end] >= 0 {
			res[size] = A[end] * A[end]
			end--
		} else {
			res[size] = A[start] * A[start]
			start++
		}
		size--
	}
	return res
}
