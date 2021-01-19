package main

import (
	. "LeetCode-go/utils"
	"math"
	"sort"
)

const Inf = math.MaxInt32

// Prim 算法，时间复杂度：O(n^2)，空间复杂度：O(n^2）
func prim(matrix [][]int, dist []int, set []bool) int {
	res, n := 0, len(set)
	for i := 0; i < n; i++ {
		t := -1
		for j := 0; j < n; j++ {
			if !set[j] && (t == -1 || dist[t] > dist[j]) {
				t = j
			}
		}

		set[t] = true

		for k := 0; k < n; k++ {
			if !set[k] {
				dist[k] = Min(dist[k], matrix[t][k])
			}
		}

		if i != 0 { res += dist[t] }
	}
	return res
}

func minCostConnectPoints(points [][]int) int {
	n := len(points)
	if n == 0 {
		return 0
	}

	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j] = Abs(points[i][0] - points[j][0]) + Abs(points[i][1] - points[j][1])
			matrix[j][i] = matrix[i][j]
		}
	}

	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = Inf
	}

	set := make([]bool, n)

	return prim(matrix, dist, set)
}

// Kruskal 算法，时间复杂度：O(nlogn)，空间复杂度：O(n^2）
type Edge struct {
	p1 int
	p2 int
	w int
}

func minCostConnectPoints1(points [][]int) int {
	n := len(points)
	if n == 0 {
		return 0
	}

	edges := []Edge{}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			edge := Edge {p1: i, p2: j, w: Abs(points[i][0] - points[j][0]) + Abs(points[i][1] - points[j][1])}
			edges = append(edges, edge)
		}
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].w < edges[j].w
	})

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

	res := 0
	for i := 0; i < len(edges); i++ {
		a, b, w := edges[i].p1, edges[i].p2, edges[i].w
		f1, f2 := find(a), find(b)
		if f1 != f2 {
			parents[f1] = f2
			res += w
		}
	}

	return res
}