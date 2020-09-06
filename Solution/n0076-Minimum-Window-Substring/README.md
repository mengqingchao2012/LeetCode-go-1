# 76. Minimum Window Substring
## Problem

Given a string S and a string T, find the minimum window in S which will contain all the characters in T in complexity O(n).

**Example:**

```
Input: S = "ADOBECODEBANC", T = "ABC"
Output: "BANC"
```

**Note:**

- If there is no such window in S that covers all characters in T, return the empty string `""`.
- If there is such window, you are guaranteed that there will always be only one unique minimum window in S.

## Solution

- 滑动窗口：
  - 思路类似于 567 题，区别在于这题允许窗口中存在除了 pattern 之外的字符，并且要求这些额外的字符尽量少
  - 窗口收缩的时机：只要匹配了所有的字符，就可以收缩窗口，一直收缩到 matched != len(pattern) 为止
  - 在收缩过程中用当前窗口中元素的个数更新最小字符数，同时更新一个临时的起点变量，表示满足最小字符串结果时，该字符串的起始位置