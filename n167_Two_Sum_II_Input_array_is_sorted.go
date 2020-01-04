package main

//使用双指针
func twoSumInSortedArray(numbers []int, target int) []int {
	length := len(numbers)
	if length == 0 {
		return []int{-1, -1}
	}

	left, right := 0, length - 1
	for left < right {
		if numbers[left] + numbers[right] > target {
			right--
		} else if numbers[left] + numbers[right] < target {
			left++
		} else {
			return []int{left + 1, right + 1}
		}
	}
	return []int{-1, -1}
}
