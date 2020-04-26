# 437. Path Sum III

## Problem

You are given a binary tree in which each node contains an integer value.

Find the number of paths that sum to a given value.

The path does not need to start or end at the root or a leaf, but it must go downwards (traveling only from parent nodes to child nodes).

The tree has no more than 1,000 nodes and the values are in the range -1,000,000 to 1,000,000.

**Example:**

```
root = [10,5,-3,3,2,null,11,3,-2,null,1], sum = 8

      10
     /  \
    5   -3
   / \    \
  3   2   11
 / \   \
3  -2   1

Return 3. The paths that sum to 8 are:

1.  5 -> 3
2.  5 -> 2 -> 1
3. -3 -> 11
```

## Solution

- 方法一：

  - 遍历每个节点，把每个节点都当成根节点，求出可能的路径数量，见代码
  - 时间复杂度：$O(n^2)$，空间复杂度：$O(n)$

- 方法二：前缀和

  - 前缀和的思路见第 560 题

  - 时间复杂度：$O(n)$，空间复杂度：$O(n)$

    