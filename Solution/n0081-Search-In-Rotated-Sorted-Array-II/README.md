# 81. Search in Rotated Sorted Array II

## Problem
Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e., `[0,0,1,2,2,5,6]` might become `[2,5,6,0,0,1,2]`).

You are given a target value to search. If found in the array return `true`, otherwise return `false`.

**Example 1:**

```
Input: nums = [2,5,6,0,0,1,2], target = 0
Output: true
```

**Example 2:**

```
Input: nums = [2,5,6,0,0,1,2], target = 3
Output: false
```

**Follow up:**

- This is a follow up problem to [Search in Rotated Sorted Array](https://leetcode.com/problems/search-in-rotated-sorted-array/description/), where `nums` may contain duplicates.
- Would this affect the run-time complexity? How and why?

## Solution
- 解法与 33 题类似。
- 元素可以重复带来的影响：
    - 数组的二段性不复存在，即：无法保证前半部分中的元素满足：`nums[0] < nums[i]`，后半部分满足：`nums[i] < nums[0]`，此时对于 `nums[i] >= nums[0]`，无法确定其处于前半段还是后半段
    - 如： [1,0,1,1,1] 和 [1,1,1,0,1]
    - 因为题目只需要判断是否存在，此时可以在首尾元素相同时，选择删掉任意一边的元素，使得数组恢复二段性
    - 如果要求下标，则考虑在出现 `nums[start] == nums[mid] && nums[mid] == nums[end]` 时同时更新 start++，end--
- 时间复杂度：O(logn)，空间复杂度：O(1)
- 最坏时间复杂度是 O(n)，比如所有元素都相同时，会持续的执行删除元素那一步