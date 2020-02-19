# 3. Longest Substring Without Repeating Characters

## Problem
Given a string, find the length of the longest substring without repeating characters.

Example 1:

Input: "abcabcbb"
Output: 3 
Explanation: The answer is "abc", with the length of 3. 

Example 2:

Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3. 
             Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
             
## Solution
- 使用滑动窗口法：start 和 end 两个指针从字符串开头开始往后移动，使用一个辅助 map（可以换成数组）来记录字符出现的位置。
  注意字符出现的位置应该等于 end 下标加 1，意味着如果出现重复字符的话，下一个字符开始才是不重复的。

- 示例：
    s = "pwwkew"
    
    <img src="..\..\pic\lc3.png" alt="avatar" style="zoom:67%;" />