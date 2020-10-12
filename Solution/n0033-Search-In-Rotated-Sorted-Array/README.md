# 33. Search in Rotated Sorted Array

## Problem
Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).

You are given a target value to search. If found in the array return its index, otherwise return -1.

You may assume no duplicate exists in the array.

Your algorithm's runtime complexity must be in the order of O(log n).

Example 1:

```
Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4
```

Example 2:

```
Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1
```

## Solution
- 由题意，旋转后的数组被分成了两个部分，以旋转点为边界分成了两个递增数组，因此依旧可以使用二分法。
    - 即，数组具有二段性：前半部分中的元素满足：`nums[0] < nums[i]`，后半部分满足：`nums[i] < nums[0]`
- 二分法的关键在于要找到有序的区间，所以在每次取到中间值是，要判断中间值所在的区间和目标值所在的区间是否是同一个：
    - 判断的方法就是：利用数组的有序性，如果 `nums[mid] >= nums[low]`，则此时 mid 在前半区间，
    如果此时target 满足 `target < nums[mid] && target >= nums[low]`，即可说明 target 和 `nums[mid]` 在同一个区间内
- 时间复杂度：O(logn)， 空间复杂度：O(1)