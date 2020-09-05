# 424. Longest Repeating Character Replacement
## Problem

Given a string `s` that consists of only uppercase English letters, you can perform at most `k` operations on that string.

In one operation, you can choose **any** character of the string and change it to any other uppercase English character.

Find the length of the longest sub-string containing all repeating letters you can get after performing the above operations.

**Note:**
Both the string's length and *k* will not exceed 104.

**Example 1:**

```
Input:
s = "ABAB", k = 2

Output:
4

Explanation:
Replace the two 'A's with two 'B's or vice versa.
```

 

**Example 2:**

```
Input:
s = "AABABBA", k = 1

Output:
4

Explanation:
Replace the one 'A' in the middle with 'B' and form "AABBBBA".
The substring "BBBB" has the longest repeating letters, which is 4.
```

 ## Solution

- 滑动窗口：
  - 由于可以替换 k 个字符，故窗口中最多的元素为任意多个重复元素 + k 个不重复元素
  - 由于 k 个不重复元素是固定的，故重复元素的个数越多，最终的结果越大
  - 每轮循环右边界右移，向窗口中加入一个元素，同时记录下加入元素在当前窗口中的重复次数，并使用该次数来更新最大重复元素的个数 maxCount
  - 窗口收缩的条件为：`end - start + 1 - maxCount > k`，此时，将 start 位置的元素计数去除，同时将 start 右移 
  - 时间复杂度：$O(n)$，空间复杂度：$O(1)$（因为只有 26 个大写字母，故辅助空间的大小是固定的，不随数据规模变化而变化）