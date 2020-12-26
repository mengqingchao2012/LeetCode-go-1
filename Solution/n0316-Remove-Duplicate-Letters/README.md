# 316. Remove Duplicate Letters
## Problem

Given a string `s`, remove duplicate letters so that every letter appears once and only once. You must make sure your result is **the smallest in lexicographical order** among all possible results.

**Note:** This question is the same as 1081: https://leetcode.com/problems/smallest-subsequence-of-distinct-characters/

 

**Example 1:**

```
Input: s = "bcabc"
Output: "abc"
```

**Example 2:**

```
Input: s = "cbacdcbc"
Output: "acdb"
```

 

**Constraints:**

- `1 <= s.length <= 104`
- `s` consists of lowercase English letters.

## Solution

- 贪心
- 对于字符串中的任意下标对应的字符，可以选择让其出现在最终结果中，或者不选
- 要使得最终字符串的字典和最小，则对于当前下标 i 对应的字符 x，将 x 和最终结果集里的最后一个元素 y 进行比较：
  - 如果 y 大于 x ，且字符 y 在下标 i 之后又重新出现，则可以将 y 删掉，加入 x
- 时间复杂度：$O(n)$，空间复杂度：$O(n)$