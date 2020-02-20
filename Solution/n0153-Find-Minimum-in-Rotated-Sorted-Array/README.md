# 153. Find Minimum in Rotated Sorted Array

## Problem
Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e.,  [0,1,2,4,5,6,7] might become  [4,5,6,7,0,1,2]).

Find the minimum element.

You may assume no duplicate exists in the array.

Example 1:

Input: [3,4,5,1,2] 
Output: 1
Example 2:

Input: [4,5,6,7,0,1,2]
Output: 0

## Solution
- 数组是由有序数组（递增）旋转后得到，故最小值就是旋转点，旋转后的数组模式为[递增|旋转点|递增]
- 使用二分法，求出 mid，如果 nums[mid] > nums[high]，说明当前 mid 位于前半递增区间内，最小值肯定不在 mid 左半边，故更新
low，如果 nums[mid] <= nums[high] 时，只能将 high 更新为 mid，因为有可能 mid 就是当前的旋转点，也就是最小值。
- 循环退出的条件是 low < high，注意此时不能取等号，退出时，low - high 区间内只有一个唯一元素，就是旋转点，所以直接返回 nums[low]
  <img src="..\..\pic\lc153.png" alt="avatar" style="zoom:67%;" />
- 优化：由分析只，low 指针是一直单向往右移动的，所以在每次进入与循环时，通过判断 nums[low] 是否已经小于 nums[high]，如果成立就可以
提前返回 nums[low]
- 时间复杂度：O(logn)，空间复杂度：O(1)