package main

import (
	. "LeetCode-go/utils"
)

// 方法一：dp
func lengthOfLIS(nums []int) int {
	length := len(nums)
	if length == 0 { //注意不要漏掉判断数组为空的情况
		return 0
	}

	max := 1
	d := make([]int, length)
	d[0] = 1
	for i := 1; i < length; i++ {
		cur := 0
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				cur = d[j] + 1
			} else {
				cur = 1
			}
			d[i] = Max(cur, d[i])
		}
		max = Max(max, d[i])
	}
	return max
}

// 方法二：dp + 二分
func lengthOfLIS1(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	d := make([]int, length)
	l := 0
	for _, v := range nums {
		i := binarySearchInsertPosition(&d, l, v) // 注意是在结果数组 d 中使用二分法
		d[i] = v
		if i == l {
			l++
		}
	}
	return l
}

func binarySearchInsertPosition(nums *[]int, length int, target int) int {
	low, high := 0, length
	for low <= high {
		mid := low + ((high - low) >> 1)
		if target < (*nums)[mid] {
			high = mid - 1
		} else if target > (*nums)[mid] {
			low = mid + 1
		} else {
			return mid
		}
	}
	return low
}
