# 132. Palindrome Partitioning II
## Problem

Given a string `s`, partition `s` such that every substring of the partition is a palindrome.

Return *the minimum cuts needed* for a palindrome partitioning of `s`.

 

**Example 1:**

```
Input: s = "aab"
Output: 1
Explanation: The palindrome partitioning ["aa","b"] could be produced using 1 cut.
```

**Example 2:**

```
Input: s = "a"
Output: 0
```

**Example 3:**

```
Input: s = "ab"
Output: 1
```

 

**Constraints:**

- `1 <= s.length <= 2000`
- `s` consists of lower-case English letters only.

## Solution

- 见代码，基本思路还是第 5 题的扩展
- 时间复杂度：$O(n^2)$，空间复杂度：$O(n^2)$