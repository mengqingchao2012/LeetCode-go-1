# 81. Search in Rotated Sorted Array II

## Problem
Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e., [0,0,1,2,2,5,6] might become [2,5,6,0,0,1,2]).

You are given a target value to search. If found in the array return true, otherwise return false.

Example 1:

Input: nums = [2,5,6,0,0,1,2], target = 0
Output: true
Example 2:

Input: nums = [2,5,6,0,0,1,2], target = 3
Output: false
Follow up:

This is a follow up problem to Search in Rotated Sorted Array, where nums may contain duplicates.
Would this affect the run-time complexity? How and why?

## Solution
- 解法同 33 题。
- 元素可以重复带来的影响：
    在出现[1,0,1,1,1]和[1,1,1,0,1]这种情况时，会出现nums[2]==nums[0]的情况，此时无法分辨出到底哪里才是有序区间[1,0]还是[0,1,1,1]，
    由于题目只需判断是否存在，故可以选择直接将low++，去掉干扰项
- 时间复杂度：O(logn)，空间复杂度：O(1)