package main

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	solution := make([]int, 0, len(candidates))
	res := [][]int{}

	sort.Ints(candidates)
	combination(candidates, 0, target, solution, &res)
	return res
}

func combination(candidates []int, start, target int, solution []int, res *[][]int) {
	if target == 0 {
		*res = append(*res, solution)
		return
	}

	if target < 0 { return }

	for i := start; i < len(candidates); i++ {
		if candidates[i] > target { break }
		// 注意这一步：因为可能存在重复的值导致出现重复结果，要把这些重复的结果去掉
		if i > start && candidates[i] == candidates[i - 1] { continue }
		solution = append(solution, candidates[i])
		combination(candidates, i + 1, target - candidates[i], solution, res)
		solution = solution[: len(solution) - 1]
		solution = solution[: len(solution) : len(solution)]
	}
}
