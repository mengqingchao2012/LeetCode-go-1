# 378.Kth Smallest Element in a Sorted Matrix

## Problem

Given a *n* x *n* matrix where each of the rows and columns are sorted in ascending order, find the kth smallest element in the matrix.

Note that it is the kth smallest element in the sorted order, not the kth distinct element.

**Example:**

```
matrix = [
   [ 1,  5,  9],
   [10, 11, 13],
   [12, 13, 15]
],
k = 8,

return 13.
```



**Note:**
You may assume k is always valid, 1 ≤ k ≤ n^2.

## Solution 

- 解法一：
  - 借助合并 k 个有序链表的思路，将数组的每一行看成是一个链表，则第 k 个最小的元素就是链表合并后的第 k - 1 个值
  - 由于数组是行列分别递增的，故第 k 个最小的元素必然在前 k 行中，因此可以先将前 k 行的第一个元素加入堆中，然后循环 k - 1 次，每次弹出堆顶元素，如果堆顶元素不是所在行中的最后一个元素，则将其后面的元素再加入到堆中
  - 最终循环结束后，堆顶元素就是第 k 个最小的元素
  - 时间复杂度：$O(klogk)$，空间复杂度：$O(k)$
- 解法二：
  - 解法一存在的问题：k 的范围是 n^2，由此最坏时间复杂度是 n^2 级别的，非常慢。慢的原因在于只利用了行的有序性而没有使用列的有序性
  - 数组中第 k 小的元素意味着在 k 前面有 k - 1 个元素比他小，根据数组的有序性，最小值是 `matrix[0][0]`，最大值是 `matrix[n][n]`，使用二分法进行搜索
  - 求出 mid 后，要建立 mid 和 k 之间的关系，即可以统计每一行中小于等于 mid 的个数，如果这些个数加起来大于等于 k ，则说明第 k 个元素要么等于 mid ，要么在 mid 的前半段，反之则位于后半段
  - 时间复杂度：$O(nlogv)$，n 表示对于求出的每个 mid，需要线性搜索每一行来统计小于等于 mid 的个数，v 则是最大值与最小值之间的差，是个整数 int32，因此 logv 最大是 32
  - 空间复杂度：$O(1)$