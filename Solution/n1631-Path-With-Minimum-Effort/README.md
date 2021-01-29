# 1631. Path With Minimum Effort
## Problem

You are a hiker preparing for an upcoming hike. You are given `heights`, a 2D array of size `rows x columns`, where `heights[row][col]` represents the height of cell `(row, col)`. You are situated in the top-left cell, `(0, 0)`, and you hope to travel to the bottom-right cell, `(rows-1, columns-1)` (i.e., **0-indexed**). You can move **up**, **down**, **left**, or **right**, and you wish to find a route that requires the minimum **effort**.

A route's **effort** is the **maximum absolute difference** in heights between two consecutive cells of the route.

Return *the minimum **effort** required to travel from the top-left cell to the bottom-right cell.*

 

**Example 1:**

<img src="..\..\\pic\lc1631-1.png" alt="avatar" style="zoom:67%;" />

```
Input: heights = [[1,2,2],[3,8,2],[5,3,5]]
Output: 2
Explanation: The route of [1,3,5,3,5] has a maximum absolute difference of 2 in consecutive cells.
This is better than the route of [1,2,2,2,5], where the maximum absolute difference is 3.
```

**Example 2:**

<img src="..\..\\pic\lc1631-2.png" alt="avatar" style="zoom:67%;" />

```
Input: heights = [[1,2,3],[3,8,4],[5,3,5]]
Output: 1
Explanation: The route of [1,2,3,4,5] has a maximum absolute difference of 1 in consecutive cells, which is better than route [1,3,5,3,5].
```

**Example 3:**

<img src="..\..\pic\lc1631-3.png" alt="avatar" style="zoom:67%;" />

```
Input: heights = [[1,2,1,1,1],[1,2,1,2,1],[1,2,1,2,1],[1,2,1,2,1],[1,1,1,2,1]]
Output: 0
Explanation: This route does not require any effort.
```

 

**Constraints:**

- `rows == heights.length`
- `columns == heights[i].length`
- `1 <= rows, columns <= 100`
- `1 <= heights[i][j] <= 10^6`

## Solution

- 解法一：二分 + bfs
  - 高度差的范围是 [0, 1e6]，故可以假设当前最大高度差为 i，然后检查在这个高度差的约束下能否从起点走到终点。
  - 如果能，则使用二分法将 i 减小，继续检查，直到找到最小的高度差，能够满足从起点走到终点，则这个高度差就是最终的结果
  - 时间复杂度：$O(mnlogC)$，m 和 n 是矩阵的行数和列数，C 是格子的最大高度
  - 空间复杂度：$O(mn)$，广度优先搜索开的辅助队列的大小
- 解法二：贪心 + 并查集
  - 找出所有消耗体力为 i 的路径，将他们加入同一个连通集中，检查这个连通集能否满足从起点走到终点，如果不能，则查找下一个可能连通集，直到找到结果
  - 使用一个三元组来表示一条边：（这条边的高度差，端点 1，端点2）
  - 遍历矩阵，对于矩阵中的每个点(x, y)，我们只构造它和它右侧相邻的点（x, y + 1)，下方相邻的点（x + 1, y）组成的两条边，这样遍历矩阵后就能够不重不漏的得到该矩阵中所有可能组成的邻边。将这些边存到一个数组中
  - 将这些边按照高度差从小到大排序，然后依次将排好序的边加入到并查集中。每加入一条边就检查起点和终点是否满足在同一个集合中，如果满足，则说明找到了消耗体力最小的结果
  - 时间复杂度：近似为$O(mnlogmn)$，其中，图中所有的边的数量为 mn，故排序时间复杂度为 mnlogmn，而并查集近似为 O(1)
  - 空间复杂度：$O(mn)$