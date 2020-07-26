# 96. Unique Binary Search Trees

## Problem

Given *n*, how many structurally unique **BST's** (binary search trees) that store values 1 ... *n*?

**Example:**

```
Input: 3
Output: 5
Explanation:
Given n = 3, there are a total of 5 unique BST's:

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3
```

 

**Constraints:**

- `1 <= n <= 19`

## Solution

- 思路：

  - 找规律：

    - 二叉搜索树必须满足：根节点的值大于其左子树中所有节点的值，小于其右子树中所有节点的值
    - 当 n = 1 时，只能组成 1 棵二叉搜索树；
    - 当 n = 2 时，可以组成 2 棵二叉搜索树；
    - 当 n = 3 时：
      - 抽象 dp[i] 为 i 个递增的数字可以组成二叉搜索树的数量；
      - 取第一个数作为根节点，则其右边的两个数可以组成 2 棵二叉搜索树，其左边没有数，则该情况下能组成二叉搜索树的数量为：`dp[0] * dp[2]`
      - 取第二个数作为根节点，其左右各有一个数，则该情况下能组成二叉搜索树的数量为：`dp[1] * dp[1]`
      - 取第三个数作为根节点，其左边有两个数，右边没有数，则该情况下能组成二叉搜索树的数量为：`dp[0] * dp[2]`
      - 令 dp[0] = 1，则将以上结果相加可得，n = 3 时可以组成二叉搜索树的数量为 `1*2 + 1*1 + 2* 1 = 5`
    - 故有：$dp[i] = \sum_{j = 1}^{i} = dp[j-1] * dp[i - j]$；

  - 时间复杂度：$O(n^2)$，空间复杂度：$O(n)$

    