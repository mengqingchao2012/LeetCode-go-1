# 337. House Robber III
## Problem

The thief has found himself a new place for his thievery again. There is only one entrance to this area, called the "root." Besides the root, each house has one and only one parent house. After a tour, the smart thief realized that "all houses in this place forms a binary tree". It will automatically contact the police if two directly-linked houses were broken into on the same night.

Determine the maximum amount of money the thief can rob tonight without alerting the police.

**Example 1:**

```
Input: [3,2,3,null,3,null,1]

     3
    / \
   2   3
    \   \ 
     3   1

Output: 7 
Explanation: Maximum amount of money the thief can rob = 3 + 3 + 1 = 7.
```

**Example 2:**

```
Input: [3,4,5,1,3,null,1]

     3
    / \
   4   5
  / \   \ 
 1   3   1

Output: 9
Explanation: Maximum amount of money the thief can rob = 4 + 5 = 9.
```

## Solution

- 思路：
  - 条件解析：
    - 以当前节点为根节点，则有
    - 偷根节点的最大收益等于根节点的钱加上所有孙子节点的钱
    - 不偷根节点的最大收益等于两个孩子节点能贡献的最大收益，即偷或者不偷两个孩子能拿到的最大收益
  - 使用一个长度为 2 的状态数组 res 来表示最大收益
    - res[0] 表示不偷当前节点能得到的最大收益
    - res[1] 表示偷当前节点能得到的最大收益
    - 由上述分析可知：
      - res[0] = Max(偷左孩子，不偷左孩子) + Max(偷右孩子，不偷右孩子)
      - res[1] = root.Val + 不偷左孩子 + 不偷右孩子
  - 时间复杂度：$O(n)$，空间复杂度：$O(n)$