# 154. Find Minimum in Rotated Sorted Array II

## Problem

Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e.,  `[0,1,2,4,5,6,7]` might become  `[4,5,6,7,0,1,2]`).

Find the minimum element.

The array may contain duplicates.

**Example 1:**

```
Input: [1,3,5]
Output: 1
```

**Example 2:**

```
Input: [2,2,2,0,1]
Output: 0
```

**Note:**

- This is a follow up problem to [Find Minimum in Rotated Sorted Array](https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/description/).
- Would allow duplicates affect the run-time complexity? How and why?

## Solution

- 思路与 153 题相同，问题在于，存在重复值时会影响对分区的判断
- 抓住数组是由原递增数组旋转得到这个条件，把关注点放在旋转后的部分
  - 求出 mid 后，如果 `nums[mid] < nums[high]`，说明落在了正确的分区里，只需修正 `high = mid` 即可
  - 如果 `nums[mid] > nums[high]`，说明落在了错误的分区里，修正 `low = mid + 1`
  - 重复值造成的影响：`[1, 3, 3, 3] -> [3, 3, 1, 3]`，求出的 `nums[mid] == nums[high]`，此时选择修正 `high` 为 `high--`
- 平均时间复杂度：$O(logn)$，空间复杂度：$O(1)$
- 最坏时间复杂度：$O(n)$，出现在所有值都相同时，要遍历完整数组