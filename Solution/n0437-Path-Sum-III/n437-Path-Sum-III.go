package main

import . "LeetCode-go/utils"

// 方法一
func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	// 经过根节点的情况 + 不经过根节点的情况
	return traversal(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

func traversal(root *TreeNode, rest int) int {
	if root == nil {
		return 0
	}
	cnt := 0 // 用来存储找到的路径的数量
	if rest == root.Val {
		cnt++
	} // rest 等于当前节点值，说明找到一条路径
	// 注意找到一条路径后不能直接返回，因为后面的路径可能出现和为0的情况，所以还要继续递归
	cnt += traversal(root.Left, rest-root.Val)
	cnt += traversal(root.Right, rest-root.Val)
	return cnt
}

// 方法二：前缀和
func pathSum1(root *TreeNode, sum int) int {
	preSum := map[int]int{}
	preSum[0] = 1 // 注意要把（0,1）这一组值先加入到哈希表中
	return prefixSum(root, sum, 0, preSum)
}

func prefixSum(root *TreeNode, sum int, curSum int, preSum map[int]int) int {
	if root == nil {
		return 0
	}
	curSum += root.Val
	count := 0
	if v, ok := preSum[curSum-sum]; ok { // 注意这里取的是 curSum - sum
		count += v
	}
	preSum[curSum]++
	count += prefixSum(root.Left, sum, curSum, preSum)
	count += prefixSum(root.Right, sum, curSum, preSum)
	preSum[curSum]-- // 退递归时要把加进去的值再减掉
	return count
}
