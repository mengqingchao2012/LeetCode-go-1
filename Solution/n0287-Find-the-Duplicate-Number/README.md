# 287. Find the Duplicate Number

## Problem
Given an array nums containing n + 1 integers where each integer is between 1 and n (inclusive), prove that at least one duplicate number must exist. Assume that there is only one duplicate number, find the duplicate one.

Example 1:

Input: [1,3,4,2,2]
Output: 2
Example 2:

Input: [3,1,3,4,2]
Output: 3
Note:

1. You must not modify the array (assume the array is read only).
2. You must use only constant, O(1) extra space.
3. Your runtime complexity should be less than O(n2).
4. There is only one duplicate number in the array, but it could be repeated more than once.

## Solution
- 解法一：二分法
    - 使用 low 和 high 表示区间 [1,n] 的范围，如：[1, 5]，则 low = 1, high = 5；
    - 二分法求出中间值 mid，示例中 mid = 3，然后遍历数组，找出数组中值小于等于 mid 的元素的个数 count；
    - 如果 count 比 mid 大，比如 count 为4，说明重复元素出现在mid之前（mid为3，说明前面只有3个位置来填入元素，而现在却数出了4个，
    即有重复），故更新 high = mid。注意不能减一，因为此时重复元素所在的区间是 [0, mid]
    - 反之，如果count 比 mid 小，则说明重复元素出现的在 [mid+1, n]，更新 low = mid + 1
    - 循环退出的条件是 low < high，退出后，low == high，返回 low 即是重复的元素
    - 时间复杂度：O(nlogn)，空间复杂度：O(1)
    
- 解法二：借用查找环形链表的入口元素的思路：

    <img src="..\..\pic\lc287.png" alt="avatar" style="zoom:67%;" />

    - 构建一个虚拟的带环链表，链表的每个节点值为当前元素的值，将当前元素的值当成下标查询出来的下一个值当成节点的 next 指针所指的下一个值：如当前值为3，则该段链表为：`-> 3 -> nums[3] ->`
    - 有了虚拟链表后，就可以使用快慢指针来找到环的入口，即重复的那个元素
    - 时间复杂度为：O(n)，空间复杂度为：O(1)