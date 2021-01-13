package main

import "sort"

func smallestStringWithSwaps(s string, pairs [][]int) string {
	n := len(s)
	if n < 2 {
		return s
	}

	// 创建并查集
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

	for _, v := range pairs { // 合并连通的集合
		parents[find(v[0])] = find(v[1])
	}

	ids := make([][]int, n) // 用于记录每个连通集合中元素的 id，第一维是集合的根节点的 id
	subStrings := make([][]byte, n) // 用于记录每个连通集合中的元素组成的字符串，第一维是集合的根节点 id

	// 遍历字符串，将字符加入到连通集中
	for i, v := range []byte(s) {
		id := find(i)
		ids[id] = append(ids[id], i)
		subStrings[id] = append(subStrings[id], v)
	}

	res := make([]byte, n)
	for groupID := 0; groupID < n; groupID++ {
		if len(subStrings[groupID]) == 0 {
			continue
		}

		// 将连通集中的字符串进行排序，然后通过查找 ids 数组找到每个元素在字符串中的插入位置，生成结果
		sort.Slice(subStrings[groupID], func(j, k int) bool {
			return subStrings[groupID][j] < subStrings[groupID][k]
		})
		for j := 0; j < len(subStrings[groupID]); j++ {
			res[ids[groupID][j]] = subStrings[groupID][j]
		}
	}
	return string(res)
}