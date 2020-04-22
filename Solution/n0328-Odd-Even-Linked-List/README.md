# 328. Odd Even Linked List

## Problem

Given a singly linked list, group all odd nodes together followed by the even nodes. Please note here we are talking about the node number and not the value in the nodes.

You should try to do it in place. The program should run in O(1) space complexity and O(nodes) time complexity.

**Example 1:**

```
Input: 1->2->3->4->5->NULL
Output: 1->3->5->2->4->NULL
```

**Example 2:**

```
Input: 2->1->3->5->6->4->7->NULL
Output: 2->3->6->7->1->5->4->NULL
```

**Note:**

- The relative order inside both the even and odd groups should remain as it was in the input.
- The first node is considered odd, the second node even and so on ...

## Solution

- 思路：

  - 使用两个指针来分别构建奇偶链表，因为最后需要将偶链表连接到奇链表的后面，所以还需要一个单独的指针来保存偶链表的头结点

  - 注意链表长度为奇数和偶数时退出的条件不同

    <img src="..\..\pic\lc-328.png" alt="avatar" style="zoom:80%;" />

- 时间复杂度：$O(n)$，空间复杂度：$O(1)$

  