package main

import "container/heap"

// 方法一：快速选择法（线性求第K大）
func partition(nums []int, low, high int) int {
	pivot, l, r := nums[low], low, high // 注意 pivot 的选择是 nums[low]，不能写成 nums[0]
	for l < r {                         // 注意是小于，l == r 时循环退出
		for l < r && nums[r] < pivot {
			r--
		}
		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
		}
		for l < r && nums[l] >= pivot {
			l++
		} // 注意大于等于 pivot 的元素放到左边
		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
		}
	}
	return l
}

func findKthLargest(nums []int, k int) int {
	low, high := 0, len(nums)-1
	for low <= high { // 注意这里是小于等于
		p := partition(nums, low, high)
		if p == k-1 {
			return nums[p]
		} else if p > k-1 {
			high = p - 1
		} else {
			low = p + 1
		}
	}
	return -1
}

// 方法二：最小堆
func findKthLargest1(nums []int, k int) int {
	hp := new(minheap)
	heap.Init(hp)

	for _, v := range nums {
		if hp.Len() < k {
			heap.Push(hp, v)
		} else if v > (*hp)[0] {
			heap.Pop(hp)
			heap.Push(hp, v)
		}
	}
	return heap.Pop(hp).(int)
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
	res := (*m)[n-1]
	*m = (*m)[:n-1]
	return res
}
