# LeetCode note

### 16. 最接近的三数之和

>给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。
>
>例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.
>
>与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).
>

- 解法：

```go
func threeSumClosest(nums []int, target int) int {
    sort.Ints(nums)
    res, min := 0, math.MaxInt64
    
    for i := len(nums) - 1; i >= 2; i-- {
        j, k := 0, i - 1
        for j < k {
            sum := nums[i] + nums[j] + nums[k]
            if sum == target {
                return sum
            } else if sum < target {
                j++
            } else {
                k--
            }
            
            diff := differ(sum, target)
            
            if diff < min {
                min = diff
                res = sum
            }
        }
    }
    return res
}

func differ(a, b int) int {
    if a > b {
        return a - b
    }
    return b - a
}
```

- 思路：
  - 先将数组排序，然后从右边开始，固定住最后一个数，赋值为i，令j=0, k=i-1，即遍历i左边的区间，取出三个数求和并与target进行比较，若小于target，则增大j，若大于target则减小k，若相等则说明找到了要求的三个数；每轮循环结束都将i左移，直到最后区间内数字的个数小于3个，则结束；
  - 使用变量min来保存和与目标之间的差值的最小值，在循环中更新min；
  - 时间复杂度：$O(n^2)$，空间复杂度：$O(1)$



### 66. 加一

>给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
>
>最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
>
>你可以假设除了整数 0 之外，这个整数不会以零开头。
>
>示例 1:
>
>输入: [1,2,3]
>输出: [1,2,4]
>解释: 输入数组表示数字 123。
>示例 2:
>
>输入: [4,3,2,1]
>输出: [4,3,2,2]
>解释: 输入数组表示数字 4321。
>

- 解法：

```go
func plusOne(digits []int) []int {
    carry := 1
    for i := len(digits) - 1; i >= 0; i-- {
        sum := digits[i] + carry
        digits[i] = sum % 10
        carry = sum / 10
    }
    if carry != 0 {
        return append([]int{1}, digits...)
    }
    return digits
    
}
```

- 思路：
  - 关键点：进位会导致相关位上的数值变化，数组的长度也可能会发生变化；
  - 将求和后的结果分成两部分，保留在原位上的数值和要进位的数值；
  - 初始化时，将加一操作抽象为更低位的进位值，即另carry=1，逆序遍历数组，将数组中的值与carry值相加，和对10取余即得到当前位上的数值，和对10取商，即得到进位的数值；
  - 时间复杂度为：$O(n)$