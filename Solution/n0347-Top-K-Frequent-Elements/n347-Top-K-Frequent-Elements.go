package main

import (
	"container/heap"
)

// 方法一：最小堆
func topKFrequent(nums []int, k int) []int {
	mp := map[int]int{}

	for _, v := range nums {
		mp[v]++
	}

	hp := new(minheap)
	heap.Init(hp)

	for key, v := range mp {
		if hp.Len() < k {
			heap.Push(hp, []int{key, v})
		} else if (*hp)[0][1] < v {
			heap.Pop(hp)
			heap.Push(hp, []int{key, v})
		}
	}

	res := []int{}
	for _, v := range *hp {
		res = append(res, v[0])
	}
	return res
}

type minheap [][]int

func (m minheap) Len() int {
	return len(m)
}

func (m minheap) Less(i, j int) bool {
	return m[i][1] < m[j][1]
}

func (m minheap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *minheap) Push(x interface{}) {
	*m = append(*m, x.([]int))
}

func (m *minheap) Pop() interface{} {
	n := len(*m)
	res := (*m)[n - 1]
	*m = (*m)[: n - 1]
	return res
}

// 方法二：快速选择法
func topKFrequent1(nums []int, k int) []int {
	mp := map[int]int{}

	for _, v := range nums {
		mp[v]++
	}

	pairs := make([][2]int, 0, len(mp))
	for key, v := range mp {
		pairs = append(pairs, [...]int{key, v})
	}

	low, high := 0, len(pairs) - 1
	for low <= high {
		p := partitions(pairs, low, high)
		if p == k - 1 {
			res := make([]int, 0, k)
			for _, v := range pairs[: k] { // 注意返回前 k 大的数，是【0, k)
				res = append(res, v[0])
			}
			return res
		} else if p > k - 1 {
			high = p - 1
		} else {
			low = p + 1
		}
	}
	return nil
}

func partitions(pairs [][2]int, low, high int) int {
	pivot, l, r := pairs[low][1], low, high
	for l < r {
		for l < r && pairs[r][1] < pivot { r-- }
		if l < r { pairs[l], pairs[r] = pairs[r], pairs[l] }
		for l < r && pairs[l][1] >= pivot { l++ }
		if l < r { pairs[l], pairs[r] = pairs[r], pairs[l] }
	}
	return l
}

// 方法三：桶排序
func topKFrequent2(nums []int, k int) []int {
	mp := map[int]int{}
	for _, v := range nums {
		mp[v]++
	}

	// 注意，对于大小为 n 的数组，如果所有元素都相同，则桶的下标是 n，故桶的数量
	// 为 n + 1，要加上 0
	buckets := make([][]int, len(nums) + 1)
	for key, v := range mp {
		buckets[v] = append(buckets[v], key)
	}

	res := make([]int, 0, k)
	for i := len(buckets) - 1; i >= 0 && k > 0; i-- {
		bucket := buckets[i]
		if len(bucket) == 0 { // 如果该桶是空，则跳过
			continue
		}
		for j := 0; j < len(bucket) && k > 0; j++ {
			res = append(res, bucket[j])
			k--
		}
	}
	return res
}