# 1319 Number of Operations to Make Network Connected
## Problem

There are `n` computers numbered from `0` to `n-1` connected by ethernet cables `connections` forming a network where `connections[i] = [a, b]` represents a connection between computers `a` and `b`. Any computer can reach any other computer directly or indirectly through the network.

Given an initial computer network `connections`. You can extract certain cables between two directly connected computers, and place them between any pair of disconnected computers to make them directly connected. Return the *minimum number of times* you need to do this in order to make all the computers connected. If it's not possible, return -1. 

 

**Example 1:**

**![avatar](..\..\pic\l1319-1.png)**

```
Input: n = 4, connections = [[0,1],[0,2],[1,2]]
Output: 1
Explanation: Remove cable between computer 1 and 2 and place between computers 1 and 3.
```

**Example 2:**

**![avatar](..\..\pic\l1319-2.png)**

```
Input: n = 6, connections = [[0,1],[0,2],[0,3],[1,2],[1,3]]
Output: 2
```

**Example 3:**

```
Input: n = 6, connections = [[0,1],[0,2],[0,3],[1,2]]
Output: -1
Explanation: There are not enough cables.
```

**Example 4:**

```
Input: n = 5, connections = [[0,1],[0,2],[3,4],[2,3]]
Output: 0
```

 

**Constraints:**

- `1 <= n <= 10^5`
- `1 <= connections.length <= min(n*(n-1)/2, 10^5)`
- `connections[i].length == 2`
- `0 <= connections[i][0], connections[i][1] < n`
- `connections[i][0] != connections[i][1]`
- There are no repeated connections.
- No two computers are connected by more than one cable.

## Solution

- 并查集
- 两台计算机连入同一个网络需要一根网线，则如果有 n 台计算机要新加入同一个网络，则需要 n - 1 根网线
- 该题求的不是新加入，而是移动冗余的网线，那么假设有 x 台单独的计算机或计算机组要加入到当前网络，就需要 x - 1 根冗余的网线。则有：
  - 如果 connections 的长度小于 n - 1，则说明没有足够的冗余网线来移动，返回 -1
  - 使用并查集求出连通节点的个数，同时在遍历 connections 数组时，如果当前分组中的两台计算机已经处在同一个网络中，即他们都有共同的根节点，则说明有一根冗余的网线可以移动。记录下冗余的网线数和连通节点的个数
  - 如果冗余网线数小于 n - 连通节点的个数，则说明无法完成组网，返回 -1，否则最小移动的次数就等于 n - 连通节点的个数
- 时间复杂度：近似于 $O(n)$，空间复杂度：$O(n)$