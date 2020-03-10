# 121. Best Time to Buy and Sell Stock

## Problems

Say you have an array for which the *i*th element is the price of a given stock on day *i*.

If you were only permitted to complete at most one transaction (i.e., buy one and sell one share of the stock), design an algorithm to find the maximum profit.

Note that you cannot sell a stock before you buy one.

**Example 1:**

```
Input: [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
             Not 7-1 = 6, as selling price needs to be larger than buying price.
```

**Example 2:**

```
Input: [7,6,4,3,1]
Output: 0
Explanation: In this case, no transaction is done, i.e. max profit = 0.
```



## Solution

- 思路：

  - 使用两个变量 buy 和 profile 来分别存储买入的价格和最大利润；
  - 遍历数组，初始化买入价格为第一个价格元素，而后如果出现当前价格低于买入价格的，则该轮就更新买入价格，否则计算并更新最大利润

- 时间复杂度：$O(n)$，空间复杂度：$O(1)$

  