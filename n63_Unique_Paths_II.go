package main

//动态规划求解，Time：O(m * n)，m是行数，n是列数，Space：O(n)
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	row := len(obstacleGrid)
	if row == 0 {
		return 0
	}

	col := len(obstacleGrid[0])
	var temp = make([]int, col)
	if obstacleGrid[0][0] == 1 {
		temp[0] = 0
	} else {
		temp[0] = 1
	}

	for i := 0; i < row; i++ {
		if obstacleGrid[i][0] == 1 {
			temp[0] = 0
		} //这里省略了else语句：temp[0] = temp[0]
		for j := 1; j < col; j++ {
			if obstacleGrid[i][j] == 1 {
				temp[j] = 0
			} else {
				//如果位置(i,j)处没有障碍，那么到达位置(i,j)的方法等于从上面下来的方法数加上从左面过来的方法数
				//上面下来的方法数就是temp[j]，左面过来的方法数就是temp[j - 1]
				temp[j] = temp[j] + temp[j - 1]
			}
		}
	}
	return temp[col - 1]
}
