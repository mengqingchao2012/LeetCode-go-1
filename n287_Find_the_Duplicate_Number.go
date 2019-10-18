package main

//方法一：二分法，Time：O(nlog(n))，Space：O(1)
func findDuplicate(nums []int) int {
	low, high := 1, len(nums) - 1
	for low < high {
		mid := low + ((high - low) >> 1)
		count := 0
		for _, num := range nums {
			if num <= mid {
				count++
			}
		}

		if count > mid {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return low
}

//方法二：模拟带环链表，Time：O(n)，Space：O(1)
func findDuplicate1(nums []int) int {
	slow := nums[0]
	fast := nums[nums[0]]

	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	ptr := 0
	for slow != ptr {
		slow = nums[slow]
		ptr = nums[ptr]
	}

	return slow
}