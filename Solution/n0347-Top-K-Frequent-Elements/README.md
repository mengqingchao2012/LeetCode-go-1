# 347. Top K Frequent Elements

## Problem

Given a non-empty array of integers, return the ***k\*** most frequent elements.

**Example 1:**

```
Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]
```

**Example 2:**

```
Input: nums = [1], k = 1
Output: [1]
```

**Note:**

- You may assume *k* is always valid, 1 ≤ *k* ≤ number of unique elements.
- Your algorithm's time complexity **must be** better than O(*n* log *n*), where *n* is the array's size.
- It's guaranteed that the answer is unique, in other words the set of the top k frequent elements is unique.
- You can return the answer in any order.

## Solution

- 解法一：最小堆
  - 时间复杂度：$O(nlogk)$，空间复杂度：$O(n)$
- 解法二：快速选择法
  - 平均时间复杂度：$O(n)$，空间复杂度：$O(n)$
- 前两种解法类似于第215题
- 解法三：桶排序
  - 有 `n = len(nums)` 个数，则可以构造 0-n 总共 n + 1 个桶，将每个元素按照出现的频率放入桶中
  - 从后往前遍历，取出 k 个元素即是最终结果
  - 时间复杂度：$O(n)$，空间复杂度：$O(n)$