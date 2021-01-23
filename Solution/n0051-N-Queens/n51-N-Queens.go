package main

import "strings"

func solveNQueens(n int) [][]string {
	cols := make([]bool, n) // 记录使用过哪些列
	dg := make([]bool, 2 * n) // 记录使用过的正向对角线（斜率为1），n * n 的矩阵的对角线数为 2n - 1
	udg := make([]bool, 2 * n) // 记录使用过的反向对角线（斜率为 -1）

	tmp := strings.Repeat(".", n)
	ans := make([][]byte, n)
	for i := 0; i < n; i++ {
		ans[i] = []byte(tmp)
	}
	res := [][]string{}
	dfs(0, n, &res, ans, cols, dg, udg)
	return res
}

func dfs(row int, n int, res *[][]string, ans [][]byte, cols, dg, udg []bool) {
	if row >= n { // row >= n，说明已经得到了一组可行的解
		tmp := make([]string, n)
		for i := 0; i < n; i++ {
			tmp[i] = string(ans[i])
		}
		*res = append(*res, tmp)
		return
	}

	for c := 0; c < n; c++ {
		// 斜率的计算公式为：y = x + d 或 y = -x + d，推出 b = y - x 或 b = y + x，故可以使用 b 来
		// 表示每条对角线，需要注意 y - x 可能为负，所以加上一个偏移量 n 来防止越界
		if !cols[c] && !dg[c + row] && !udg[c - row + n] {
			ans[row][c] = 'Q'
			cols[c], dg[c + row], udg[c - row + n] = true, true, true
			// 回溯
			dfs(row + 1, n, res, ans, cols, dg, udg)
			ans[row][c] = '.'
			cols[c], dg[c + row], udg[c - row + n] = false, false, false
		}
	}
}
