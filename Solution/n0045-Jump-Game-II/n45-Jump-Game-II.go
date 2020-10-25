package main

import ."LeetCode-go/utils"

// 方法一：动态规划，时间复杂度：O(n^2), 空间复杂度：O(n)
func jump(nums []int) int {
	n := len(nums)

	// dp[i] 表示到达第 i 个位置需要的最小步数
	dp := make([]int, n)
	dp[0] = 0 // 对于第一个元素，最小步数为 0
	// 假设所有步数都是最小步数 1，则 n 个元素最多也只需要 n 步，因此将
	// 状态数组中的剩余至赋值为 n + 1，避免使用 math.MaxInt32 导致加法操作溢出
	for i := 1; i < n; i++ {
		dp[i] = n + 1
	}

	for i := 0; i < n; i++ {
		// 对于当前元素 i ，其步数能影响到的下标范围是：[i + 1, i + nums[i]]，
		// 因此只需要修改这个范围呢的状态数组即可，注意加上 j < n 的限制
		for j := i + 1; j <= i + nums[i] && j < n; j++ {
			// 在 [i + 1, i + nums[i]] 范围内，所有的节点都可以通过当前元素 i
			// 的步数 + 1 来到达，因此到达范围内节点的最小步数等于之前其记录的到达
			// 步数与先到达当前元素所需要的最小步数再加 1，两者之间取最小值
			dp[j] = Min(dp[j], dp[i] + 1)
		}
	}
	return dp[n - 1]
}

// 方法二：贪心，时间复杂度：O(n)，空间复杂度：O(n)
func jump1(nums []int) int {
	n := len(nums)

	if n == 0 {
		return -1
	}

	// record 记录到达每个位置需要的最小步数
	record := make([]int, n)
	maxIdx := 0
	for idx, num := range nums {
		if idx > maxIdx { // 说明不可达，返回
			return -1
		}
		if maxIdx >= n - 1 { // 已经找到到达终点的最小步数，返回
			return record[n - 1]
		}
		maxIdx = Max(maxIdx, idx + num) // 更新当前可到达的最远下标
		// 对于 [idx + 1, idx + num] 之间的元素，如果其在 record 中的记录还没有被更新过，
		// 说明之前都未能到达该位置，则到达该位置的最小步数就等于到达 idx 的最小步数 + 1，否则
		// 说明之前已经到达过该位置，现在到达该位置的步数肯定比之前多，不需要再更新
		for i := idx + 1; i <= idx + num && i < n; i++ {
			if record[i] == 0 {
				record[i] = record[idx] + 1
			}
		}
	}
	return -1
}

// 空间优化
func jump2(nums []int) int {
	n := len(nums)

	if n == 0 {
		return -1
	}

	if n == 1 {
		return 0
	}

	maxIdx, curEnd, jumps := 0, 0, 0
	for idx, num := range nums {
		if idx > maxIdx {
			return -1
		}
		if maxIdx >= n - 1 {
			return jumps + 1
		}
		if idx > curEnd {
			jumps++
			curEnd = maxIdx
		}
		maxIdx = Max(maxIdx, idx + num)
	}
	return -1
}
