package main

import (
	"math"
)

func splitIntoFibonacci(S string) []int {
	res := []int{}
	dfs(0, len(S), S, &res)
	return res
}

func dfs(position, length int, s string, res *[]int) bool {
	// 遍历完整个数组后，结果集的长度只有大于等于 3 才是合法解
	if position == length {
		return len(*res) >= 3
	}

	// 每个块的开头如果是 0，则该块要么不合法，要么只能是一个单独的 0
	// 由于每个元素 i 都只能是32位整数，故每个块的长度为 1-10，10就是
	// 最大的32位整数的长度
	maxLen := 0
	if s[position] == '0' {
		maxLen = 1
	} else {
		maxLen = 10
	}

	num := 0
	// 对于每一个合法的分割块，其对应元素在字符串 S 中的范围为：[position, Min(position + maxLen, length))
	for i := position; i < Min(position + maxLen, length); i++ {
		num = num * 10 + int(s[i] - '0')
		// 超过 32 位整形范围，剪枝
		if num > math.MaxInt32 {
			break
		}

		// 如果结果集中已经有 2 个块，则第三个块也是确定的，即 F[0] + F[1] = F[2]
		if len(*res) >= 2 {
			last := len(*res) - 1
			sum := (*res)[last] + (*res)[last - 1]
			if sum < num { // 对于 F[2] > F[0] + F[1]，剪枝
				break
			} else if sum > num { // 对于 F[2] < F[0] + F[1]，则还可以继续往后遍历来寻找可能的结果
				continue
			}
		}

		*res = append(*res, num)
		if dfs(i + 1, length, s, res) {
			return true
		}
		// 回溯
		*res = (*res)[:len(*res) - 1]
		*res = (*res)[:len(*res):len(*res)]
	}
	return false
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}