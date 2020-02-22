package main

// 方法一：使用哈希表，Time：O(n)，Space：O(n)
func majorityElement(nums []int) int {
	size := len(nums)

	mp := make(map[int]int, size)
	for _, v := range nums {
		mp[v]++
		if mp[v] > size/2 {
			return v
		}
	}
	return 0
}

// 方法二：摩尔投票法，Time：O(n)，Space：O(1)
func majorityElement1(nums []int) int {
	num, count := nums[0], 1

	for i := 1; i < len(nums); i++ {
		if count == 0 {
			num = nums[i]
			count = 1
		} else if nums[i] == num {
			count++
		} else {
			count--
		}
	}
	return num
}
