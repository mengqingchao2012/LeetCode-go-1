package main

type ArrayReader struct {
	num []int
}

func (a *ArrayReader) get(idx int) int {
	if idx < len(a.num) {
		return a.num[idx]
	}
	return -1
}

func search(reader ArrayReader, target int) int {
	size, start, end:= 2, 0, 1
	for reader.get(end) < target {
		size <<= 1
		start = end + 1
		end = start + size - 1
	}

	low, high := start, end
	for low <= high {
		mid := low + ((high - low) >> 1)
		if reader.get(mid) == target {
			return mid
		} else if reader.get(mid) > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}