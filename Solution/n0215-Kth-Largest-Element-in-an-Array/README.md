# 215. Kth Largest Element in an Array

## Problem

Find the **k**th largest element in an unsorted array. Note that it is the kth largest element in the sorted order, not the kth distinct element.

**Example 1:**

```
Input: [3,2,1,5,6,4] and k = 2
Output: 5
```

**Example 2:**

```
Input: [3,2,3,1,2,4,5,5,6] and k = 4
Output: 4
```

**Note:**
You may assume k is always valid, 1 ≤ k ≤ array's length.

## Solution

- 方法一：最小堆
  - 维护一个大小为 k 的最小堆，遍历数组：
    - 如果堆未满，则直接加入堆中
    - 如果堆已满，则比较堆顶元素与当前元素的大小，如果堆顶元素小于当前元素，则弹出堆顶元素，当前元素加入堆中
    - 最后返回堆顶元素即可
  - 时间复杂度：$O(nlogk)$，空间复杂度：k
- 方法二：快速选择法（线性求第 k 大）
  - 任选一个元素作为分隔元素，遍历数组，将大于等于该元素的值放到元素左边，比小于等于该元素的值放到元素右边，这样确定分隔点的位置后，通过统计分隔点之前和之后元素的个数，再与 k 进行比较，就可以知道最终结果落在哪个区间内，再递归求解即可
  - 平均时间复杂度：$O(n)$，空间复杂度：$O(1)$
  - 最坏时间复杂度：$O(n^2)$，出现在数组有序的情况下
  - 该方法受分隔点的选择影响较大，在实际场景中之所以有效是因为实际场景中大部分数据都是乱序的