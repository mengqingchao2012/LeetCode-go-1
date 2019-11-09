package main

//Time：O(2^n)，Space：O(n)
func subsets(nums []int) [][]int {
	length := len(nums)
	if length == 0 {
		return [][]int{}
	}

	var res [][]int
	var solution []int

	backTracking(nums, solution, 0, &res)
	return res
}

//回溯法，注意区别第39题
func backTracking(nums, solution []int, start int, res *[][]int) {
	*res = append(*res, solution)

	for i := start; i < len(nums); i++ {
		solution = append(solution, nums[i])
		backTracking(nums, solution, i + 1, res)
		solution = solution[: len(solution) - 1]
		solution = solution[: len(solution): len(solution)] //注意此处，不执行这一步分离底层数组的话结果会出错
	}
}

//func main() {
//	nums := []int{1, 2, 3}
//	res := subsets(nums)
//	fmt.Println(res)
//}
