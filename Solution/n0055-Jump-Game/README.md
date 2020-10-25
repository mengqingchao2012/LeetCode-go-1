# 55. Jump Game
## Problem

Given an array of non-negative integers, you are initially positioned at the first index of the array.

Each element in the array represents your maximum jump length at that position.

Determine if you are able to reach the last index.

 

**Example 1:**

```
Input: nums = [2,3,1,1,4]
Output: true
Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
```

**Example 2:**

```
Input: nums = [3,2,1,0,4]
Output: false
Explanation: You will always arrive at index 3 no matter what. Its maximum jump length is 0, which makes it impossible to reach the last index.
```

 

**Constraints:**

- `1 <= nums.length <= 3 * 10^4`
- `0 <= nums[i][j] <= 10^5`

## Solution

- 思路：
  - 遍历数组，使用一个变量 maxIdx 来记录当前能够到达的最远下标
  - 对于数组中的每个元素：
    - 如果其下标小于 maxIdx，则说明不可达，返回 false
    - 如果 maxIdx 已经大于数组长度，则说明可以到达终点，返回 true
    - 否则就是用当前下标加上元素值来检查是否可以更新 maxIdx
  - 时间复杂度：$O(n)$，空间复杂度：$O(1)$