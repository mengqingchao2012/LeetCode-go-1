# 98. Validate Binary Search Tree

## Problem

Given a binary tree, determine if it is a valid binary search tree (BST).

Assume a BST is defined as follows:

- The left subtree of a node contains only nodes with keys **less than** the node's key.
- The right subtree of a node contains only nodes with keys **greater than** the node's key.
- Both the left and right subtrees must also be binary search trees.

 

**Example 1:**

```
    2
   / \
  1   3

Input: [2,1,3]
Output: true
```

**Example 2:**

```
    5
   / \
  1   4
     / \
    3   6

Input: [5,1,4,null,null,3,6]
Output: false
Explanation: The root node's value is 5 but its right child's value is 4.
```

## Solution

- 思路：

  - 二叉搜索树的特征归结起来其实就是：
    - 根节点的值大于左子树节点中的最大值，小于右子树节点中的最小值
    - 且根节点的左右子树也要满足二叉搜索树

- 方法一：

  - 对每个节点，都按照上述两条标准去检查是否满足条件
  - 求某个节点的子树中的最值的时间复杂度是 $logn$，需要对 n 个结点重复执行该操作，所以时间复杂度是：$O(nlogn)$
  - 空间复杂度是$O(n)$

- 方法二：

  - 方法一中每个节点值的上下界重复计算了很多次，考虑将上下界在递归中传递

  - 对于根节点，初始化其上下界为 $[nil, nil]$

  - 对于左节点，其下界不变，上界更新为根节点的值

  - 对于右节点，其上界不变，下界更新为根节点的值

  - 时间复杂度：$O(n)$，空间复杂度：$O(n)$

    