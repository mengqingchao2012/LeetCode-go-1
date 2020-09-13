package main

import ."LeetCode-go/utils"

func insert(intervals [][]int, newInterval []int) [][]int {
	n := len(intervals)
	res := [][]int{}
	if n == 0 {
		return append([][]int{}, newInterval)
	}

	// 二分法分别找出 newInterval 中 start 和 end 在 intervals 中的插入位置
	start := binarySearch(intervals, newInterval[0])
	end := binarySearch(intervals, newInterval[1])

	// 对于 newInterval[0] 之前的元素，可以直接加入结果集
	res = append(res, intervals[:start]...)

	// 待合并的元素实际上就是 interval[start : end + 1]
	tmp := make([][]int, 0, end - start + 1)
	tmp = append(tmp, newInterval)

	// 注意处理 end + 1 超出数组范围的情况
	size := Min(end + 1, n)
	tmp = append(tmp, intervals[start:size]...)
	
	for _, v := range tmp {
		n := len(res)
		if n == 0 || res[n-1][1] < v[0] {
			res = append(res, v)
		} else {
			res[n-1][1] = Max(res[n-1][1], v[1])
		}
	}
	// 把剩余的元素放入结果集中
	res = append(res, intervals[size:]...)
	return res
}

func binarySearch(intervals [][]int, target int) int {
	low, high := 0, len(intervals) - 1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if intervals[mid][0] == target {
			return mid
		} else if intervals[mid][0] > target {
			high--
		} else {
			low++
		}
	}
	return low
}
