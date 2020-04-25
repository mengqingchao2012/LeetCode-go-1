# 560. Subarray Sum Equals K

## Problem

Given an array of integers and an integer **k**, you need to find the total number of continuous subarrays whose sum equals to **k**.

**Example 1:**

```
Input:nums = [1,1,1], k = 2
Output: 2
```



**Note:**

1. The length of the array is in range [1, 20,000].
2. The range of numbers in the array is [-1000, 1000] and the range of the integer **k** is [-1e7, 1e7].

## Solution

- 方法一：暴力法：

  - 两层循环，外层循环 i 确定子数组的左边界，内层循环 j 每次从 i 开始，确定子数组的右边界，即组成子数组为 [i, j]
  - 使用变量 sum 来保存边界内子数组的和，如果和等于 k，就将计数器加一
  - 时间复杂度：$O(n^2)$，空间复杂度：$O(1)$

- 方法二：前缀和：

  - 和区间 S[i, j] 等效于 S[0, j] - s[0, i - 1]，S[0, j]  和 S[0, i - 1] 称为前缀和，由于都是从下标 0 开始计算到特定值，所以只使用一个下标 S[j] 和 S[i-1] 来表示前缀和；

  - 所以有：$k = S[j] - S[i - 1] \Rightarrow S[j] - k = S[i - 1]$

  - 问题降维到类似于求两数之和的问题，使用一个辅助哈希表来存储到下标 i 的前缀和，键为前缀和，值为前缀和出现的次数；

    - 注意之所以要记录前缀和出现的次数，是因为可能出现同一个前缀和由不同的区间构造组成
    - 如：[1, 2, 1, -1] 中，[1, 2] 构成前缀和 3，[1, 2, 1, -1] 也构成前缀和 3

  - 如果当前下标 j 的前缀和减去 k 的值可以在哈希表中查到，就说明在 S[i, j] 之间的和满足条件

  - 一个隐含的条件是：当 S[0, j] = k 时，此时 S[i - 1] = 0，故需要将（0，1）这一组值提前加入到哈希表中

  - 时间复杂度：$O(n)$，空间复杂度：$O(n)$

    