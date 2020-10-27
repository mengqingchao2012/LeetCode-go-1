# 516. Longest Palindromic Subsequence
## Problem

Given a string s, find the longest palindromic subsequence's length in s. You may assume that the maximum length of s is 1000.

**Example 1:**
Input:

```
"bbbab"
```

Output:

```
4
```

One possible longest palindromic subsequence is "bbbb".

 

**Example 2:**
Input:

```
"cbbd"
```

Output:

```
2
```

One possible longest palindromic subsequence is "bb".

 

**Constraints:**

- `1 <= s.length <= 1000`
- `s` consists only of lowercase English letters.

## Solution

- 思路：
  - 使用一个二维数组 `dp[i][j]` 来记录下标为 i，j 之间的子序列的最大回文子序列长度
  - 对于任意 i，j：
    - 如果满足 s[i] == s[j]，则 `dp[i][j]= 2 + dp[i + 1][j - 1]`，即 2 加上下标为 i + 1，j - 1 之间的子序列的最大回文子序列长度
    -  如果 s[i] 和 s[j] 不相等，则 `dp[i][j] = max(dp[i + 1][j], dp[i][j - 1])`
  - 最终返回的是 `dp[0][n - 1]`，即下标为 0，n - 1 之间的子序列的最大回文子序列长度
  - 注意：
    - `dp[i][i] = 1`，因为单个字符的最大回文子序列长度就是 1
    - 要想求出 `dp[i][j]`，则需要先求出 `dp[i + 1][]` 和 `dp[][j - 1]` 的所有值，因此 i 是从大到小遍历，j 是从小到大遍历，同时 j 要大于等于 i，因为下标是从 i 到 j
  - 时间复杂度：$O(n^2)$，空间复杂度：$O(n^2)$