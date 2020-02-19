package main

func searchII(nums []int, target int) bool {
	length := len(nums)
	if length == 0 {
		return false
	}

	low, high := 0, length -1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if nums[mid] == target { //注意要优先判断相等，否则先判断 nums[mid]==nums[low]的话会漏结果，如：在[1]中找1
			return true
		}
		if nums[mid] == nums[low] { //出现特殊情况时，直接将 low++
			low++
			continue
		}
		if nums[mid] >= nums[low] {
			if target < nums[mid] && target >= nums[low] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return false
}