package main

func makeConnected(n int, connections [][]int) int {
	size := len(connections)
	if size < n - 1 {
		return -1
	}

	parents := make([]int, n)
	for i := 0; i < n; i++ {
		parents[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		if parents[x] != x {
			parents[x] = find(parents[x])
		}
		return parents[x]
	}

	cnt := n // 用于记录独立的连通节点的个数
	redundancy := 0 // 用于记录冗余网线的数目
	for _, v := range connections {
		p1, p2 := find(v[0]), find(v[1])
		if p1 != p2 {
			parents[p1] = p2
			cnt-- // 每新加入一个节点，独立的连通节点数就减一
		} else {
			redundancy++
		}
	}

	if cnt - 1 > redundancy {
		return -1
	} else {
		return cnt - 1
	}
}