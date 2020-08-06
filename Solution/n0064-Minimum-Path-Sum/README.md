# 64. Minimum Path Sum
## Problem

Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right which minimizes the sum of all numbers along its path.

Note: You can only move either down or right at any point in time.

Example:

```
 Input:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
Output: 7
Explanation: Because the path 1→3→1→1→1 minimizes the sum.
```

## Solution

- 动态规划
- 解法一：
  - 使用二维数组 `dp[m][n]` 表示状态，遍历从左上角开始，故有：
    - `dp[0][0] = a[0][0]`
    - 第一列：`dp[i][0] = dp[i - 1][0] + a[i][0]`
    - 第一行：`dp[0][j] = dp[0][j - 1] + a[0][j]`
    - 状态转移方程：`dp[i][j] = Min(dp[i][j - 1], dp[i - 1][j]) + a[i][j]`
  - 时间复杂度：$O(m * n)$，空间复杂度：$O(m * n)$
- 解法二：
  - 使用一维数组 `dp[n]` 进行优化：
    - `dp[0] = a[0][0]`
    - 先计算第一行：`dp[i] = dp[0][i - 1] + a[0][i]`
    - 此时，对于坐标为 `(i, j)` 的元素，有：
      - 未更新的 `dp[j]` 表示的是其上面一行同一列的元素，等效于 `dp[i - 1][j]`
      - `dp[j - 1]` 表示的是其左边的元素，等效于 `dp[i][j - 1]`
    - 故：`dp[j] = Min(dp[j], dp[j - 1]) + a[i][j]`
    - 注意，每次遍历完一行，换下一行时，都需要重新初始化：`dp[0] = dp[0] + a[i][0]`，等效于 `dp[i - 1][0] + a[i][0]`
  - 时间复杂度：$O(m * n)$，空间复杂度：$O(n)$