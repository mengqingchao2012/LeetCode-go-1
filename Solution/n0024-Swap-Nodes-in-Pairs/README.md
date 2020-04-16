# 24. Swap Nodes in Pairs

## Problems

Given a linked list, swap every two adjacent nodes and return its head.

You may **not** modify the values in the list's nodes, only nodes itself may be changed.

 

**Example:**

```
Given 1->2->3->4, you should return the list as 2->1->4->3.
```

## Solutions

- 解法一：迭代法

  <img src="../../pic\l24-1.png" alt="avatar" style="zoom:50%;" />

  - 如图，需要三个指针来进行交换
  - 时间复杂度：$O(n)$，空间复杂度：$O(1)$

- 解法二：递归法

  <img src="..\..\pic\l24-2.png" alt="avatar" style="zoom:67%;" />

  - 退递归条件：节点为空或者链表只剩下一个节点

  - 每次处理两个节点，使用 h 和 n 两个指针来操作，完成当前轮交换后，递归对以 n 的下一个节点为头结点的子链表进行操作

  - 最终返回新的头结点 n

  - 不需要 dummy 节点的原因是，递归执行是从后向前的，也就是说，子链表永远都是先进行操作，最后返回的 n 就是新的头结点

  - 时间复杂度：$O(n)$，空间复杂度：$O(n)$

    