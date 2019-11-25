package main

import (
	"sort"
)

func permuteUnique(nums []int) [][]int {
	if nums == nil {
		return [][]int{}
	}

	var res [][]int

	sort.Ints(nums)
	res = append(res, nums)
	for isFind, num := nextPermutations(nums); isFind == true; {
		res = append(res, num)
		isFind, num = nextPermutations(num)
	}
	return res

}

func nextPermutations(num []int) (bool, []int) {
	n := len(num)
	p := n - 2

	nums := make([]int, n)
	copy(nums, num)

	for p >= 0 && nums[p] >= nums[p+1] {
		p--
	}

	if p >= 0 {
		i := n - 1
		for i > p && nums[i] <= nums[p] {
			i--
		}
		nums[p], nums[i] = nums[i], nums[p]
	}

	j := p + 1
	for i := n - 1; j < i; {
		nums[j], nums[i] = nums[i], nums[j]
		j++
		i--
	}

	return p != -1, nums
}

//func permuteUnique(nums []int) [][]int {
//	if nums == nil {
//		return [][]int{}
//	}
//
//	var res [][]int
//
//	sort.Ints(nums)
//	res = append(res, nums)
//	for isFind, num := nextPermutations(&nums); isFind == true;{
//		res = append(res, *num)
//		isFind, num = nextPermutations(num)
//	}
//	return res
//
//}
//
//func nextPermutations(num *[]int) (bool, *[]int) {
//	n := len(*num)
//	p := n - 2
//
//	nums := make([]int, n)
//	copy(nums, *num)
//
//	for p >= 0 && nums[p] >= nums[p + 1] {
//		p--
//	}
//
//	if p >= 0 {
//		i := n - 1
//		for i > p && nums[i] <= nums[p] {
//			i--
//		}
//		nums[p], nums[i] = nums[i], nums[p]
//	}
//
//	j := p + 1
//	for i := n - 1; j < i; {
//		nums[j], nums[i] = nums[i], nums[j]
//		j++
//		i--
//	}
//
//	return p != -1, &nums
//}

//func main() {
//	var nums = []int{1, 1, 2}
//	res := permuteUnique(nums)
//	fmt.Println(res)
//}
