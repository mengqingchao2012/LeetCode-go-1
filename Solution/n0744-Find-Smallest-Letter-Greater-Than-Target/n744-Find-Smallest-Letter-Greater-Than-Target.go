package main

func nextGreatestLetter(letters []byte, target byte) byte {
	n := len(letters)
	if n == 0 {
		return 'a'
	}

	low, high := 0, n - 1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if letters[mid] == target + 1 {
			return letters[mid]
		} else if letters[mid] > target + 1 {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	if low < 0 || low >= n {
		return letters[0]
	} else {
		return letters[low]
	}
}
