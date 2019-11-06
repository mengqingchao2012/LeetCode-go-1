package main

//Time：O(n^2), Space: O(1)
func rotate1(matrix [][]int)  {
	n := len(matrix)
	if n == 0 || len(matrix[0]) == 0 {
		return
	}

	//第一步将沿主对角线对称的两个数相互换位
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ { //注意 j 的取值
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	//第二步沿垂直中心线将左右两边对称的两个数相互换位
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ { //注意 j 的取值
			matrix[i][j], matrix[i][n - 1 - j] = matrix[i][n - 1 - j], matrix[i][j]
		}
	}

}
