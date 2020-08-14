# 152. Maximum Product Subarray
## Problem

Given an integer array nums, find the contiguous subarray within an array (containing at least one number) which has the largest product.

Example 1:

```
Input: [2,3,-2,4]
Output: 6
Explanation: [2,3] has the largest product 6.
```

Example 2:

```
Input: [-2,0,-1]
Output: 0
Explanation: The result cannot be 2, because [-2,-1] is not a subarray.
```

## Solution

- 类似于第 53 题
- 不同之处在于，这里是乘法，有几个注意点：
  - 乘以一个负数会将最大值变成最小值，将最小值变成最大值，因此需要同时记录每一轮的最大值和最小值
  - 乘以 0 会导致所有值都变为 0
- 其余见代码