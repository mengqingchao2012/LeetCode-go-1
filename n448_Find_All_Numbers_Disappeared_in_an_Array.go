package main

//解法一：Time：O(n), Space：O(n)
func findDisappearedNumbers(nums []int) []int {
	size := len(nums)
	var res = make([]int, 0, size)
	if size == 0 {
		return nil
	}

	var temp = make([]bool, size)

	for _, v := range nums {
		temp[v-1] = true
	}
	for k, v := range temp {
		if v != true {
			res = append(res, k+1)
		}
	}
	return res
}

//解法二：Time：O(n), Space：O(1)
func findDisappearedNumbers1(nums []int) []int {
	size := len(nums)
	if size == 0 {
		return nil
	}

	var res = make([]int, 0, size)
	for _, value := range nums {
		idx := abs(value) - 1
		if nums[idx] > 0 {
			nums[idx] = -nums[idx]
		}
	}

	for i, v := range nums {
		if v > 0 {
			res = append(res, i+1)
		}
	}
	return res
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
