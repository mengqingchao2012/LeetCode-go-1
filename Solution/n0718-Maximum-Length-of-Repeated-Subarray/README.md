# 718. Maximum Length of Repeated Subarray

## Problem

Given two integer arrays `A` and `B`, return the maximum length of an subarray that appears in both arrays.

**Example 1:**

```
Input:
A: [1,2,3,2,1]
B: [3,2,1,4,7]
Output: 3
Explanation: 
The repeated subarray with maximum length is [3, 2, 1].
```

 

**Note:**

1. 1 <= len(A), len(B) <= 1000
2. 0 <= A[i], B[i] < 100

## Solution

- 解法一：动态规划

  - 使用 `dp[i][j]` 表示 `A[i:]` 和 `B[j:]` 的最长公共前缀的长度。

    - 如果此时 `A[i] == B[j]`，那么 `dp[i][j]` 的值就等同于 `A[i+1:]` 和 `B[j+1:]` 的最长公共前缀的长度 + 1，即 `dp[i][j] =dp[i + 1][j + 1] + 1 `；
    - 如果 `A[i] != B[j]`，则 `dp[i][j] =  0`

  - 由于 `dp[i][j]` 的状态需要由 `dp[i + 1][j + 1]` 的状态来确定，故选择从后向前遍历

  - 时间复杂度：$O(n * m)$，空间复杂度：$O(n * m)$

    