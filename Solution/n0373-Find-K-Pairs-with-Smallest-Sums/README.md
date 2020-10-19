# 373. Find K Pairs with Smallest Sums
## Problem

You are given two integer arrays **nums1** and **nums2** sorted in ascending order and an integer **k**.

Define a pair **(u,v)** which consists of one element from the first array and one element from the second array.

Find the k pairs **(u1,v1),(u2,v2) ...(uk,vk)** with the smallest sums.

**Example 1:**

```
Input: nums1 = [1,7,11], nums2 = [2,4,6], k = 3
Output: [[1,2],[1,4],[1,6]] 
Explanation: The first 3 pairs are returned from the sequence: 
             [1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]
```

**Example 2:**

```
Input: nums1 = [1,1,2], nums2 = [1,2,3], k = 2
Output: [1,1],[1,1]
Explanation: The first 2 pairs are returned from the sequence: 
             [1,1],[1,1],[1,2],[2,1],[1,2],[2,2],[1,3],[1,3],[2,3]
```

**Example 3:**

```
Input: nums1 = [1,2], nums2 = [3], k = 3
Output: [1,3],[2,3]
Explanation: All possible pairs are returned from the sequence: [1,3],[2,3]
```

## Solution 

- 思路：
  - 要求和最小的 k 对数，维护一个大顶堆
  - 两层循环遍历两个数组，对于 nums1 的每个元素，遍历 nums2 取出每个元素进行和计算
  - 如果当前堆中元素还不足 k 个，则将两个元素的下标以及其和组成的数组加入堆中
  - 如果当前堆中元素已经有 k 个，则：
    - 如果新加入的元素和大于堆顶元素，退出内层循环
      - 因为数组是递增的，之后的所有元素和都会比当前堆顶元素大
    - 否则弹出堆顶元素，将新的元素和加入堆中
  - 时间复杂度：$O(mnlogk)$，空间复杂度：$O(k)$