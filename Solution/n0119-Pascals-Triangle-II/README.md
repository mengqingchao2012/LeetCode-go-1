# 119. Pascal's Triangle II
## Problem

Given a non-negative index *k* where *k* ≤ 33, return the *k*th index row of the Pascal's triangle.

Note that the row index starts from 0.

**Example:**

```
Input: 3
Output: [1,3,3,1]
```

**Follow up:**

Could you optimize your algorithm to use only *O*(*k*) extra space?

## Solution

- 思路：

  - 由于只需要返回最后一行的结果，因此可以在同一个数组中进行修改

  - 需要注意的是，这样修改数组的时候要从右往左遍历

  - 时间复杂度：$O(n^2)$，空间复杂度：$O(1)$

    