package main

import (
	"container/heap"
	"sort"
)

// 方法一：
func minMeetingRooms(intervals [][]int) int {
	n := len(intervals)
	if n == 0 {
		return 0
	}
	start, end := make([]int, 0, n), make([]int, 0, n)
	for i := 0; i < n; i++ {
		start = append(start, intervals[i][0])
		end = append(end, intervals[i][1])
	}

	sort.Slice(start, func (i, j int) bool {
		return start[i] < start[j]
	})
	sort.Slice(end, func (i, j int) bool {
		return end[i] < end[j]
	})

	cnt := 0
	e := 0
	for s := 0; s < n; s++ {
		if start[s] < end[e] {
			cnt++
		} else {
			e++
		}
	}
	return cnt
}

// 方法二：最小堆
func minMeetingRooms1(intervals [][]int) int {
	n := len(intervals)
	if n == 0 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	hp := new(minheap)
	heap.Init(hp)
	hp.Push(intervals[0][1])

	for i := 1; i < n; i++ {
		// 如果 intervals[i] 的 start 大于等于堆顶元素，则说明之前有会议在 intervals[i]
		// 开始之前结束，可以复用之前的会议室，弹出已结束会议的时间，加入新的会议开始的时间
		if intervals[i][0] >= (*hp)[0] {
			heap.Pop(hp)
		}
		heap.Push(hp, intervals[i][1])
	}
	return hp.Len()
}

type minheap []int

func (m minheap) Len() int {
	return len(m)
}

func (m minheap) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m minheap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *minheap) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *minheap) Pop() interface{} {
	n := len(*m)
	res := (*m)[n - 1]
	*m = (*m)[:n - 1]
	return res
}
