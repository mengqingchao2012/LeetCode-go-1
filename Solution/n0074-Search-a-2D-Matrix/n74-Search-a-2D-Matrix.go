package main

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}

	n := len(matrix[0])
	if n == 0 {
		return false
	}

	low, high := 0, m*n-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		row := mid / n
		col := mid % n
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}
