# 523. Continuous Subarray Sum

## Problem

Given a list of **non-negative** numbers and a target **integer** k, write a function to check if the array has a continuous subarray of size at least 2 that sums up to a multiple of **k**, that is, sums up to n*k where n is also an **integer**.

 

**Example 1:**

```
Input: [23, 2, 4, 6, 7],  k=6
Output: True
Explanation: Because [2, 4] is a continuous subarray of size 2 and sums up to 6.
```

**Example 2:**

```
Input: [23, 2, 6, 4, 7],  k=6
Output: True
Explanation: Because [23, 2, 6, 4, 7] is an continuous subarray of size 5 and sums up to 42.
```

 

**Note:**

1. The length of the array won't exceed 10,000.
2. You may assume the sum of all the numbers is in the range of a signed 32-bit integer.

## Solution

- 同余：如果两个数对同一个数取余的结果相同，则说明这两个数同余数，即`a % k == b % k`

- 如果两个数对 k 同余，则这两个数的差是 k 的整数倍，即：$a \% k == b \% k \Rightarrow a - b = m * k$

  - 看到 a - b 这种形式，应该想到前缀和

- 思路：

  - 使用一个哈希表来存储前缀和，键是前缀和对k的余数，值是该前缀和对应的下标
  - 储存下标是因为题目要求结果数组的长度要大于等于2
  - 要注意两个特殊情况：
    - 如果 k 为 0，则不能对前缀和取余，直接保存前缀和
    -  初始化时要将 (0, -1) 加入哈希表

- 时间复杂度：$O(n)$，空间复杂度：$O(n)$

  