# 209. Minimum Size Subarray Sum
## Problem

Given an array of **n** positive integers and a positive integer **s**, find the minimal length of a **contiguous** subarray of which the sum ≥ **s**. If there isn't one, return 0 instead.

**Example:** 

```
Input: s = 7, nums = [2,3,1,2,4,3]
Output: 2
Explanation: the subarray [4,3] has the minimal length under the problem constraint.
```

## Solution

- 滑动窗口：
  - 令窗口范围内的元素和为 winSum，则当 winSum < s 时，不断增大窗口，即窗口右边界右移；
  - 当满足 winSum >= s 时，找到满足条件的结果，这时使用窗口中元素的个数来更新最终结果 res
  - 同时，需要在当前窗口中查找是否可能出现更优解，即，只要当前窗口内元素和还满足 winSum >= s，就不断的右移左边界，同时用新窗口中元素的个数来更新最终结果
- 时间复杂度：$O(n + n)$，空间复杂度：$O(1)$