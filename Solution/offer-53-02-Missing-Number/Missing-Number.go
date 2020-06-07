package main

func missingNumber(nums []int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}

	low, high := 0, n - 1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if nums[mid] != mid { // 如果 nums[mid] != mid，说明不一致出现在 mid 之前
			if mid == 0 || nums[mid - 1] == mid - 1 { // 注意要先判断 mid == 0 的情况，否则 mid - 1 可能越界，如 [1] 应该返回 0
				return mid
			}
			high = mid - 1
		} else { // 否则说明不一致出现在 mid 之后
			low = mid + 1
		}
	}

	// 注意这个判断不能漏，如 [0] 应该返回 1， [0, 1] 应该返回 2
	if low == n {
		return n
	}

	return -1
}
