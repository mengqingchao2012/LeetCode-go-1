package main

import "sort"

// 方法一：
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates) //排序candidates是为了方便之后进行剪枝操作，减少递归的次数

	solution := make([]int, 0, len(candidates)) //用来储存可能满足条件的单个结果集
	result := [][]int{} //用来储存最终返回的结果

	combsum(candidates, target, solution, &result)
	return result
}

func combsum(candidates []int, target int, solution []int, result *[][]int) {
	if target == 0 { // 递归终止条件1：如果target为0，说明找到了一组满足条件的解
		*result = append(*result, solution) //将解加入到结果集中
		return
	}

	// 递归终止条件2：candidates中的所有值都已经取过一遍，等于是完成了所有可能结果的枚举
	// target < candidates[0]：剪枝，因为candidates排过序，如果target小于第一个candidates的值，则说明之后的所有值都不满足要求
	// 注意：要先判断 len(candidates) == 0 的情况，利用短路原则防止 candidates[0] 越界
	if len(candidates) == 0 || target < candidates[0] {
		return
	}

	//注意一定不能漏！！！重新再切切片是为了保证下一次append的时候，solution产生一个新的底层数组，如果不执行这一步的话，
	//由于切片共享底层数组会导致最终得出错误的答案
	solution = solution[:len(solution):len(solution)]

	combsum(candidates, target-candidates[0], append(solution, candidates[0]), result)
	combsum(candidates[1:], target, solution, result)
}

// 方法二：回溯
func combinationSum1(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	solution := make([]int, 0, len(candidates))
	result := [][]int{}

	combin(candidates, target, 0, solution, &result)
	return result
}

func combin(candidates []int, target, start int, solution []int, result *[][]int) {
	if target == 0 {
		*result = append(*result, solution)
		return
	}

	if target < 0 {
		return
	}

	for i := start; i < len(candidates); i++ {
		// 因为已经提前对数组进行了排序，故如果 candidates[i] 已经大于 target，其后面的数
		// 也大于 target，可以提前返回
		if candidates[i] > target { break }
		solution = append(solution, candidates[i])
		// 注意这里 start 还是从 i 开始，因为可以重复取值
		combin(candidates, target - candidates[i], i, solution, result)
		solution = solution[:len(solution) - 1]
		solution = solution[:len(solution):len(solution)]
	}
}
