# 1312. Minimum Insertion Steps to Make a String Palindrome
## Problem

Given a string `s`. In one step you can insert any character at any index of the string.

Return *the minimum number of steps* to make `s` palindrome.

A **Palindrome String** is one that reads the same backward as well as forward.

 

**Example 1:**

```
Input: s = "zzazz"
Output: 0
Explanation: The string "zzazz" is already palindrome we don't need any insertions.
```

**Example 2:**

```
Input: s = "mbadm"
Output: 2
Explanation: String can be "mbdadbm" or "mdbabdm".
```

**Example 3:**

```
Input: s = "leetcode"
Output: 5
Explanation: Inserting 5 characters the string becomes "leetcodocteel".
```

**Example 4:**

```
Input: s = "g"
Output: 0
```

**Example 5:**

```
Input: s = "no"
Output: 1
```

 

**Constraints:**

- `1 <= s.length <= 500`
- All characters of `s` are lower case English letters.

## Solution 

- 思路：
  - 使用第 516 题的方法先找到字符串的最大子序列的长度（注意是最大子序列，不是最大子串），如 `abdbca` 的最大子序列是 `abdba`，这样只剩下一个 `c`，则要使 `abdbca` 变成回文就只需要在恰当的位置加上另一个 `c` 即可
  - 由此可得，需要的操作步骤等于字符串的长度减去最大子序列的长度
  - 扩展：如果题目将插入变成删除，也是同样的思路
  - 时间复杂度：$O(n^2)$，空间复杂度：$O(n^2)$