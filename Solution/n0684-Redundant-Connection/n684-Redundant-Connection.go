package main

func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	if n == 0 {
		return []int{}
	}

	parents := make([]int, n + 1)
	for i := 1; i <= n; i++ {
		parents[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		if parents[x] != x {
			parents[x] = find(parents[x])
		}
		return parents[x]
	}

	for _, v := range edges {
		if find(v[0]) == find(v[1]) {
			return v
		} else {
			parents[find(v[0])] = find(v[1])
		}
	}
	return []int{}
}