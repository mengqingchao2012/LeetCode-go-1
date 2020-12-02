# 402. Remove K Digits
## Problem

Given a non-negative integer *num* represented as a string, remove *k* digits from the number so that the new number is the smallest possible.

**Note:**

- The length of *num* is less than 10002 and will be ≥ *k*.
- The given *num* does not contain any leading zero.



**Example 1:**

```
Input: num = "1432219", k = 3
Output: "1219"
Explanation: Remove the three digits 4, 3, and 2 to form the new number 1219 which is the smallest.
```



**Example 2:**

```
Input: num = "10200", k = 1
Output: "200"
Explanation: Remove the leading 1 and the number is 200. Note that the output must not contain leading zeroes.
```



**Example 3:**

```
Input: num = "10", k = 2
Output: "0"
Explanation: Remove all the digits from the number and it is left with nothing which is 0.
```

## Solution

- 单调栈
  - 数字的顺序是固定的
  - 遍历数组，如果栈顶元素大于当前待加入元素，则移除栈顶元素显然可以使得最终结果变的更小，故选择移除栈顶元素，加入当前元素
  - 需要注意前缀 0 的问题和数组自身是升序的情况
- 时间复杂度：$O(n)$，空间复杂度：$O(n)$