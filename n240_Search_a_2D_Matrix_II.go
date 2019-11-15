package main

//Time：O(m+n)，Space：O(1)
func searchMatrixII(matrix [][]int, target int) bool {
	n := len(matrix)
	if n == 0 {
		return false
	}

	row := 0
	col := len(matrix[0]) - 1

	for row < n && col >= 0 {
		if matrix[row][col] > target {
			col--
		} else if matrix[row][col] < target {
			row++
		} else {
			return true
		}
	}
	return false
}
