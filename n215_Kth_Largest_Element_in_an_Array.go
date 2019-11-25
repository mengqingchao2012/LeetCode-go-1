package main

import "container/heap"

//方法一：快速选择法：Time(avg)：O(n)，Time(worse)：O(n^2)，Space:O(1)
func findKthLargest(nums []int, k int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		p := partitionFind(&nums, left, right)
		if p == k-1 {
			return nums[p]
		} else if p > k-1 {
			right = p - 1
		} else {
			left = p + 1
		}
	}
	return -1
}

func partitionFind(nums *[]int, left, right int) int {
	pivot := (*nums)[left]
	low, high := left, right
	for low < high {
		for low < high && (*nums)[high] < pivot {
			high--
		}
		if low < high {
			(*nums)[low], (*nums)[high] = (*nums)[high], (*nums)[low]
		}
		for low < high && (*nums)[low] >= pivot {
			low++
		}
		if low < high {
			(*nums)[low], (*nums)[high] = (*nums)[high], (*nums)[low]
		}
	}
	return low
}

//方法二：最小堆，Time：O(nlog(k))，Space：O(k)
func findKthLargest1(nums []int, k int) int {
	temp := new(minheap)
	heap.Init(temp)

	for _, v := range nums {
		if temp.Len() < k {
			heap.Push(temp, v)
		} else if v > (*temp)[0] {
			heap.Pop(temp)
			heap.Push(temp, v)
		}
	}
	return heap.Pop(temp).(int)

}

type minheap []int

func (h minheap) Len() int {
	return len(h)
}

func (h minheap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h minheap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minheap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *minheap) Pop() interface{} {
	n := len(*h)
	res := (*h)[n-1]
	*h = (*h)[:n-1]
	return res
}
