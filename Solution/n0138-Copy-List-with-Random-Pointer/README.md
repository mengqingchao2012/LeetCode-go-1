# 138. Copy List with Random Pointer

## Problems

A linked list is given such that each node contains an additional random pointer which could point to any node in the list or null.

Return a [**deep copy**](https://en.wikipedia.org/wiki/Object_copying#Deep_copy) of the list.

The Linked List is represented in the input/output as a list of `n` nodes. Each node is represented as a pair of `[val, random_index]` where:

- `val`: an integer representing `Node.val`
- `random_index`: the index of the node (range from `0` to `n-1`) where random pointer points to, or `null` if it does not point to any node.

 

**Example 1:**

<img src="..\..\pic\l138-1.png" alt="avatar" style="zoom:80%;" />

```
Input: head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
Output: [[7,null],[13,0],[11,4],[10,2],[1,0]]
```

**Example 2:**

<img src="..\..\pic\l138-2.png" alt="avatar" style="zoom:80%;" />

```
Input: head = [[1,1],[2,1]]
Output: [[1,1],[2,1]]
```

**Example 3:**

![avatar](..\..\pic\l138-3.png)

```
Input: head = [[3,null],[3,0],[3,null]]
Output: [[3,null],[3,0],[3,null]]
```

**Example 4:**

```
Input: head = []
Output: []
Explanation: Given linked list is empty (null pointer), so return null.
```

 

**Constraints:**

- `-10000 <= Node.val <= 10000`
- `Node.random` is null or pointing to a node in the linked list.
- Number of Nodes will not exceed 1000.

## Solutions

- 方法一：

  - 使用一个辅助哈希表，用来存储旧节点到新节点的映射关系
  - 首先遍历一次链表，不考虑 Random 节点，只负责构建哈希表，同时完成主链表的复制
  - 再遍历一次链表，对照哈希表，将 Random 节点的对应新节点加到新链表上
  - 时间复杂度：$O(n)$，空间复杂度：$O(n)$

- 方法二：

  <img src="..\..\pic\l138-4.png" alt="avatar" style="zoom:80%;" />

  - 第一次遍历，使用当前节点值构建新节点，并将新节点插入到当前节点的后面

  - 第二次遍历，由于新节点紧跟在旧节点后，则某个节点的 next 指针指向的就是该节点对应到新链表中的节点，由此通过 `p.Next.Random = p.Random.Next` 即可完成 Random 节点的添加

  - 第三次遍历，使用三个指针将新旧两条链表分开

  - 时间复杂度$O(n)$，空间复杂度：$O(1)$

    