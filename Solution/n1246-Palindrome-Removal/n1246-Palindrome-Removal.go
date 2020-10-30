package main

import ."LeetCode-go/utils"

func minimumMoves(arr []int) int {
	n := len(arr)
	if n == 0 {
		return 0
	}

	// dp[i][j] 表示移除 [i, j] 这段子数组所需要的最小操作数
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = 1 // 移除单个子数组只需要 1 步
	}

	for j := 0; j < n; j++ {
		for i := j - 1; i >= 0; i-- {
			// 如果 i,j 相邻，则当 arr[i] == arr[j] 时，子数组 s[i,j] 是回文，移除只需要一步，如 11
			// 否则需要两步，如 12
			if i + 1 == j {
				if arr[i] == arr[j] {
					dp[i][j] = 1
				} else {
					dp[i][j] = 2
				}
				continue
			}

			min := n // 记录移除子数组 [i,j] 所需要的最小步数，初始化为 n
			// 如果子数组的两端元素相同，则移除子数组 [i,j] 所需要的步数与移除子数组 [i + 1, j - 1] 是
			// 相同的
			// 场景一：子数组本身就是回文，如 123321, 1234321，这种情况下上述结论显然成立
			// 场景二：子数组中间移除字符后是回文，如 1231, 124521，这种情况下，最终可以递归到非回文的部分，
			// 移除非回文部分所需要的最小步数就等于移除子数组所需要的步数，上述结论也成立
			if arr[i] == arr[j] {
				min = Min(min, dp[i + 1][j - 1])
			}
			// 对于两端元素不相同的情况，则将该部分数组划分为两个部分，[i, k][k + 1, j]，移除这两部分所需要的步数
			// 的和就是移除子数组需要的步数，要想步数最小，就需要枚举分解点 k，找到最小的步数
			for k := i; k < j; k++ {
				min = Min(min, dp[i][k] + dp[k + 1][j])
			}
			// 最后用这个最小值更新 dp[i][j]
			dp[i][j] = min
		}
	}
	return dp[0][n - 1]
}
