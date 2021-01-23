# 51. N-Queens
## Problem

The **n-queens** puzzle is the problem of placing `n` queens on an `n x n` chessboard such that no two queens attack each other.

Given an integer `n`, return *all distinct solutions to the **n-queens puzzle***.

Each solution contains a distinct board configuration of the n-queens' placement, where `'Q'` and `'.'` both indicate a queen and an empty space, respectively.

 

**Example 1:**

![avatar](..\..\pic\lc51.jpg)

```
Input: n = 4
Output: [[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
Explanation: There exist two distinct solutions to the 4-queens puzzle as shown above
```

**Example 2:**

```
Input: n = 1
Output: [["Q"]]
```

 

**Constraints:**

- `1 <= n <= 9`

## Solution

- N 皇后问题，dfs

- 时间复杂度：

  - n 个皇后不能在同行同列同对角线，考虑列方向，第一行有 n 个位置可选，第二行有 n - 1 个，以此类推得到 `n * (n - 1) * (n - 2) * ... * 1 = n!`种可选的方案。由于对角线方向的影响远不如列方向的影响来的大，故可以直接使用 $$O(n!)$$ 作为一个宽松的上界

- 空间复杂度：有三个部分：

  - 递归深度：有 n 行，故为 n
  - 辅助数组：n
  - 棋盘的大小 n * n
  - 故空间复杂度为 $O(n^2)$

   