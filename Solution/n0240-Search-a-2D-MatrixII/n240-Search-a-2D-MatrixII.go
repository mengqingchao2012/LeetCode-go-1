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

	row := 0
	col := n - 1
	for row < m && col >= 0 {
		if target < matrix[row][col] {
			col--
		} else if target > matrix[row][col] {
			row++
		} else {
			return true
		}
	}

	return false
}
