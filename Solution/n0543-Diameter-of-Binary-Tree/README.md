# 543. Diameter of Binary Tree

## Problem

Given a binary tree, you need to compute the length of the diameter of the tree. The diameter of a binary tree is the length of the **longest** path between any two nodes in a tree. This path may or may not pass through the root.

**Example:**
Given a binary tree

```
          1
         / \
        2   3
       / \     
      4   5    
```


Return **3**, which is the length of the path [4,2,1,3] or [5,2,1,3].

**Note:** The length of path between two nodes is represented by the number of edges between them.


## Solution

- 思路：

  - 如果指定路径必须经过根节点，则只需固定根节点，求出其左右子树的直径，相加即是树的直径

  - 由于题目没有指定必须经过根节点，故需要遍历所有节点，将每个节点依次看做是根节点，求出其直径

  - 自顶向下的方法存在多次重复计算，因此采用自底向上来计算，可以使用递归或者栈来执行

    <img src="..\..\pic\lc543.png" alt="avatar" style="zoom:60%;" />

- 时间复杂度：$O(n)$，空间复杂度：$O(n)$

  