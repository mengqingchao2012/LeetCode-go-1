# 416. Partition Equal Subset Sum
## Problem
Given a **non-empty** array `nums` containing **only positive integers**, find if the array can be partitioned into two subsets such that the sum of elements in both subsets is equal.

 

**Example 1:**

```
Input: nums = [1,5,11,5]
Output: true
Explanation: The array can be partitioned as [1, 5, 5] and [11].
```

**Example 2:**

```
Input: nums = [1,2,3,5]
Output: false
Explanation: The array cannot be partitioned into equal sum subsets.
```

 

**Constraints:**

- `1 <= nums.length <= 200`
- `1 <= nums[i] <= 100`

## Solution

- 0/1 背包：
  - 先对数组求和，如果和是奇数，则直接返回 false，如果和是偶数，将题目变为从数组中任意选数，选出来的数的和能否凑成原数组和的一半（设为 target)，变成了 0/1 背包问题
  - 状态数组：`bool dp[i][j]`：表示前 i 个数的和是否满足小于等于 target，j 的值是枚举从 0 - target
  - 状态转移方程：对于第 i 个元素 nums[i]，有：
    - 如果不选择第 i 个元素，则 `dp[i][j] = dp[i - 1][j]`
    - 如果选择了第 i 个元素且 nums[i] <= j，则 `dp[i][j] = dp[i - 1][j - nums[i]]`
    - 上述两个结果取或，即任一为真即可
  - 时间复杂度：$O(n * target)$，空间复杂度：$O(n * target)$
  - 可以使用一维数组来优化空间复杂度，优化后的空间复杂度是 $O(target)$

