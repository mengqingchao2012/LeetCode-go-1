package main

func searchRange(nums []int, target int) []int {
	n := len(nums)
	if n == 0 {
		return []int{-1, -1}
	}

	end := binarySearchEndIndex(nums, target) // 找到 target 最后出现的位置
	if nums[end] != target { // 如果 target 不存在，则直接返回
		return []int{-1, -1}
	}

	start := binarySearchEndIndex(nums, target - 1) // 通过搜索 target - 1 来间接找 target 的起始位置
	if start >= 0 && start <= end && end < n {
		// 注意，这里确定 start 不能直接使用 start + 1，而要加上判断条件的原因是：
		// 如果 target - 1 比 nums[0] 要小，则返回的 start 为 0，此时如果我们的 target == nums[0]，
		// 则 start++ 就得到了错误的结果，因此要加上判定条件。
		// 同时注意不能使用 start != 0 来作为判定，比如在 [1, 3, 3, 4] 中找 3，如果使用 start != 0 作为
		// 判定，返回的就是 [0, 2]，是个错误的结果
		if nums[start] != target {
			start++
		}
		return []int{start, end}
	}

	// 同理，如果选择先找出 target 出现的第一个位置，在使用 end = binarySearchStartIndex(nums, target + 1)
	// 确定 target 的最后出现位置时，要注意取到右边界的情况，故条件应该为：if nums[end] != target { end-- }

	return []int{-1, -1}
}

// 该方法返回数组中元素出现的最后位置
// 注意：如果查找某个不存在的数，返回的是比该数小的所有数中最大的那个
//  nums[mid] <= target
// ---------------------|
// ---------------------|----------------
// a                  b`b                c
// [a, b] 之间的所有数都满足 <= target, (b, c] 之间的所有数都大于 target
// 故 b 点是 target 可能出现的最大下标
// 如果 target 不存在，则返回的是最接近 target 的 b`
func binarySearchEndIndex(nums []int, target int) int {
	low, high := 0, len(nums) - 1
	for low < high {
		mid := (low + high + 1) >> 1
		if nums[mid] <= target {
			low = mid
		} else {
			high = mid - 1
		}
	}
	return low
}

// 该方法返回数组中元素出现的第一个位置
// 注意：如果查找某个不存在的数，返回的是比该数大的所有数中最小的那个
//                 nums[mid] >= target
//              |-----------------------
// -------------|-----------------------
// a            b b'                    c
// [b, c] 之间的所有数都满足 >= target, [a, b) 之间的所有数都小于 target
// 故 b 点是 target 可能出现的最小下标
// 如果 target 不存在，则返回的是最接近 target 的 b`
func binarySearchStartIndex(nums []int, target int) int {
	low, high := 0, len(nums) - 1
	for low < high {
		mid := (low + high) >> 1
		if nums[mid] >= target {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}
