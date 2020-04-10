# 116. Populating Next Right Pointers in Each Node

## Problems

You are given a **perfect binary tree** where all leaves are on the same level, and every parent has two children. The binary tree has the following definition:

```
struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
```

Populate each next pointer to point to its next right node. If there is no next right node, the next pointer should be set to `NULL`.

Initially, all next pointers are set to `NULL`.

 

**Follow up:**

- You may only use constant extra space.
- Recursive approach is fine, you may assume implicit stack space does not count as extra space for this problem.

 

**Example 1:**

<img src="E:..\..\pic\lc116problem.png" alt="avatar" style="zoom:60%;" />

```
Input: root = [1,2,3,4,5,6,7]
Output: [1,#,2,3,#,4,5,6,7,#]
Explanation: Given the above perfect binary tree (Figure A), your function should populate each next pointer to point to its next right node, just like in Figure B. The serialized output is in level order as connected by the next pointers, with '#' signifying the end of each level.
```

 

**Constraints:**

- The number of nodes in the given tree is less than `4096`.
- `-1000 <= node.val <= 1000`



## Solutions

- 递归法

  <img src="../../pic\lc116-1.png" alt="avatar" style="zoom:50%;" />

  - 针对每个节点，可以执行以下两步
    -  `node.Left.Next = node.Right`
    - 同时，如果该节点的 next 指针不为空，则 `node.Right.Next = node.Next.Left`
  - 之后递归调用该节点的左节点和右节点分别进行设置，最终需要对所有节点都递归执行以上操作
  - 递归的退出条件 ：当前节点为空或者已经走到了叶子节点
    - 由于满二叉树必然有左右节点，所以可以通过判断某个节点的左节点是否为空来判断其是否是叶子节点
  - 时间复杂度：$O(n)$，空间复杂度：$O(logn)$

- 迭代法

  <img src="../../pic\lc116-2.png" alt="avatar" style="zoom:50%;" />

  - 设置一个 l 指针指向每行的最左边元素，设置一个 p 指针指向 l 元素

  - 对于每一行，p 从 l 开始，设置

    - `p.Left.Next = p.Right`
    - 如果 `p.Next != nil`则` p.Right.Next = p.Next.Left`

  - 换行的时机：当 `p.Next == nil` 时说明当前行已经遍历完毕，执行换行：设置 `l = l.Left`

  - 时间复杂度：$O(n)$；空间复杂度：$O(1)$

    