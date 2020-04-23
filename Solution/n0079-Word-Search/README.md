# 79. Word Search

## Problem

Given a 2D board and a word, find if the word exists in the grid.

The word can be constructed from letters of sequentially adjacent cell, where "adjacent" cells are those horizontally or vertically neighboring. The same letter cell may not be used more than once.

**Example:**

```
board =
[
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]

Given word = "ABCCED", return true.
Given word = "SEE", return true.
Given word = "ABCB", return false.
```

 

**Constraints:**

- `board` and `word` consists only of lowercase and uppercase English letters.
- `1 <= board.length <= 200`
- `1 <= board[i].length <= 200`
- `1 <= word.length <= 10^3`

## Solution

- 见代码
- 时间复杂度：$O(m*n*3^k)$
  - 最坏情况下至少需要遍历一次数组，所以最少也有 m*n
  - 对每个位置深入搜索，一开始有4个方向可以选择，之后每深入一步都有3个方向可以选择（有一个方向是重复位置），要深入检查k次，所以上限是 3的k次方
  - 乘起来即得到时间复杂度的上限
- 空间复杂度：$O(m*n)$，使用了一个二维辅助数组进行标记