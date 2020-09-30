# 124. Binary Tree Maximum Path Sum
## Problem

Given a **non-empty** binary tree, find the maximum path sum.

For this problem, a path is defined as any sequence of nodes from some starting node to any node in the tree along the parent-child connections. The path must contain **at least one node** and does not need to go through the root.

**Example 1:**

```
Input: [1,2,3]

       1
      / \
     2   3

Output: 6
```

**Example 2:**

```
Input: [-10,9,20,null,null,15,7]

   -10
   / \
  9  20
    /  \
   15   7

Output: 42
```

## Solution

- 注意题目要求：路径不要求一定要从根节点到叶子节点，不一定经过根节点，也不一定经过叶子节点，同时可以只包含一个节点
- 解法一：
  - 将问题分解为：从根节点出发到叶子节点的最大路径和，将这个问题的求解方式抽象成一个方法 f，则对于树中的任意节点，直接使用 f，就可以将问题变为遍历树中的节点，将每个节点都看成根节点求最大路径和，最终在这些结果中选出最大的那个
  - 对于任意节点为根节点的树，其最大路径和等于节点值加上左右子树的最大路径和中较大的那个
  - 需要注意的是，路径和可能为负，负的路径和对于结果没有任何正向作用，可以丢弃。因此求出来的子树路径和要跟 0 求最大值
  - 时间复杂度：$O(n^2)$（n 个节点，每个节点求最大路径和的时间复杂度为 n）
  - 空间复杂度：$O(n)$
- 解法二：
  - 解法一的问题在于对子树的路径和计算存在重复，因此考虑在遍历结点的过程中，每求出一个节点的最大路径和就直接更新最终结果，这样就只需要遍历一次树中所有节点即可
  - 时间复杂度：$O(n)$，空间复杂度：$O(n)$