package main

func checkPossibility(nums []int) bool {
	length := len(nums)
	if length == 0 {
		return false
	}
	if length == 1 {
		return true
	}
	modified := nums[0] > nums[1]
	for i := 1; i < len(nums) - 1; i++ {
		if nums[i] > nums[i + 1] {
			if modified {
				return false
			}
			if nums[i + 1] >= nums[i - 1] {
				nums[i] = nums[i - 1]
			} else {
				nums[i + 1] = nums[i]
			}
			modified = true
		}
	}
	return true
}
