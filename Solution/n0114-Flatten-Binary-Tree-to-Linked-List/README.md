# 114. Flatten Binary Tree to Linked List
## Problem

Given a binary tree, flatten it to a linked list in-place.

For example, given the following tree:

    	1
       / \
      2   5
     / \   \
    3   4   6
The flattened tree should look like:

```
1
 \
  2
   \
    3
     \
      4
       \
        5
         \
          6
```

## Solution

- 解法一：

  - 使用前序遍历保存节点，然后再修改 root
  - 时间复杂度：$O(n)$，空间复杂度：$O(n)$（辅助切片和递归深度）

- 解法二：自底向上构建

  - 前序遍历的顺序是：根 —— 左 —— 右，则自底向上构建树的顺序是：右 —— 左 —— 根

  - 由此得到递归关系：

    ```go
    recur(root.Right)
    recur(root.Left)
    // 退递归后处理当前节点，即根节点
    ```

  - 使用一个临时变量 p 指向当前节点，则 p 表示的是已经完成构建的部分

  - 退递归后，将根节点的右子树指向 p，并修改根节点的左子树为空，然后再将 p 指向根节点

  - 时间复杂度：$O(n)$，空间复杂度：$O(n)$（递归深度），实现了一边遍历，一边构建树