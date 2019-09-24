package main

import "sort"

//Time：O(n^2)，Space：O(1)
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int

	for i := len(nums) - 1; i >= 2; i-- {
		if nums[i] < 0 { //目标值要求等于0，如果最大数已经小于0了，就不需要再比较了
			break
		}
		j, k, target := 0, i - 1, -nums[i]
		for j < k {
			sum := nums[j] + nums[k]
			switch {
			case sum == target:
				res = append(res, []int{nums[j], nums[k], nums[i]})
				j, k = choose(nums, j, k)
			case sum < target:
				j++
			case sum > target:
				k--
			}
		}
		for i >= 2 && nums[i] == nums[i - 1] { //不要忘记nums[i]也要跳过重复值
			i--
		}
	}
	return res
}

func choose(nums []int, j, k int) (int, int) { //当nums[j]和nums[k]遇到重复值时，跳过重复值
	for j < k {
		switch {
		case nums[j] == nums[j + 1]:
			j++
		case nums[k] == nums[k - 1]:
			k--
		default:
			j++
			k--
			return j, k
		}
	}
	return j, k
}
