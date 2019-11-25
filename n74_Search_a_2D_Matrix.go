package main

//Time：O(m * n)，m和n分别代表行数和列数，Space:O(1)
func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}

	n := len(matrix[0])

	left, right := 0, m*n-1
	for left <= right {
		mid := left + ((right - left) >> 1)
		row := mid / n
		col := mid % n
		if target < matrix[row][col] {
			right = mid - 1
		} else if target > matrix[row][col] {
			left = mid + 1
		} else {
			return true
		}
	}

	return false
}
