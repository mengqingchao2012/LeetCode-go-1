# 77. Combinations

## Problem

Given two integers *n* and *k*, return all possible combinations of *k* numbers out of 1 ... *n*.

**Example:**

```
Input: n = 4, k = 2
Output:
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]
```

## Solution

- 回溯算法；

- 时间复杂度：$O(k*C_n^k)$：要找 k 个数，每个数有 $C_n^k$ 中取法

- 空间复杂度：$O(k)$：包括递归的深度 k，和临时存储的数组空间 k

  