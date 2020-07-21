# 236. Lowest Common Ancestor of a Binary Tree

## Problem
Given a binary tree, find the lowest common ancestor (LCA) of two given nodes in the tree.

According to the [definition of LCA on Wikipedia](https://en.wikipedia.org/wiki/Lowest_common_ancestor): “The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow **a node to be a descendant of itself**).”

Given the following binary tree: root = [3,5,1,6,2,0,8,null,null,7,4]

![avatar](..\..\pic\l236.png)

 

**Example 1:**

```
Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
Output: 3
Explanation: The LCA of nodes 5 and 1 is 3.
```

**Example 2:**

```
Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
Output: 5
Explanation: The LCA of nodes 5 and 4 is 5, since a node can be a descendant of itself according to the LCA definition.
```

 

**Note:**

- All of the nodes' values will be unique.
- p and q are different and both values will exist in the binary tree.

## Solution

- 解法一：
  - 扩展函数返回值的意义：
    - 对某个节点，如果节点为空，或者节点等于 p、q 中的任一一个节点，则直接返回该节点
    - 否则，分别递归遍历其左右子树，在左右子树上同时寻找 p、q
    - 如果某个节点的左右子树都有返回值，则说明 p、q 各在一颗子树上，节点即是最近祖先，直接返回
    - 如果某个节点只有一棵子树上有返回值，则说明有一个节点在该子树上，返回那个不为空的值
  - 时间复杂度：$O(n)$，空间复杂度：$O(n)$
- 解法二：
  - 遍历二叉树，节点的祖先其实就是从根节点到指定节点路径上所有节点的集合，包括了指定节点在内
  - 分别遍历二叉树，找到两个节点对应的祖先的集合
  - 然后遍历这两个集合，找出其中最后一个相同的节点，就是最近公共祖先
  - 时间复杂度：$O(n)$，空间复杂度：$O(n)$