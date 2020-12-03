# 204. Count Primes
## Problem

Count the number of prime numbers less than a non-negative number, `n`.

 

**Example 1:**

```
Input: n = 10
Output: 4
Explanation: There are 4 prime numbers less than 10, they are 2, 3, 5, 7.
```

**Example 2:**

```
Input: n = 0
Output: 0
```

**Example 3:**

```
Input: n = 1
Output: 0
```

 

**Constraints:**

- `0 <= n <= 5 * 106`

## Solution

- 质数的相关性质：
  - 定义：从 2 开始，如果一个数只有1 和他本身两个约数，则这个数是质数。小于 2 的数既不是质数，也不是合数
  - 对于一个数 n，最多只有一个大于 $\sqrt{n}$ 的质因子
  - 调和级数：$1 + \frac{1}{2} + \frac{1}{3} +... +\frac{1}{n} = lnn + C$
  - 质数定理：在 1-n 中有 $\frac{n}{lnn}$ 个质数
- 解法一：朴素筛
  - 将 [2, n) 个数放入数组中，使用一个辅助数组 `is_composite` 来存储第 i 个数是不是质数
  - 对于第 i 个数，如果 `is_composite[i]` 为 false，则说明当前数是质数，则将计数加一
  - 同时对于遍历到的第 i 个数，分别将其倍数标记为 true
  - 时间复杂度：
    - 运算次数：$\frac{n}{2} + \frac{n}{3} +... +\frac{n}{n} \Rightarrow nlnn$
    - 时间复杂度：$nlnn = nlog_en < nlog_2n = O(nlogn)$
- 解法二：埃氏筛
  - 在解法一的基础上进行优化，只有在元素 i 是质数的情况下，才将其倍数标记为 true
  - 原理：对于元素 p，只要 p 不是 [2, p - 1] 中质数的倍数，即 p 在 [2, p - 1] 中没有其他质因数，则 p 一定是质数，因为只有 1 和 p 两个约数
  - 时间复杂度：$O(nloglogn)$