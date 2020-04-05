#  108. Convert Sorted Array to Binary Search Tree

## Problems

Given an array where elements are sorted in ascending order, convert it to a height balanced BST.

For this problem, a height-balanced binary tree is defined as a binary tree in which the depth of the two subtrees of *every* node never differ by more than 1.

**Example:**

```
Given the sorted array: [-10,-3,0,5,9],

One possible answer is: [0,-3,9,-10,null,5], which represents the following height balanced BST:

      0
     / \
   -3   9
   /   /
 -10  5
```



## Solutions

- 二叉搜索树的特点：拍平以后是有序数组，即，根节点永远是中间元素

  - 取数组的中间元素作为根节点，则其左边的元素就是左子树，右边的元素是右子树；
  - 递归的使用左右两边的数组元素构建子树；
  - 递归终止条件：某个数组区间内没有元素存在

- 时间复杂度：$O(n)$，空间复杂度：$O(logn)$

  