# 295. Find Median from Data Stream
## Problem

Median is the middle value in an ordered integer list. If the size of the list is even, there is no middle value. So the median is the mean of the two middle value.

For example,

```
[2,3,4]`, the median is `3
[2,3]`, the median is `(2 + 3) / 2 = 2.5
```

Design a data structure that supports the following two operations:

- void addNum(int num) - Add a integer number from the data stream to the data structure.
- double findMedian() - Return the median of all elements so far.

 

**Example:**

```
addNum(1)
addNum(2)
findMedian() -> 1.5
addNum(3) 
findMedian() -> 2
```

 

**Follow up:**

1. If all integer numbers from the stream are between 0 and 100, how would you optimize it?
2. If 99% of all integer numbers from the stream are between 0 and 100, how would you optimize it?

## Solution

- 解法：对顶堆
  - 依照题意，中位数的值只跟中间一个或两个元素的值相关，因此将数组分为左右两部分，左边维护一个大顶堆，右边维护一个小顶堆；
  - 执行插入操作时优先插入左边，即如果左边堆为空或者待插入元素小于大顶堆的堆顶元素，则插入左边，否则插入右边；
  - 这里选择在求中位数时才对两个堆中的元素执行重平衡：
    - 如果左边的元素个数小于右边，则弹出右边堆顶元素插入左边
    - 如果数组中的总元素个数为偶数，则要求左右两边元素相同，此时将左边多出来的元素弹出插入右边