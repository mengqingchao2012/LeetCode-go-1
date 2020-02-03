package main

import ."LeetCode-go/utils"

//方法一：Time：O(n)，Space:O(n)
func trap(height []int) int {
	length := len(height)
	if length == 0 {
		return 0
	}

	leftMax, rightMax, water := -1, -1, 0
	temp := make([]int, length)
	for i := 0; i < length; i++ {
		leftMax = Max(leftMax, height[i])
		temp[i] = leftMax
	}
	for i := length - 1; i >= 0; i-- {
		rightMax = Max(rightMax, height[i])
		temp[i] = Min(temp[i], rightMax)
		water += temp[i] - height[i]
	}
	return water
}

//方法二：Time：O(n)，Space:O(1)
func trap1(height []int) int {
	length := len(height)
	if length == 0 {
		return 0
	}

	left, right := 0, length - 1
	leftMax, rightMax := -1, -1
	water := 0
	for left <= right {
		leftMax = Max(leftMax, height[left])
		rightMax = Max(rightMax, height[right])
		if leftMax < rightMax {
			water += leftMax - height[left]
			left++
		} else {
			water += rightMax - height[right]
			right--
		}
	}
	return water
}
