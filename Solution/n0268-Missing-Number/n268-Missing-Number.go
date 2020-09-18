package main

// 方法一
func missingNumber(nums []int) int {
	n := len(nums)

	for i := 0; i < n; i++ {
		for nums[i] < n && nums[i] != i {
			v := nums[i]
			nums[i], nums[v] = nums[v], nums[i]
		}
	}

	for i := 0; i < n; i++ {
		if nums[i] != i {
			return i
		}
	}
	// 注意：这里返回的是 n，对应于 nums 为空以及 nums=[9] 这种情况
	// nums 为空返回 0，nums只有一个数（不管是什么）返回 1
	return n
}

// 方法二：位运算
func missingNumber1(nums []int) int {
	res := len(nums)

	// 注意：不能用 res ^= nums[i] ^ i 的方式，go 中会得到不正确的结果
	for i, v := range nums {
		res ^= i ^ v
	}
	return res
}
