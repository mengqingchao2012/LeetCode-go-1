# 268. Missing Number
## Problem

Given an array containing *n* distinct numbers taken from `0, 1, 2, ..., n`, find the one that is missing from the array.

**Example 1:**

```
Input: [3,0,1]
Output: 2
```

**Example 2:**

```
Input: [9,6,4,2,3,5,7,0,1]
Output: 8
```

**Note**:
Your algorithm should run in linear runtime complexity. Could you implement it using only constant extra space complexity?

## Solution

- 方法一：Cyclic Sort
  - 对于有序数组，只需要遍历数组，检查 nums[i] 是否等于 i 即可
  - 对于无序数组，可以先排序，由于数组中存储的是 0-n 之间的数，可以在 O(n) 时间内完成排序：
    - 读出下标 i 对应的元素：nums[i]，只要 nums[i] 不等于 [i]，就将 nums[i] 和 nums[nums[i]] 的值进行交换，直到相等，则继续下一个元素
    - 需要注意的是：数据来源是 [0, n]，而数组长度只有 n - 1，故需要判断下标是否合法
  - 排序完后再遍历一次就可知道结果
  - 时间复杂度：$O(n)$，空间复杂度：$O(1)$
- 方法二：异或运算
  - 已知 a ^ a = 0，同时有 (a ^ b) ^ c = a ^ (b ^ c)，故将数组中的每个元素分别异或上其下标，并将每组结果异或起来，最终得到的就是缺失的元素值