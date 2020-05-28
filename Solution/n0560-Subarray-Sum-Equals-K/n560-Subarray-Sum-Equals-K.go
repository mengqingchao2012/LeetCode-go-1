package main

// 方法一：
func subarraySum(nums []int, k int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	cnt := 0
	for i := 0; i < length; i++ {
		sum := 0
		for j := i; j < length; j++ {
			sum += nums[j]
			if sum == k {
				cnt++
			}
		}
	}
	return cnt
}

// 方法二：
func subarraySum1(nums []int, k int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	preSum := map[int]int{}
	preSum[0] = 1
	sum, cnt := 0, 0
	for i := 0; i < length; i++ {
		sum += nums[i]
		if v, ok := preSum[sum-k]; ok {
			cnt += v
		}
		preSum[sum]++
	}
	return cnt
}
