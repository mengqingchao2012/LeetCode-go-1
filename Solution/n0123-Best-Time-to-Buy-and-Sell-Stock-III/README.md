# 123. Best Time to Buy and Sell Stock III
## Problem

Say you have an array for which the *i*th element is the price of a given stock on day *i*.

Design an algorithm to find the maximum profit. You may complete at most *two* transactions.

**Note:** You may not engage in multiple transactions at the same time (i.e., you must sell the stock before you buy again).

 

**Example 1:**

```
Input: prices = [3,3,5,0,0,3,1,4]
Output: 6
Explanation: Buy on day 4 (price = 0) and sell on day 6 (price = 3), profit = 3-0 = 3.
Then buy on day 7 (price = 1) and sell on day 8 (price = 4), profit = 4-1 = 3.
```

**Example 2:**

```
Input: prices = [1,2,3,4,5]
Output: 4
Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
Note that you cannot buy on day 1, buy on day 2 and sell them later, as you are engaging multiple transactions at the same time. You must sell before buying again.
```

**Example 3:**

```
Input: prices = [7,6,4,3,1]
Output: 0
Explanation: In this case, no transaction is done, i.e. max profit = 0.
```

**Example 4:**

```
Input: prices = [1]
Output: 0
```

 

**Constraints:**

- `1 <= prices.length <= 105`
- `0 <= prices[i] <= 105`

## Solution

- 方法一：前后缀分离
  - 由题意可知，最多只能交易两次，且前后两次交易之间是相互独立的。因此可以通过枚举第二次交易的买入点，将数组划分为两部分，前半部分是第一次交易，后半部分是第二次交易，两次交易求最大利润的方法是相同的；
  - 创建一个映射（这里使用数组），f[i] 表示 [0, i] 范围内完成一次交易可得到的最大利润
  - 时间复杂度：$O(n)$，空间复杂度：$O(n)$
- 方法二：动态规划 