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



### 61. 旋转链表

>给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。
>
>示例 1:
>
>输入: 1->2->3->4->5->NULL, k = 2
>输出: 4->5->1->2->3->NULL
>解释:
>向右旋转 1 步: 5->1->2->3->4->NULL
>向右旋转 2 步: 4->5->1->2->3->NULL
>示例 2:
>
>输入: 0->1->2->NULL, k = 4
>输出: 2->0->1->NULL
>解释:
>向右旋转 1 步: 2->0->1->NULL
>向右旋转 2 步: 1->2->0->NULL
>向右旋转 3 步: 0->1->2->NULL
>向右旋转 4 步: 2->0->1->NULL
>

- 解法：构建成环

```go
type ListNode struct {
    Val int
    Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil {
        return nil
    }
    if head.Next == nil {
        return head
    }
    
    oldTail := head
    var n = 1 //注意n从1开始计数
    for ; oldTail.Next != nil; n++ {
        oldTail = oldTail.Next
    }
    oldTail.Next = head //头尾相连，构建成环
    
    newTail := head
    for i := 0; i < n - k % n - 1; i++ {
        newTail = newTail.Next
    }
    res := newTail.Next
    newTail.Next = nil
    
    return res
}
```

- 思路：
  - 构建成环，则旋转操作其实就变为找到新的头结点和尾结点，然后断开链表；
  - 新链表头的位置：`n - k % n`，链表尾的位置：`n - k % n - 1`，n是链表中结点的个数，注意n从1开始数才对；
  - 时间复杂度：$O(n)$，空间复杂度：$O(1)$



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
  - 初始化时，将加一操作抽象为更低位的进位值，即令carry=1，逆序遍历数组，将数组中的值与carry值相加，和对10取余即得到当前位上的数值，和对10取商，即得到进位的数值；
  - 时间复杂度为：$O(n)$
  
    

### 448.找到所有数组中消失的数字

>给定一个范围在  1 ≤ a[i] ≤ n ( n = 数组大小 ) 的 整型数组，数组中的元素一些出现了两次，另一些只出现一次。
>
>找到所有在 [1, n] 范围之间没有出现在数组中的数字。
>
>您能在不使用额外空间且时间复杂度为O(n)的情况下完成这个任务吗? 你可以假定返回的数组不算在额外空间内。
>
>示例:
>
>输入:
>[4,3,2,7,8,2,3,1]
>
>输出:
>[5,6]
>

- 解法一：

```go
func findDisappearedNumbers(nums []int) []int {
	size := len(nums)
	var res = make([]int, 0, size)
	if size == 0 {
		return nil
	}

	var temp = make([]bool, size)

	for _, v := range nums {
		temp[v - 1] = true
	}
	for k, v := range temp {
		if v != true {
			res = append(res, k + 1)
		}
	}
	return res
}
```

- 思路：
  - 引入一个辅助数组，遍历原数组，将数组中取到的元素值减一当成辅助数组的索引，并将该索引对应的值设为true，代表该数已经出现过；
  - 遍历辅助数组，找出值不为true的所有下标，加一后即可得到数组中没有出现的数；
  - 时间复杂度：$O(n)$，空间复杂度：$O(n)$
- 方法二：

```go
func findDisappearedNumbers(nums []int) []int {
	size := len(nums)
	if size == 0 {
		return nil
	}

	var res = make([]int, 0, size)
	for _, value := range nums {
		idx := abs(value) - 1
		if nums[idx] > 0 {
			nums[idx] = -nums[idx]
		}
	}

	for i, v := range nums {
		if v > 0 {
			res = append(res, i + 1)
		}
	}
	return res
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
```

- 思路：
  - 直接在原数组上进行修改；
  - 要想利用原数组的空间存储数字是否出现的信息，同时又要保证能够取到原来的数值信息，考虑使用相反数；
  - 遍历数组，将取到的元素值取绝对值后减1，当做数组的下标，访问该下标，并将下标对应的值改为负数，遍历完成后，数组中，元素值依旧为正的值所对应的下标+1即是缺失的数；
  - 时间复杂度：$O(n)$，空间复杂度：$O(1)$