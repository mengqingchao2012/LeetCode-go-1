package main

import ."LeetCode-go/utils"

// 时间复杂度：O(n^3 * K），空间复杂度：O(n^3)
func mergeStones(stones []int, K int) int {
	n := len(stones)
	if (n - 1) % (K - 1) != 0 {
		return -1
	}
	prefixSum := make([]int, n + 1)
	for i := 0; i < n; i++ {
		prefixSum[i + 1] = prefixSum[i] + stones[i]
	}

	// dp[i][j][k] 表示把子数组 [i,j] 合并成 k 堆的最小开销
	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, K + 1)
			for k := 0; k < K + 1; k++ {
				dp[i][j][k] = 1e8 // 初始化为一个极大值
			}
		}
		dp[i][i][1] = 0 // 将单个元素合成 1 堆的开销是 0
	}

	for l := 2; l <= n; l++ { // 待合并的数组区间长度，最小是 2
		for i := 0; i <= n - l; i++ { // 区间的起点
			j := i + l - 1 // 区间的终点
			for k := 2; k <= K; k++ { // 合成的堆数，枚举从 2 堆到 K 堆
				for p := i; p < j; p = p + K - 1 { // 将区间分割成 1 堆 和 k - 1 堆，p 是分割点
					// 注意，不是任意 p 都能满足条件，因为枚举分割点最终要求左边能合并成 1 堆，这个时候
					// 要满足 (p - i) % (K - 1) == 0，因此分割点每次更新的步长可以直接设置为 K - 1，
					// 免去了很多不必要的计算

					// 将 [i,p] 合成 1 堆，[p + 1, j] 合成 k - 1 堆，这样整个区间就有 k 堆
					dp[i][j][k] = Min(dp[i][j][k], dp[i][p][1] + dp[p + 1][j][k - 1])
				}
			}
			// 到了这里，已经得到了将 [i,j] 合并成 K 堆的最小开销，还差最后一步，将这 K 堆合并成 1 堆，
			// 实际上就等于将区间合并成 K 堆的最小开销加上区间内所有元素的和
			dp[i][j][1] = dp[i][j][K] + prefixSum[j + 1] - prefixSum[i]
		}
	}
	return dp[0][n - 1][1]
}

// 方法一实际上可以继续改进，去掉第三维，因为 K 是固定的，枚举分割点时，左边合并为 1 堆，
// 右边则尽量的去合并区间，合并完成后如果右边恰好是 K - 1 堆，则整个区间可以合并成 1 堆，
// 此时需满足 (j - i) % (K - 1) == 0，即能分割成的堆数是确定的，不是任选都能满足条件
// 方法二：时间复杂度：O(n^3/K), 空间复杂度：O(n^2)
func mergeStones1(stones []int, K int) int {
	n := len(stones)
	if (n - 1) % (K - 1) != 0 {
		return -1
	}
	prefixSum := make([]int, n + 1)
	for i := 0; i < n; i++ {
		prefixSum[i + 1] = prefixSum[i] + stones[i]
	}

	// dp[i][j] 表示尽力去合并 [i, j] 的最小开销
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = 0
	}

	for l := K; l <= n; l++ {
		for i := 0; i + l <= n; i++ {
			j := i + l - 1
			dp[i][j] = 1e8
			for p := i; p < j; p = p + K - 1 {
				dp[i][j] = Min(dp[i][j], dp[i][p] + dp[p + 1][j])
			}
			// [i, j] 合并完成后，如果整个区间长度满足 (j - i) % (K - 1) == 0，则将其合并为 1 堆
			if (l - 1) % (K - 1) == 0 {
				dp[i][j] += prefixSum[j + 1] - prefixSum[i]
			}
		}
	}
	return dp[0][n - 1]
}