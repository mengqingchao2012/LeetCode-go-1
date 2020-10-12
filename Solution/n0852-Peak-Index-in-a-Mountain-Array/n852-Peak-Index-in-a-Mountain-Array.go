package main

func peakIndexInMountainArray(arr []int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}

	low, high := 0, n - 1
	for low < high {
		mid := low + ((high - low) >> 1)
		if arr[mid] > arr[mid + 1] {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}