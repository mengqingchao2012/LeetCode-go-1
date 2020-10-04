package main

func subsets(nums []int) [][]int {
	n := len(nums)
	if n == 0 {
		return [][]int{}
	}

	cur := make([]int, 0, n)
	// 共有 2^n - 1 个结果，故这里直接通过 1 << n 来初始化内存
	res := make([][]int, 0, 1 << n)
	// 注意 res 要用指针
	backTrace(nums, 0, cur, &res)
	return res
}

func backTrace(nums []int, start int, cur []int, res *[][]int) {
	*res = append(*res, cur)
	for i := start; i < len(nums); i++ {
		cur = append(cur, nums[i])
		backTrace(nums, i + 1, cur, res)
		// 回溯，注意要通过 cur[:len(cur):len(cur)] 来分离 cur 的底层数组
		cur = cur[:len(cur) - 1]
		cur = cur[:len(cur):len(cur)]
	}
}
