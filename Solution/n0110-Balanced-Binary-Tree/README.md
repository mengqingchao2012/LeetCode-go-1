# 110.Balanced Binary Tree

## Problem

Given a binary tree, determine if it is height-balanced.

For this problem, a height-balanced binary tree is defined as:

> a binary tree in which the left and right subtrees of *every* node differ in height by no more than 1.

 

**Example 1:**

Given the following tree `[3,9,20,null,null,15,7]`:

```
    3
   / \
  9  20
    /  \
   15   7
```

Return true.

**Example 2:**

Given the following tree `[1,2,2,3,3,null,null,4,4]`:

```
       1
      / \
     2   2
    / \
   3   3
  / \
 4   4
```

Return false.

## Solution

- 方法一：

  - 使用一个辅助函数来求解各个结点的高度
  - 对于每一个结点，分别计算其左右子树的高度，如果满足左右子树高度差不大于 1 且左右子树都是平衡二叉树，则该二叉树是平衡二叉树
  - 时间复杂度：
    - 树高为 logn，对于根结点（第一层），需要遍历检查所有结点，故检查根节点的时间复杂度是 n，对于第二层，有两个结点，检查每个结点都需要遍历一半的结点，故检查第二层的时间复杂度也是 n，以此类推，检查每一层的时间结点都是 n
    - 故时间复杂度：$O(nlogn)$，空间复杂度：$O(n)$

- 方法二：

  - 方法一的问题在于，检查是自顶向下开始的，即下面的结点被重复计算了多次，一个改进就是在遍历的过程中直接完成检查，如果某个子树结点不满足要求，则直接返回 false

  - 时间复杂度：$O(n)$，空间复杂度：$O(n)$

    