# 90. Subsets II
## Problem

Given a collection of integers that might contain duplicates, ***nums\***, return all possible subsets (the power set).

**Note:** The solution set must not contain duplicate subsets.

**Example:**

```
Input: [1,2,2]
Output:
[
  [2],
  [1],
  [1,2,2],
  [2,2],
  [1,2],
  []
]
```

## Solution

- 回溯，思路与 78 题相同
- 为了去掉重复的结果，需要先将数组进行排序，对于每个位置，只能往后取数，不能往前取，同时需要跳过重复的元素