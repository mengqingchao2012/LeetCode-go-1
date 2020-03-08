# Coin Change

## Problems

You are given coins of different denominations and a total amount of money *amount*. Write a function to compute the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return `-1`.

**Example 1:**

```
Input: coins = [1, 2, 5], amount = 11
Output: 3 
Explanation: 11 = 5 + 5 + 1
```

**Example 2:**

```
Input: coins = [2], amount = 3
Output: -1
```

**Note**:
You may assume that you have an infinite number of each kind of coin.



## Solution

- 思路：
  - 还是将情况分为针对指定的金额，使用指定硬币和不使用指定硬币两种情况
  - 要求的是凑成金额的最小硬币数，则结果就是上面两种情况中较小的那个
- 时间复杂度：$O(amount * len(coin))$；空间复杂度：$O(amount)$

