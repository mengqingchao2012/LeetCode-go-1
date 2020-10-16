package main

func findClosestElements(arr []int, k int, x int) []int {
	low, high := 0, len(arr) - k
	for low < high {
		mid := low + ((high - low) >> 1)
		if x - arr[mid] > arr[mid + k] - x {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return arr[low : low + k]
}
