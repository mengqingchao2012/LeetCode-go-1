# 494. Target Sum
## Problem

You are given a list of non-negative integers, a1, a2, ..., an, and a target, S. Now you have 2 symbols `+` and `-`. For each integer, you should choose one from `+` and `-` as its new symbol.

Find out how many ways to assign symbols to make sum of integers equal to target S.

**Example 1:**

```
Input: nums is [1, 1, 1, 1, 1], S is 3. 
Output: 5
Explanation: 

-1+1+1+1+1 = 3
+1-1+1+1+1 = 3
+1+1-1+1+1 = 3
+1+1+1-1+1 = 3
+1+1+1+1-1 = 3

There are 5 ways to assign symbols to make the sum of nums be target 3.
```

 

**Constraints:**

- The length of the given array is positive and will not exceed 20.
- The sum of elements in the given array will not exceed 1000.
- Your output answer is guaranteed to be fitted in a 32-bit integer.

## Solution

- 思路：
  - 可以添加的符号是正负号，可以统一看成是加法，满足交换律，则：
    - `+1+1-1+1+1 = 3 等价于 (1 + 1 + 1 + 1) - (1) = 3`
    - 即可以看成是把数组分成了两组，一组全添加加号，另一组全添加减号
  - 由此可以推出：
    - `Sum(N) - Sum(P) = S`
    - `Sum(N) + Sum(P) = Sum(nums)`
    - 两式合并后有：`Sum(N) = (Sum(nums) + S) / 2`
  - 则可以使用第 416 题的解法进行求解