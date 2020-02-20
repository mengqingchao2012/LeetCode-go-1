# 74. Search a 2D Matrix

##Problem
Write an efficient algorithm that searches for a value in an m x n matrix. This matrix has the following properties:

Integers in each row are sorted from left to right.
The first integer of each row is greater than the last integer of the previous row.
Example 1:

Input:
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 3
Output: true
Example 2:

Input:
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 13
Output: false

## Solution
- 由题意可以将该矩阵看成是一个相邻两行之间呈z字形递增的数组，可以使用二分法
- 求出 mid 后，要将 mid 转换成行列坐标表示的形式：
    行坐标 = mid / 列数
    列坐标 = mid % 列数
- 时间复杂度：O(log(m*n))，空间复杂度：O(1)