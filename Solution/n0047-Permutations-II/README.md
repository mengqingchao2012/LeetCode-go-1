# 47. Permutations II

## Problem

Given a collection of numbers that might contain duplicates, return all possible unique permutations.

**Example:**

```
Input: [1,1,2]
Output:
[
  [1,1,2],
  [1,2,1],
  [2,1,1]
]
```

## Solution

- 借用第31题 Next Permutation 的思路，每次获取比原数组大的下一个数组，即可实现去重

- 注意 go 中 slice 的底层数组引用会带来非常多的问题=-=

- 时间复杂度：$O(n * n!)$，空间复杂度：$O(1)$

  