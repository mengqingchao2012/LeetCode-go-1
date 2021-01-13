# 1202. Smallest String With Swaps
## Problem

You are given a string `s`, and an array of pairs of indices in the string `pairs` where `pairs[i] = [a, b]` indicates 2 indices(0-indexed) of the string.

You can swap the characters at any pair of indices in the given `pairs` **any number of times**.

Return the lexicographically smallest string that `s` can be changed to after using the swaps.

 

**Example 1:**

```
Input: s = "dcab", pairs = [[0,3],[1,2]]
Output: "bacd"
Explaination: 
Swap s[0] and s[3], s = "bcad"
Swap s[1] and s[2], s = "bacd"
```

**Example 2:**

```
Input: s = "dcab", pairs = [[0,3],[1,2],[0,2]]
Output: "abcd"
Explaination: 
Swap s[0] and s[3], s = "bcad"
Swap s[0] and s[2], s = "acbd"
Swap s[1] and s[2], s = "abcd"
```

**Example 3:**

```
Input: s = "cba", pairs = [[0,1],[1,2]]
Output: "abc"
Explaination: 
Swap s[0] and s[1], s = "bca"
Swap s[1] and s[2], s = "bac"
Swap s[0] and s[1], s = "abc"
```

 

**Constraints:**

- `1 <= s.length <= 10^5`
- `0 <= pairs.length <= 10^5`
- `0 <= pairs[i][0], pairs[i][1] < s.length`
- `s` only contains lower case English letters.

## Solution

- 并查集
  - 由题意可知：只有指定位置的元素之间可以彼此交换，则这些可交换的元素之间如果相互连通的话，则构成一个连通集
  - 要让最终结果字典序最小，其实就是每个连通集中的字符排列最小，因此可以先找出所有的连通集，再对连通集中的字符进行排序，最后再将排好序的字符插回到原字符串中即可
- 时间复杂度：$O(nlogn + m)$，其中 `nlogn` 是排序的时间复杂度，最坏情况下，整个字符串都属于同一个连通集，字符之间可以两两交换，m 是并查集的近似时间复杂度，表示连通集的个数
- 空间复杂度：$O(n)$