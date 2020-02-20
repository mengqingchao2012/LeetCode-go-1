# 240. Search a 2D Matrix II

## Problem
Write an efficient algorithm that searches for a value in an m x n matrix. This matrix has the following properties:

Integers in each row are sorted in ascending from left to right.
Integers in each column are sorted in ascending from top to bottom.
Example:

Consider the following matrix:

[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]
Given target = 5, return true.

Given target = 20, return false.

## Solution
- 从右上角开始，如果目标值小于当前值，则说明目标值比当前列的所有值都小，所以列左移，如果目标值大于当前值，说明目标值比当前行的所有值
都大，行下移；
- 注意各个边界的条件限制
- 时间复杂度：O(m+n)，空间复杂度：O(1)