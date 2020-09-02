# 91. Decode Ways
## Problem

A message containing letters from A-Z is being encoded to numbers using the following mapping:

```'A' -> 1
'A' -> 1
'B' -> 2
...
'Z' -> 26
```

Given a non-empty string containing only digits, determine the total number of ways to decode it.

Example 1:

```
Input: "12"
Output: 2
Explanation: It could be decoded as "AB" (1 2) or "L" (12). ```
```

Example 2:

```
Input: "226"
Output: 3
Explanation: It could be decoded as "BZ" (2 26), "VF" (22 6), or "BBF" (2 2 6).```
```

## Solution

- 动态规划
- 由题意可知，合法的情况有：
  - 对于单个字符组成的数，要求该数在 [1, 9] 之间
  - 对于两个字符组成的数，要求该数在 [10, 26] 之间
- 设 `dp[i]` 表示前 i 个字符能组成的合法结果，则由于最多只能选两个字符组成结果，故：`dp[i] = dp[i - 1] + dp[i - 2]`，其中，对于 `dp[i - 1]` 和 `dp[i - 2]` 要分别判断其合法性，只有合法才能加入到结果集中
- 初始条件为：
  - `dp[0]` 只有一种结果，即空字符串，故 `dp[0] = 1`
  - `dp[1]` 要判断 `s[0]` 是否合法，即 `s[0] == 0` ，不合法，设为 0，否则设为 1