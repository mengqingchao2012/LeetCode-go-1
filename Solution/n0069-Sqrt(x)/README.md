# 69. Sqrt(x)

## Problem
Implement int sqrt(int x).

Compute and return the square root of x, where x is guaranteed to be a non-negative integer.

Since the return type is an integer, the decimal digits are truncated and only the integer part of the result is returned.

Example 1:

Input: 4
Output: 2
Example 2:

Input: 8
Output: 2
Explanation: The square root of 8 is 2.82842..., and since 
             the decimal part is truncated, 2 is returned.
             
## Solution
- 方法一：二分法
    时间复杂度：O(logn)，空间复杂度：O(1)
    
- 方法二：牛顿迭代
    时间复杂度：O(logn)，空间复杂度：O(1)
    
    <img src="..\..\pic\lc69.png" alt="avatar" style="zoom:60%;" />