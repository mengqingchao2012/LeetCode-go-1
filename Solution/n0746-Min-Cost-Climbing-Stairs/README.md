# 746. Min Cost Climbing Stairs

## Problems

On a staircase, the `i`-th step has some non-negative cost `cost[i]` assigned (0 indexed).

Once you pay the cost, you can either climb one or two steps. You need to find minimum cost to reach the top of the floor, and you can either start from the step with index 0, or the step with index 1.

**Example 1:**

```
Input: cost = [10, 15, 20]
Output: 15
Explanation: Cheapest is start on cost[1], pay that cost and go to the top.
```



**Example 2:**

```
Input: cost = [1, 100, 1, 1, 1, 100, 1, 1, 100, 1]
Output: 6
Explanation: Cheapest is start on cost[0], and only step on 1s, skipping cost[3].
```



**Note:**

1. `cost` will have a length in the range `[2, 1000]`.
2. Every `cost[i]` will be an integer in the range `[0, 999]`.

## Solutions

- 思路：动态规划

  - 爬到第 i 阶所消耗的最小体力，等于第 i - 1 阶和第 i - 2 阶中体力消耗较小的那个值加上第 i 阶的体力消耗值

- 时间复杂度：$O(n)$，空间复杂度：$O(1)$

  