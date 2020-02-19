# Two Sum

## Problem
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].

## Solution
- 思路：t1 + t2 = target => t1 = target - t2
- 实现：遍历数组 nums，使用一个辅助 map ，键是元素的值，值是元素对应的下标。如果目标值与当前下标所指元素的差在 map 中已经存在，
则说明该数组中存在两数之和等于目标值，如果不存在，则将元素及其下标存入数组中继续遍历。
- 时间复杂度：$O(n)$，空间复杂度：$O(n)$