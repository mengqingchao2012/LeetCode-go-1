# 92. Reverse Linked List II

## Problem

Reverse a linked list from position *m* to *n*. Do it in one-pass.

**Note:** 1 ≤ *m* ≤ *n* ≤ length of list.

**Example:**

```
Input: 1->2->3->4->5->NULL, m = 2, n = 4
Output: 1->4->3->2->5->NULL
```

## Solution

- 思路：

  <img src="..\..\pic\l92.png" alt="avatar" style="zoom:45%;" />

  - 将链表分为三个部分：`[ : m - 1], [m : n], [n + 1 : ]`
    - 位置 m - 1 是第一部分的最后一个节点，该节点应该连接到翻转后的第一个节点
    - 位置 m 是翻转后的最后一个节点，该节点应该连接到第三部分的第一个节点
  - 确定了对应的位置后，就可以通过遍历链表来找到这些节点并存储起来，同时对第二部分进行翻转，翻转完成后再进行三个部分的连接，具体见代码
  - 特别要注意各类边界条件
  - 时间复杂度：$O(n)$，空间复杂度：$O(1)$