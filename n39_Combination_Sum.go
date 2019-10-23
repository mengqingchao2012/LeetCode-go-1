package main

import "sort"

func combinationSum(candidates []int, target int) [][]int {

	sort.Ints(candidates)
	var result [][]int
	var solution []int

	combsum(candidates, solution, target, &result)
	return result
}

func combsum(candidates, solution []int, target int, result *[][]int) {
	if target == 0 {
		*result = append(*result, solution)
		return
	}

	if len(candidates) == 0 || target < candidates[0] {
		return
	}
	//注意这里重新再切切片是为了保证下一次append的时候，solution产生一个新的底层数组，如果不执行这一步的话，
	// 由于切片共享底层数组会导致最终得出错误的答案
	solution = solution[: len(solution) : len(solution)]

	combsum(candidates, append(solution, candidates[0]), target - candidates[0], result)
	combsum(candidates[1:], solution, target, result)
}
