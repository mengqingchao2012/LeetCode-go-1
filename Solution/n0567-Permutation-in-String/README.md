# 567. Permutation in String
## Problem

Given two strings **s1** and **s2**, write a function to return true if **s2** contains the permutation of **s1**. In other words, one of the first string's permutations is the **substring** of the second string.

 

**Example 1:**

```
Input: s1 = "ab" s2 = "eidbaooo"
Output: True
Explanation: s2 contains one permutation of s1 ("ba").
```

**Example 2:**

```
Input:s1= "ab" s2 = "eidboaoo"
Output: False
```

 

**Constraints:**

- The input strings only contain lower case letters.
- The length of both given strings is in range [1, 10,000].

## Solution

- 滑动窗口：
  - 先遍历 s1，得到组成 pattern 的元素分布情况，存储在 map 中
  - 使用一个临时变量 match 来记录当前字符匹配的情况
  - 然后遍历 s2，创建一个滑动窗口，如果 end 对应的值在 map 中，则更新 map[s2[end]] -= 1
    - 此时如果 map[s2[end]] == 0，说明对于该字符已经完全匹配完成，match++
  - 比较 match 和 len(map) 的大小，如果相等，说明所有元素的匹配完成，返回 true
  - 如果 end - start >= len(s1) - 1，即窗口大小超过 pattern 的长度，则：
    - 先判断 map[s2[start]] 是否等于0，等于则要将 match--
    - 更新 map[s2[start]] += 1
    - start 右移
  - 如果循环结束还未提前返回，说明没有找到匹配项，返回 false
- 时间复杂度：$O(n + m)$，n 和 m 为 s1 和 s2 的长度，空间复杂度：$O(n)$