package main

import "container/heap"

// 方法一：快速选择法（线性求第K大）
func findKthLargest(nums []int, k int) int {
	n := len(nums)
	res := quickSelect(nums, 0, n - 1, k)
	return res
}

func quickSelect(nums []int, left, right, k int) int {
	if left == right {
		return nums[left]
	}

	// 注意：因为是要求第 k 大的数，所以选择将数组按照降序排序，这样第 k 大的数会落在最终结果的前半部分
	pivot, i, j := nums[left + ((right - left) >> 1)], left, right
	for {
		for nums[i] > pivot { i++ }
		for nums[j] < pivot { j-- }
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
		i++; j--
	}

	// 分界点将数组分成两部分，前半部分的所有数都 >= pivot，后半部分的所有数都 <= pivot
	sl := j - left + 1 // 计算出前半部分共有多少个数
	if k <= sl { // 如果前半部分数的个数小于等于 k，说明第 k 大的数肯定在前半部分，则递归的到前半部分进行查找
		return quickSelect(nums, left, j, k)
	} else { // 否则到后半部分进行查找，要注意此时由于查找数的区间变了，最终的 k 要减去前半部分的数的个数
		return quickSelect(nums, j + 1, right, k - sl)
	}
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
