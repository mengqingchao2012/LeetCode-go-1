package main

func numIslands(grid [][]byte) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])

	visit := make([][]bool, rows * cols)
	for i := 0; i < rows; i++ {
		visit[i] = make([]bool, cols)
	}
	res := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' && visit[i][j] == false {
				dfs(grid, visit, i, j)
				res++
			}
		}
	}
	return res
}

func dfs(grid [][]byte, visit [][]bool, x, y int) {
	rows, cols := len(grid), len(grid[0])
	dx := [4]int{-1, 0, 1, 0}
	dy := [4]int{0, 1, 0, -1}
	visit[x][y] = true

	for i := 0; i < 4; i++ {
		a, b := x + dx[i], y + dy[i]
		if a < 0 || a >= rows || b < 0 || b >= cols || grid[a][b] == '0' || visit[a][b] {
			continue
		}
		dfs(grid, visit, a, b)
	}
}