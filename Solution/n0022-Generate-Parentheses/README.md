# 22. Generate Parentheses
## Problem

Given `n` pairs of parentheses, write a function to *generate all combinations of well-formed parentheses*.

 

**Example 1:**

```
Input: n = 3
Output: ["((()))","(()())","(())()","()(())","()()()"]
```

**Example 2:**

```
Input: n = 1
Output: ["()"]
```

 

**Constraints:**

- `1 <= n <= 8`

## Solution

- 思路
  - 两个原则：
    - 只要还有左括号，任意时刻都可以写下左括号
    - 只有在剩余的右括号个数多于左括号个数时，才能写下右括号
  - 时间复杂度：卡特兰数 O(4^n / sqrt(n))，空间复杂度：O(n)