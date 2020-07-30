# 118. Pascal's Triangle

## Problem

Given a non-negative integer *numRows*, generate the first *numRows* of Pascal's triangle.

**Example:**

```
Input: 5
Output:
[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
]
```

## Solution

- 杨辉三角形的特点：

  - 每一行的首尾两个元素都是1；
  - 每一行的元素个数是i + 1；
  - 每一行中除了首末元素外，剩下的元素都满足 `a[i][j] = a[i - 1][j - 1] + a[i - 1][j]`

- 时间复杂度：$O(n^2)$，空间复杂度：$O(1)$

  