# 50. Pow(x, n)

## Problem
Implement pow(x, n), which calculates x raised to the power n (xn).

Example 1:

Input: 2.00000, 10
Output: 1024.00000
Example 2:

Input: 2.10000, 3
Output: 9.26100
Example 3:

Input: 2.00000, -2
Output: 0.25000
Explanation: 2-2 = 1/22 = 1/4 = 0.25
Note:

1. -100.0 < x < 100.0
2. n is a 32-bit signed integer, within the range [−231, 231 − 1]

## Solution
- 将幂的值写成二进制的形式：

  - 如：11 = 1011，则：Pow(x, 11) = $x^{11}=x^8*x^2*x^1$

    <img src="..\..\pic\lc50.png" alt="avatar" style="zoom:80%;" />

  - x 的幂次刚好对应到 11 转为二进制后其不为0值的位

  - 使用一个循环，当幂次 n 不等于0时，如果 n 的最低位是1，就用底数 x 乘以 n，然后更新底数：`x*=x`，如果n的最低位是0，则该轮只更新底数，每一轮都将 n 右移一位，故每一轮 n 的值都减半；

- 时间复杂度：O(logn)，空间复杂度：O(1)