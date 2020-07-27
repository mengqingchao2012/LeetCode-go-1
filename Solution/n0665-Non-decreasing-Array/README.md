# 665. Non-decreasing Array
## Problem

Given an array `nums` with `n` integers, your task is to check if it could become non-decreasing by modifying **at most** `1` element.

We define an array is non-decreasing if `nums[i] <= nums``[i + 1]` holds for every `i` (0-based) such that `(0 <= i <= n - 2)`.

 

**Example 1:**

```
Input: nums = [4,2,3]
Output: true
Explanation: You could modify the first 4 to 1 to get a non-decreasing array.
```

**Example 2:**

```
Input: nums = [4,2,1]
Output: false
Explanation: You can't get a non-decreasing array by modify at most one element.
```

 

**Constraints:**

- `1 <= n <= 10 ^ 4`
- `- 10 ^ 5 <= nums[i] <= 10 ^ 5`

## Solution

- 思路：
  - 需要记录数组中相邻元素出现递减情况的次数
  - 关键点有两点：
    - 递减出现的位置有哪些情况？
      - 在数组刚开始的前两个元素之间出现：[4, 2, 3]
      - 在数组中间出现：[3,4,2,3]
    - 修改时是将第一个元素改小还是将第二个元素改大？
      - 对于第一种情况，只能修改第一个元素，将其改小，否则会出错，如：[4, 3, 3]
      - 对于第二种情况，需要再分类讨论：
        - 如果 `nums[i + 1] >= nums[i - 1]`，则需要将第一个元素（nums[i]）改小
          - 如：[2, 3, 5, 4, 6]，改为 [2, 3, 3, 4, 6]
        - 如果 `nums[i + 1] < nums[i - 1]`，则需要将第二个元素（nums[i + 1]）改大
          - 如：[2, 3, 4, 1, 6]，改为 [2, 3, 4, 4, 6]
  - 由于涉及到 n - 1，故下标只能从 1 开始，即需要单独对前两个元素进行判断
  - 时间复杂度：$O(n)$，空间复杂度：$O(1)$