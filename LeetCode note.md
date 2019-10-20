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



### 17.电话号码的数字组合

>给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
>
>给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
>
>
>
>示例:
>
>输入："23"
>输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
>说明:
>尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。

- 解法一：递归法：

```go
var mapping = map[byte]string {
	0: "abc",
	1: "def",
	2: "ghi",
	3: "jkl",
	4: "mno",
	5: "pqrs",
	6: "tuv",
	7: "wxyz",
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}

	var result []string
	recursion(digits, 0, "", &result)
	return result

}
//辅助函数，用于递归求解
func recursion(digits string, idx int, str string, result *[]string) { //注意数组要传入指针，因为最后要取到返回值
	if idx == len(digits) { //递归的退出条件
		*result = append(*result, str) //注意此处的解引用符号不能漏
		return
	}

	chars := mapping[digits[idx] - '2']
	for i := 0; i < len(chars); i++ {
		nxtStr := str + string(chars[i])
		recursion(digits, idx + 1, nxtStr, result)
	}
}
```

- 思路：
  - 递归解法一般是DFS的，迭代解法一般是BFS的；
  - 对于每个位置，穷举其所有可能得结果，然后回溯到下一个位置继续穷举，比如：`23`这个数字，`2` 的可能取值为 `abc`，当取`a`时，下一位 `3` 的可能取值为 `def`，即产生三个解`ad, ae, af`，回溯，然后`2`取值为`b`，继续穷举；
  - 时间复杂度：$O(4^n)$
    - 最坏情况：数字为`79`时，此时每一位上都有四种可能性（pqrs * wxyz），第一个位置要调用4次，第二个位置要调用 $4^2$ 次，第n个位置要调用 $4^n$ 次，则总时间为：$4 + 4^1 + ... + 4^n$ 
  - 空间复杂度：$O(n)$：取决于最大递归深度，n个字母组合时，最大字母长度为n，递归深度也为n
- 方法二：迭代法：

```go
func letterCombinations1(digits string) []string {
	if digits == "" {
		return nil
	}
	res := []string{""} //注意初始化，此时res的len为1

	for i := 0; i < len(digits); i++ {
		chars := mapping[digits[i] - '2']
		var temp []string
        //对于res中的每一个值，都将其与下一个值的所有可能值进行组合
		for j := 0; j < len(res); j++ { 
			for k := 0; k < len(chars); k++ {
				temp = append(temp, res[j] + string(chars[k]))
			}
		}
		res = temp
	}

	return res
}
```



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
  
    

### 287.寻找重复数

>给定一个包含 n + 1 个整数的数组 nums，其数字都在 1 到 n 之间（包括 1 和 n），可知至少存在一个重复的整数。假设只有一个重复的整数，找出这个重复的数。
>
>示例 1:
>
>输入: [1,3,4,2,2]
>输出: 2
>示例 2:
>
>输入: [3,1,3,4,2]
>输出: 3
>说明：
>
>不能更改原数组（假设数组是只读的）。
>只能使用额外的 O(1) 的空间。
>时间复杂度小于 O(n2) 。
>数组中只有一个重复的数字，但它可能不止重复出现一次。
>

- 方法一：二分法

```go
func findDuplicate(nums []int) int {
	low, high := 1, len(nums) - 1 //注意low从1开始
	for low < high { //注意边界
		mid := low + ((high - low) >> 1)
		count := 0
		for _, num := range nums {
			if num <= mid { //注意边界
				count++
			}
		}

		if count > mid { //注意边界
			high = mid
		} else {
			low = mid + 1 //注意边界
		}
	}

	return low
}
```

- 思路：
  - 题目理解：数组的长度：n+1，数字范围：[1, n]，由此可知，数组中是一定存在重复数字的，比如，n取5，则数组长度为6个数，数字范围：[1, 5]只有5个取值；
  - 用到的数学方法：抽屉原理，简单理解：一个萝卜一个坑，找到中间那个，分别统计左右两边数的个数，如果不存在重复数字，则左右两边个数相同，等于中间数，否则说明多的那一边存在重复数字；
  - low和high不是下标，而是取值范围中的两个极值，mid求出的即是中间数；
  - 统计数组中值小于等于中间数值的数字个数count，如果count的个数大于中间数值，说明重复数值在小于中间值的这一边，反之亦然；
  - 迭代更新low和high，直到两个数相等，此时这个值就是重复数；
  - 时间复杂度：$O(nlog(n))$，空间复杂度：$O(1)$；
- 方法二：模拟带环链表：

```go
func findDuplicate(nums []int) int {
    slow := nums[0]
    fast := nums[nums[0]]
    
    for slow != fast {
        slow = nums[slow]
        fast = nums[nums[fast]]
    }
    
    ptr := 0
    for slow != ptr {
        slow = nums[slow]
        ptr = nums[ptr]
    }
    
    return slow
}
```

- 思路：
  - 由于数值的取值范围小于数组的大小，由此，数值的取值可以当做数组的下标不会越界；
  - 每次以当前的值作为下标，取下一个值，模拟生成一个环形链表，由于重复数的存在，必然会在重复数的地方开始绕环；
  - 使用快慢指针遍历数组，则他们最终会相遇在链表上的一个结点，在用一个指针从链表头开始，慢指针从当前位置开始，每次两个指针各进一步，当两个指针再次相遇时，即是环的入口，即重复数字的值；
  - 时间复杂度：$O(n)$， 空间复杂度：$O(1)$



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