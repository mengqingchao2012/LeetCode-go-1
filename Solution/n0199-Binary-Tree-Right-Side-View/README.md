# 199. Binary Tree Right Side View

## Problem

Given a binary tree, imagine yourself standing on the *right* side of it, return the values of the nodes you can see ordered from top to bottom.

**Example:**

```
Input: [1,2,3,null,5,null,4]
Output: [1, 3, 4]
Explanation:

   1            <---
 /   \
2     3         <---
 \     \
  5     4       <---
```

## Solution

- 方法一：层序遍历（bfs）

  - 每层的最后一个节点值就是二叉树右视图所看到的值，所以可以使用层序遍历
  - 使用一个辅助队列来模拟层序遍历，为了实现方便，采用根-右-左的方式进行遍历
  - 每切换到新的一层，就将该层的第一个节点的值（该层的最右节点）加入到结果集中
  - 时间复杂度：$O(n)$，空间复杂度：$O(n)$

- 方法二：dfs

  - 按照 根-右-左 的顺序进行深度优先遍历

  - 每层中只把最后一个元素加入到结果集中，故结果集的大小可以用来表示已经遍历过的层数

    - 第 0 层的结果放入结果集后，size = 1，第 1 层完毕后为 2，第 2 层完毕后为 3
    - 即 size - 1 表示已经完成添加节点的层号

  - 还需要一个变量 level 来表示当前正在遍历的层数

    - 如果 level 的值小于结果集的大小 size，就说明当前层的最后一个元素已经放入到结果集中了，只需要继续遍历，不需要将其他值加入结果集
    - 如果 Level 的值等于结果集 size 的大小，说明找到了当前层的最右元素，则将值加入结果集，并将 size + 1

  - 时间复杂度：$O(n)$，空间复杂度：$O(n)$

    