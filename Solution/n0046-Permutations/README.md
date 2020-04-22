# 46. Permutations

## Problems

Given a collection of **distinct** integers, return all possible permutations.

**Example:**

```
Input: [1,2,3]
Output:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]
```

## Solutions

- 回溯法

  - 对数组中的每个元素，将其置换到第一个元素，则该元素的全排列为后面子数组元素的全排列

- 时间复杂度：$O(n * n!)$，空间复杂度：$O(n)$

  