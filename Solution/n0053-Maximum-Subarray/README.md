# 53. Maximum Subarray

## Problem

Given an integer array `nums`, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

**Example:**

```
Input: [-2,1,-3,4,-1,2,1,-5,4],
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
```

**Follow up:**

If you have figured out the O(*n*) solution, try coding another solution using the divide and conquer approach, which is more subtle.

## Solution

- 思路：

  - 用 max 变量表示最大和，cur 变量表示当前和
  - 如果 cur <= 0，则说明当前 cur 对整体最大和完全是负贡献，还不直接丢弃前面的累加和重新开始，故此时使用当前元素值重置 cur，其他情况将当前元素值累加到 cur 上，并用 cur 更新 max

- 时间复杂度：$O(n)$，空间复杂度：$O(1)$

  