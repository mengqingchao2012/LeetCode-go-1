package main

import ."LeetCode-go/utils"

func minCut(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	isPalindrome := make([][]bool, n)
	for i := 0; i < n; i++ {
		isPalindrome[i] = make([]bool, n)
		isPalindrome[i][i] = true
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if i + 1 == j {
				isPalindrome[i][j] = s[i] == s[j]
			} else {
				isPalindrome[i][j] = s[i] == s[j] && isPalindrome[i + 1][j - 1]
			}
		}
	}

	// 状态数组 cut[i]，i 表示字符的个数，cut[i] 表示包含 1 - i 个字符的子串，
	// 能够按照回文规则划分成的最小块数
	// 如：对于 aab，cut[3] 表示使用了 3 个字符，能够划分的最小块数是 2 [aa b]
	cut := make([]int, n + 1)
	cut[0] = 0 // 一个字符都不取，划分的块数只能是 0
	for i := 1; i <= n; i++ { // i 是终点，范围是 [1, i]
		cut[i] = 1e8 // 对于终点是 i 的这一类情况，先将其最小块数初始化为一个很大的数
		for j := 1; j <= i; j++ { // j 是起点
			if isPalindrome[j - 1][i - 1] { // 注意这里下标要转换到 isPalindrome 中的下标
				// 当前的状态为 [1, j - 1] [j, i]
				// 此时 [j, i] 是回文串，切割后是一个整体，则要想 [1, i] 之间切分的块数最少，
				// 就要求 [1 , j - 1] 之前的切分块数最少，对应的就是 cut[j - 1] 要尽量小，
				// 找到一个合适的 cut[j - 1] 后再加上 [j, i] 这一块，就是 cut[i] 的最小块数
				cut[i] = Min(cut[i], cut[j - 1] + 1)
			}
		}
	}
	// 注意返回值，要求的是切分的操作数，比块数少 1
	return cut[n] - 1
}
