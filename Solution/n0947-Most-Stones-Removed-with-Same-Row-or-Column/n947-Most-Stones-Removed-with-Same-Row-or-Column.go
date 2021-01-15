package main

func removeStones(stones [][]int) int {
	n := len(stones)
	if n == 0 {
		return 0
	}

	parents := make([]int, 2 * 1e4)
	for i := 0; i < len(parents); i++ {
		parents[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		if parents[x] != x {
			parents[x] = find(parents[x])
		}
		return parents[x]
	}

	for _, v := range stones {
		p1, p2 := find(v[0] + 10001), find(v[1])
		if p1 != p2 {
			parents[p1] = p2
		}
	}

	mp := map[int]int{}
	for _, v := range stones {
		mp[find(v[1])]++
	}
	remainedNodeNum := len(mp)
	return n - remainedNodeNum
}
