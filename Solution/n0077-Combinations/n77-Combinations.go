package main

func combine(n int, k int) [][]int {
	if k > n { return [][]int{} } // 注意边界情况的判断

	result := [][]int{}
	temp := make([]int, 0, k)
	combineRecursive(n, k, 1, temp, &result)
	return result
}

func combineRecursive(n, k, start int, temp []int, result *[][]int) {
	if k == 0 { // 退递归条件：k == 0，说明已经取到了足够的数，将当前的组合结果加入到结果集中
		tmp := make([]int, len(temp))
		copy(tmp, temp)
		*result = append(*result, tmp)
		return
	}

	for i := start; i <= n - k + 1; i++ {
		temp = append(temp, i)
		combineRecursive(n, k - 1, i + 1, temp, result)
		temp = temp[:len(temp) - 1]
	}
}
