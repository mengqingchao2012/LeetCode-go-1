# 169. Majority Element

## Problem
Given an array of size n, find the majority element. The majority element is the element that appears more than ⌊ n/2 ⌋ times.

You may assume that the array is non-empty and the majority element always exist in the array.

Example 1:

Input: [3,2,3]
Output: 3
Example 2:

Input: [2,2,1,1,1,2,2]
Output: 2

## Solution
- 摩尔投票法：
    众数只有一个，记下当前数，并初始化计数器为1，遍历数组，如果下一个数与当前数相同，则计数器加1，否则计数器减一，
    并更新当前数，如果计数器减到了0，则重新初始化计数器为1，并记录下当前数
- 时间复杂度：O(n)，空间复杂度：O(1)