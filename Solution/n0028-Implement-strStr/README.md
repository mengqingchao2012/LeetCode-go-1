# 28. Implement strStr()

## Problem

Implement [strStr()](http://www.cplusplus.com/reference/cstring/strstr/).

Return the index of the first occurrence of needle in haystack, or **-1** if needle is not part of haystack.

**Example 1:**

```
Input: haystack = "hello", needle = "ll"
Output: 2
```

**Example 2:**

```
Input: haystack = "aaaaa", needle = "bba"
Output: -1
```

**Clarification:**

What should we return when `needle` is an empty string? This is a great question to ask during an interview.

For the purpose of this problem, we will return 0 when `needle` is an empty string. This is consistent to C's [strstr()](http://www.cplusplus.com/reference/cstring/strstr/) and Java's [indexOf()](https://docs.oracle.com/javase/7/docs/api/java/lang/String.html#indexOf(java.lang.String)).

## Solution

- 双指针分别遍历

  - 优化：外层指针只需遍历两个字符串的长度差，因为如果外层字符串的长度小于内层字符串的长度，肯定不满足条件

- 时间复杂度：$O((n - m + 1) * m)$，n 是外层字符串的长度，m 是内层字符串的长度

- 空间复杂度：$O(1)$

  