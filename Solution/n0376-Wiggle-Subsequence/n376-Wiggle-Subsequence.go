package main

func wiggleMaxLength(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}

	prediff := nums[1] - nums[0]
	res := 0
	if prediff != 0 {
		res = 2
	} else {
		res = 1
	}

	for i := 2; i < n; i++ {
		diff := nums[i] - nums[i - 1]
		if diff < 0 && prediff >= 0 {
			res++
			prediff = diff
		} else if diff > 0 && prediff <= 0 {
			res++
			prediff = diff
		}
	}
	return res
}