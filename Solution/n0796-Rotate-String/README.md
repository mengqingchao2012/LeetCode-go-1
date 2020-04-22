# 796. Rotate String

## Problem

We are given two strings, `A` and `B`.

A *shift on `A`* consists of taking string `A` and moving the leftmost character to the rightmost position. For example, if `A = 'abcde'`, then it will be `'bcdea'` after one shift on `A`. Return `True` if and only if `A` can become `B` after some number of shifts on `A`.

```
Example 1:
Input: A = 'abcde', B = 'cdeab'
Output: true

Example 2:
Input: A = 'abcde', B = 'abced'
Output: false
```

**Note:**

- `A` and `B` will have length at most `100`.

## Solution

- 思路：

  - 把两个A拼在一起形成一个新的字符串AA，则如果B成立，B一定是AA的某个子串

- 时间复杂度：$O(n^2)$，空间复杂度：$O(1)$

  