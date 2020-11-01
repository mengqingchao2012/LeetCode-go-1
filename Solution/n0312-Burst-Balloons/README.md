# 312. Burst Balloons
## Problem

Given `n` balloons, indexed from `0` to `n-1`. Each balloon is painted with a number on it represented by array `nums`. You are asked to burst all the balloons. If the you burst balloon `i` you will get `nums[left] * nums[i] * nums[right]` coins. Here `left` and `right` are adjacent indices of `i`. After the burst, the `left` and `right` then becomes adjacent.

Find the maximum coins you can collect by bursting the balloons wisely.

**Note:**

- You may imagine `nums[-1] = nums[n] = 1`. They are not real therefore you can not burst them.
- 0 ≤ `n` ≤ 500, 0 ≤ `nums[i]` ≤ 100

**Example:**

```
Input: [3,1,5,8]
Output: 167 
Explanation: nums = [3,1,5,8] --> [3,5,8] -->   [3,8]   -->  [8]  --> []
             coins =  3*1*5      +  3*5*8    +  1*3*8      + 1*8*1   = 167
```

## Solution

- 区间 dp：

  - 状态数组：`dp[i][j]` 表示戳破区间 `[i + 1, j - 1]` 内的所有气球能够得到的最大分数

    - 注意戳破的区间不包含 i 和 j，因为需要用到 i，j 的值来加入计算

    - 对于区间 `[i + 1, j - 1] ` 内的气球，假设最后一个戳破的气球是 k

      ![avatar](..\..\pic\lc312.png)

    - 则此时戳破气球 K 的收益为 `nums[i] * nums[K] * nums[j]`，

    - 当前状态下整个区间内的收益是 `dp[I][K] + dp[K][J] + nums[i] * nums[K] * nums[j]`

    - 由此，通过确定 i 和 j，然后在区间内枚举 k 即可得到最大的收益

  - 时间复杂度：$O(n^3)$，空间复杂度：$O(n^2)$