# 547. Number of Provinces
## Problem

There are `n` cities. Some of them are connected, while some are not. If city `a` is connected directly with city `b`, and city `b` is connected directly with city `c`, then city `a` is connected indirectly with city `c`.

A **province** is a group of directly or indirectly connected cities and no other cities outside of the group.

You are given an `n x n` matrix `isConnected` where `isConnected[i][j] = 1` if the `ith` city and the `jth` city are directly connected, and `isConnected[i][j] = 0` otherwise.

Return *the total number of **provinces***.

 

**Example 1:**

![avatar](..\..\pic\l547-1.jpg)

```
Input: isConnected = [[1,1,0],[1,1,0],[0,0,1]]
Output: 2
```

**Example 2:**

![avatar](..\..\pic\l547-2.jpg)

```
Input: isConnected = [[1,0,0],[0,1,0],[0,0,1]]
Output: 3
```

 

**Constraints:**

- `1 <= n <= 200`
- `n == isConnected.length`
- `n == isConnected[i].length`
- `isConnected[i][j]` is `1` or `0`.
- `isConnected[i][i] == 1`
- `isConnected[i][j] == isConnected[j][i]`

## Solution

- 并查集
- 时间复杂度：近似看做$O(n^2)$，并查集的时间复杂度近似看成是 $O(1)$，两层循环遍历是平方复杂度
- 空间复杂度：$O(n)$

## 并查集

- 解决两类问题：
  - 将两个集合合并
  - 询问两个元素是否在一个集合中
- 基本原理：
  - 每个集合用一棵树来表示，树根的编号就是整个集合的编号
  - 每个节点存储它的父节点，用 p[x] 表示 x 的父节点
- 问题1. 如何判断树根：`if(p[x] == x)`
- 问题2. 如何求 x 的集合编号：
  - 向上遍历父节点，直到找到树根：`while(p[x] != x) x = p[x]`
- 问题3. 如何合并两个集合：
  - 把其中一个集合完整的连接到另一个根节点下面，即把其中一棵树变成另一棵树的子树
  - 假设 p[x] 是 x 的集合编号，p[y] 是 y 的集合编号，则 `p[x] = y`
- 优化：路径压缩
  - 只要执行过一次向上回溯，则把路径上所有节点的父节点都修改为根节点
  - 优化后，并查集的时间复杂度变为***近似O(1)***