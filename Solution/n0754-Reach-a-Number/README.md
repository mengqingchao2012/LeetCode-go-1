# 754. Reach a Number

## Problem

You are standing at position `0` on an infinite number line. There is a goal at position `target`.

On each move, you can either go left or right. During the *n*-th move (starting from 1), you take *n* steps.

Return the minimum number of steps required to reach the destination.

**Example 1:**

```
Input: target = 3
Output: 2
Explanation:
On the first move we step from 0 to 1.
On the second step we step from 1 to 3.
```



**Example 2:**

```
Input: target = 2
Output: 3
Explanation:
On the first move we step from 0 to 1.
On the second move we step  from 1 to -1.
On the third move we step from -1 to 2.
```



**Note:**

`target` will be a non-zero integer in the range `[-10^9, 10^9]`.

## Solution

- 思路：

  - 目标 target 可以为正也可以为负，但因为数组是对称的，所以如果 target 为负，转为正数求解即可，要注意将 target 的类型转为 int64，防止溢出
  - 假设前 n 步的和为 sum，加入 sum 小于 target，说明此时还未到达 target，因此可以假设前 n 步都是正步长，sum = 1 + 2 + 3 + ... + n
  - 如果第 n + 1 步后，sum大于 target，则假设前 n + 1 步中有一步 x 是负步长，此时有：`sum - t = 2x`，之所以是 2x 是因为 sum 是 n + 1 步正步长的和，现在要将其中的第 x 步设为负步长，则要减掉 2 个 x 才行
  - 由 `sum - t = 2x` 可知，当 sum 和 t 的差为偶数时，可以找到一个满足条件的 n 的值，故有：
    - 当 sum < t，或者 sum - t 的值为奇数时，n++, sum += n

- 时间复杂度：$O(n)$，空间复杂度：$O(1)$

  