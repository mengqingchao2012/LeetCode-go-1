package main

func pivotIndex(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	leftSum := 0
	for i := 0; i < len(nums); i++ {
		if leftSum == sum-nums[i]-leftSum {
			return i
		}
		leftSum += nums[i]
	}
	return -1
}
