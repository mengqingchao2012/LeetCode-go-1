package main

import (
	. "LeetCode-go/utils"
	"sort"
)

// 方法一：二分 + bfs
func minimumEffortPath(heights [][]int) int {
	left, right := 0, int(1e6)
	for left < right {
		mid := (left + right) >> 1
		if check(heights, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func check(h [][]int, target int) bool {
	rows, cols := len(h), len(h[0])
	direction := [][2]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	visit := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visit[i] = make([]bool, cols)
	}
	queue := [][2]int{}

	queue = append(queue, [2]int{0, 0})
	visit[0][0] = true
	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		x, y := cur[0], cur[1]

		for _, v := range direction {
			dx, dy := v[0], v[1]
			a, b := x + dx, y + dy
			if a < 0 || a >= rows || b < 0 || b >= cols || visit[a][b] || Abs(h[a][b] - h[x][y]) > target {
				continue
			}

			queue = append(queue, [2]int{a, b})
			visit[a][b] = true
		}
	}
	return visit[rows - 1][cols - 1]
}

// 贪心 + 并查集
func minimumEffortPath1(heights [][]int) int {
	rows, cols := len(heights), len(heights[0])
	edges := make([][3]int, 0, rows * cols)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if i != rows - 1 {
				edges = append(edges, [3]int{Abs(heights[i][j] - heights[i + 1][j]), i * cols + j, (i + 1) * cols + j})
			}
			if j != cols - 1 {
				edges = append(edges, [3]int{Abs(heights[i][j] - heights[i][j + 1]), i * cols + j, i * cols + (j + 1)})
			}
		}
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][0] < edges[j][0]
	})

	parents := make([]int, rows * cols)
	for i := 0; i < rows * cols; i++ {
		parents[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		if parents[x] != x {
			parents[x] = find(parents[x])
		}
		return parents[x]
	}

	for _, edge := range edges[:] {
		p1, p2 := find(edge[1]), find(edge[2])
		if p1 != p2 {
			parents[p1] = p2
		}
		if find(0) == find((rows - 1) * cols + (cols - 1)) {
			return edge[0]
		}
	}
	return 0
}
