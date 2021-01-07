package main

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	p := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		if p[x] != x {
			p[x] = find(p[x])
		}
		return p[x]
	}

	cnt := n
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if isConnected[i][j] == 1 && p[find(i)] != find(j) {
				p[find(i)] = find(j)
				cnt--
			}
		}
	}
	return cnt
}
