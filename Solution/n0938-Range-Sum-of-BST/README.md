# 938. Range Sum of BST

## Problems

Given the `root` node of a binary search tree, return the sum of values of all nodes with value between `L` and `R` (inclusive).

The binary search tree is guaranteed to have unique values.

 

**Example 1:**

```
Input: root = [10,5,15,3,7,null,18], L = 7, R = 15
Output: 32
```

**Example 2:**

```
Input: root = [10,5,15,3,7,13,18,1,null,6], L = 6, R = 10
Output: 23
```

 

**Note:**

1. The number of nodes in the tree is at most `10000`.
2. The final answer is guaranteed to be less than `2^31`.

## Solutions

- 遍历各节点，将值在范围 $[L,R]$ 之间的所有节点值相加

- 利用二叉搜索树根节点值大于左子树，小于右子树的特点进行剪枝

- 时间复杂度：$O(n)$，空间复杂度：$O(n)$

  