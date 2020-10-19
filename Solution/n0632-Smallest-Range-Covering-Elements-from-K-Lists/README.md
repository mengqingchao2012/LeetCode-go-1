# 632. Smallest Range Covering Elements from K Lists
## Problem

You have `k` lists of sorted integers in **non-decreasing order**. Find the **smallest** range that includes at least one number from each of the `k` lists.

We define the range `[a, b]` is smaller than range `[c, d]` if `b - a < d - c` **or** `a < c` if `b - a == d - c`.

 

**Example 1:**

```
Input: nums = [[4,10,15,24,26],[0,9,12,20],[5,18,22,30]]
Output: [20,24]
Explanation: 
List 1: [4, 10, 15, 24,26], 24 is in range [20,24].
List 2: [0, 9, 12, 20], 20 is in range [20,24].
List 3: [5, 18, 22, 30], 22 is in range [20,24].
```

**Example 2:**

```
Input: nums = [[1,2,3],[1,2,3],[1,2,3]]
Output: [1,1]
```

**Example 3:**

```
Input: nums = [[10,10],[11,11]]
Output: [10,11]
```

**Example 4:**

```
Input: nums = [[10],[11]]
Output: [10,11]
```

**Example 5:**

```
Input: nums = [[1],[2],[3],[4],[5],[6],[7]]
Output: [1,7]
```

 

**Constraints:**

- `nums.length == k`
- `1 <= k <= 3500`
- `1 <= nums[i].length <= 50`
- `-105 <= nums[i][j] <= 105`
- `nums[i]` is sorted in **non-decreasing** order.

## Solution

- 思路：
  - 最小区间的定义：
    - 区间长度最小
    - 相同长度时起点最小
  - 使用一个最小堆来记录堆中元素的最小值，同时维护一个额外的变量 currentMax 表示堆中的最大值
  - 定义 minRange 和 maxRange 表示结果区间的两个端点，分别初始化为 minInt32 和 maxInt32
  - 将每个子数组的第一个元素加入堆中，加入的同时更新 currentMax
  - 当堆中元素等于数组长度时，弹出堆顶元素 elem ，计算当前堆中元素的范围 [elem : currentMax]，如果此时该范围比 [minRange, maxRange] 小，则用 elem 和 currentMax 更新 minRange 和 maxRange
  - 弹出堆顶元素后，如果其后续还有元素，则将后续元素加入堆中，同时要使用加入堆的元素更新 currentMax
  - 最后得到的 minRange 和 maxRange 就是最小范围
  - 时间复杂度：$O(nklogk)$，空间复杂度：$O(k)$