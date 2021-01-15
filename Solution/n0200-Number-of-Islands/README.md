# 200. Number of Islands

## Problem

Given an `m x n` 2d `grid` map of `'1'`s (land) and `'0'`s (water), return *the number of islands*.

An **island** is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.

 

**Example 1:**

```
Input: grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
Output: 1
```

**Example 2:**

```
Input: grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
Output: 3
```

 

**Constraints:**

- `m == grid.length`
- `n == grid[i].length`
- `1 <= m, n <= 300`
- `grid[i][j]` is `'0'` or `'1'`.

## Solution 

- 思路：flood fill 算法——洪水灌溉算法
  - 遍历 grid，使用一个辅助二维数组来标记访问过的位置
  - 如果当前元素是 "1"，则使用 dfs 或者 bfs 将元素上下左右四个方向上所有满足条件的位置都进行标记，再继续遍历后续元素
- 时间复杂度：$O(m * n)$，m 和 n 是图的大小，因为每个元素最多只会被访问两次（一次正常遍历，一次检查标记），所以是线性时间复杂度