package main

import (
	"sort"
)

//Time：O(n^3)，Space：O(1)
func fourSum(nums []int, target int) [][]int {
	var res [][]int

	size := len(nums)
	if nums == nil || size < 4 {
		return res
	}
	sort.Ints(nums)

	for i := size - 1; i >= 3; i-- {
		if 4*nums[i] < target {
			break
		}
		for j := i - 1; j >= 2; j-- {
			if 3*nums[j] < target-nums[i] {
				break
			}
			newTarget := target - nums[i] - nums[j]
			left, right := 0, j-1
			for left < right {
				sum := nums[left] + nums[right]
				switch {
				case sum == newTarget:
					res = append(res, []int{nums[left], nums[right], nums[j], nums[i]})
					left, right = check(nums, left, right)
				case sum < newTarget:
					left++
				case sum > newTarget:
					right--
				}
			}
			for j >= 2 && nums[j] == nums[j-1] {
				j--
			}
		}
		for i >= 3 && nums[i] == nums[i-1] {
			i--
		}
	}
	return res
}

func check(nums []int, left, right int) (int, int) {
	for left < right {
		switch {
		case nums[left] == nums[left+1]:
			left++
		case nums[right] == nums[right-1]:
			right--
		default:
			left++
			right--
			return left, right
		}
	}
	return left, right
}

//func main() {
//	res := fourSum([]int{0, 0, 0, 0}, 0)
//	fmt.Println(res)
//}
