# 241. Different Ways to Add Parentheses
## Problem

Given a string of numbers and operators, return all possible results from computing all the different possible ways to group numbers and operators. The valid operators are `+`, `-` and `*`.

**Example 1:**

```
Input: "2-1-1"
Output: [0, 2]
Explanation: 
((2-1)-1) = 0 
(2-(1-1)) = 2
```

**Example 2:**

```
Input: "2*3-4*5"
Output: [-34, -14, -10, -10, 10]
Explanation: 
(2*(3-(4*5))) = -34 
((2*3)-(4*5)) = -14 
((2*(3-4))*5) = -10 
(2*((3-4)*5)) = -10 
(((2*3)-4)*5) = 10
```

## Solution

- 思路：
  - 中缀表达式可以表示为一颗以运算数为叶子结点，算子为根节点和内部节点的表达式树，该树的中序遍历就是中缀表达式
  - 这样该题就转化为和 95，96 题同样的解题思路