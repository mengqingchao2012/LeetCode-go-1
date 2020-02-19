# 34. Find First and Last Position of Element in Sorted Array

## Problem
Given an array of integers nums sorted in ascending order, find the starting and ending position of a given target value.

Your algorithm's runtime complexity must be in the order of O(log n).

If the target is not found in the array, return [-1, -1].

Example 1:

Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]
Example 2:

Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]

## Solution
- 二分法的特点：如果查找的元素不在数组中，则返回的下标是该元素应该插入的位置的下标
    - 如果 目标值 < 中间值，则修正右边界，目标值 == 中间值，直接返回，目标值 > 中间值，修正左边界
- 要找元素最后一次出现的位置，只需要在目标值等于中间值时，让其继续往右查找，也就是修正左边界
- 思路：
    可以直接通过二分法的变体找到元素最后一次出现的位置
    元素的开始位置可以通过求比目标小1的元素的插入位置，再加上1得到，不用考虑比目标小1的元素是否存在数组中，返回的是插入位置
- 时间复杂度：O(log(n))，空间复杂度：O(1)