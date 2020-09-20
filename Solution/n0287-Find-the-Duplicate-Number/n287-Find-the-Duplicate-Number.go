package main

// 二分法
func findDuplicate(nums []int) int {
	low, high := 1, len(nums)-1 // low 和 high 代表的不是下标，而是区间[1,n] 的首尾值
	for low < high {
		mid := low + ((high - low) >> 1)
		count := 0
		for _, v := range nums {
			if v <= mid {
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

// 模拟带环链表
func findDuplicate1(nums []int) int {
	slow, fast := nums[0], nums[nums[0]] //注意初始化时slow和fast不能相同
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	p := 0
	for p != slow {
		p = nums[p]
		slow = nums[slow]
	}
	return slow
}

// Cyclic Sort
func findDuplicate2(nums []int) int {
	n := len(nums)

	for i := 0; i < n; i++ {
		for nums[i] - 1 != i {
			v := nums[i] - 1
			// 当交换到重复数时，会发生死循环，这时候应该直接返回
			if nums[i] == nums[v] {
				return nums[i]
			}
			nums[i], nums[v] = nums[v], nums[i]
		}
	}
	return -1
}
