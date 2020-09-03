# 340. Longest Substring with At Most K Distinct Characters
## Problem

Given a string, find the length of the longest substring T that contains at most k distinct characters.

Example 1:

```
Input: s = "eceba", k = 2
Output: 3
Explanation: T is "ece" which its length is 3.
```

Example 2:

```
Input: s = "aa", k = 1
Output: 2
Explanation: T is "aa" which its length is 2.
```

## Solution

- 滑动窗口：
  - 使用一个辅助哈希表来记录当前窗口中的字符
  - 遍历字符串，在窗口中元素小于等于 k 之前，移动右边界扩大窗口，同时将新加入的字符添加到哈希表中计数
  - 只要窗口中元素大小超过 k，则右移左边界缩小窗口，每次缩小，都相应的从哈希表中减少对应字符的计数，如果计数减为 0，则删去该条目
  - 只要窗口中元素大小小于等于 k，则每轮循环都使用当前窗口中的元素个数来更新最终结果
- 时间复杂度：$O(n)$，空间复杂度：$O(n)$