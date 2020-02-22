package main

func findMin(nums []int) int {
	size := len(nums)
	if size == 0 {
		return -1
	}

	low, high := 0, size-1
	for low < high { //注意退出条件是 low < high
		if nums[low] < nums[high] {
			return nums[low]
		}
		mid := low + ((high - low) >> 1)
		if nums[mid] > nums[high] {
			low = mid + 1
		} else {
			high = mid //注意此时不能将 high 更新为 mid - 1，因为有可能此时最小值就是 mid，要把 mid 也考虑进去
		}
	}
	return nums[low] //注意返回值
}
