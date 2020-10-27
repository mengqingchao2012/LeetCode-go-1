# 5. Longest Palindromic Substring
## Problem

Given a string `s`, return *the longest palindromic substring* in `s`.

 

**Example 1:**

```
Input: s = "babad"
Output: "bab"
Note: "aba" is also a valid answer.
```

**Example 2:**

```
Input: s = "cbbd"
Output: "bb"
```

**Example 3:**

```
Input: s = "a"
Output: "a"
```

**Example 4:**

```
Input: s = "ac"
Output: "a"
```

 

**Constraints:**

- `1 <= s.length <= 1000`
- `s` consist of only digits and English letters (lower-case and/or upper-case)

## Solution

- 解法一：动态规划
  - 状态数组：`bool dp[i][j]` 表示 [i, j] 之间的子串是否是回文
  - 判断是否是回文子串：
    - 如果 i == j，则单个字符必然是回文子串
    - 如果 i + 1 == j，即 i 和 j 相邻，则当且仅当 s[i] == s[j] 时，是回文子串
    - 对于其他情况，则需要同时满足 s[i] == s[j] 且 [i + 1, j - 1] 之间的子串是回文子串
  - 由于当前 i，j 的状态与 i + 1 和 j - 1 相关，故 i 从大到小遍历，j 从小到大遍历
  - 遍历的同时记录 maxLen 和 start，当 j - i + 1 > maxLen，即回文子串的长度大于 maxLen 时，更新
  - 时间复杂度：$O(n^2)$，空间复杂度：$O(n^2)$
- 解法二：中心扩展
  - 遍历字符串，以每个字符为中心，向左右两边扩展求出其最长子串
  - 注意求最长子串时有两种情况，以当前字符为中心，子串的长度为奇数或者子串的长度为偶数，以该字符为中心的最长子串的长度是两种中较大的那一个
  - 时间复杂度：$O(n^2)$，空间复杂度：$O(1)$