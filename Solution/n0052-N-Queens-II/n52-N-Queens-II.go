package main

func totalNQueens(n int) int {
	ans := make([]int, n)
	cols := make([]bool, n)
	dg := make([]bool, 2 * n)
	udg := make([]bool, 2 * n)

	var cnt int
	cnt = 0
	dfs(0, n, &cnt, ans, cols, dg, udg)
	return cnt
}

func dfs(row, n int, cnt *int, ans[]int, cols, dg, udg []bool) {
	if row >= n {
		*cnt++
		return
	}

	for c := 0; c < n; c++ {
		if !cols[c] && !dg[c + row] && !udg[c - row + n] {
			ans[c] = row
			cols[c], dg[c + row], udg[c - row + n] = true, true, true
			dfs(row + 1, n, cnt, ans, cols, dg, udg)
			cols[c], dg[c + row], udg[c - row + n] = false, false, false
			ans[c] = 0
		}
	}
}
