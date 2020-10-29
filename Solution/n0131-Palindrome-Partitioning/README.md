# 131. Palindrome Partitioning
## Problem

Given a string *s*, partition *s* such that every substring of the partition is a palindrome.

Return all possible palindrome partitioning of *s*.

**Example:**

```
Input: "aab"
Output:
[
  ["aa","b"],
  ["a","a","b"]
]
```

## Solution

- 见代码
- 时间复杂度：$O(2^n)$，两个字符之间可以选择切开或者不切开，一共有 n - 1 个间隙可以选，故 时间复杂度是 $2^{n - 1}$
- 空间复杂度：$O(n^2)$，递归栈深度为 n，主要空间复杂度来自二维数组