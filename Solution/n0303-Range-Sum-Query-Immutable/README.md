# 303. Range Sum Query - Immutable
## Problem

Given an integer array *nums*, find the sum of the elements between indices *i* and *j* (*i* ≤ *j*), inclusive.

**Example:**

```
Given nums = [-2, 0, 3, -5, 2, -1]

sumRange(0, 2) -> 1
sumRange(2, 5) -> -1
sumRange(0, 5) -> -3
```

 

**Constraints:**

- You may assume that the array does not change.
- There are many calls to *sumRange* function.
- `0 <= nums.length <= 10^4`
- `-10^5 <= nums.length <= 10^5`
- `0 <= i <= j < nums.length`

## Solution

- 见代码