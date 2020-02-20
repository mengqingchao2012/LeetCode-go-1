package main

func twoSum(numbers []int, target int) []int {
	low, high := 0, len(numbers) - 1
	for low <= high {
		sum := numbers[low] + numbers[high]
		if sum > target {
			high--
		} else if sum < target {
			low++
		} else {
			return []int{low + 1, high + 1}
		}
	}
	return []int{-1, -1}
}
