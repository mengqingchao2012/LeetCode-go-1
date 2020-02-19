# 2. Add Two Numbers

## Problem
You are given two non-empty linked lists representing two non-negative integers. The digits are stored 
in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.

## Solution
- 要注意的点：

    a. 两个链表不一定等长，注意链表尾的判断
    
    b. 注意处理进位问题
    
- 思路：

    - 使用一个虚拟头结点来保存结果的头指针  

    - 使用一个变量 carry 来保存每次计算后进位的值，则有：
    
        `CurrentNode.Val = (l1.Val + l2.Val + carry) % 10`
        
        `carry = (l1.Val + l2.Val + carry) / 10`
        
    - 求和不一定要用 `sum = l1.Val + l2.Val + carry`，这样需要同时满足两个指针不为空，还要附带判断 carry 的状态，
    不好实现。等价的写法是：用 sum 变量分别去与两个指针求和，同时在每轮初始化时，直接将 sum 初始化为上一轮的 carry，
    既方便了条件判断，又减少了一次加法运算（初始化时，sum = 0 + carry）
    
- 时间复杂度：O(n)，空间复杂度：O(1)